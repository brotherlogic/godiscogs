package godiscogs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultDiscogsBaseURL = "https://api.discogs.com"

// SearchParams defines query parameters for Discogs GET /database/search.
type SearchParams struct {
	Query        string // q
	Type         string // e.g. "release"; if set, sent as type=...
	Title        string // title
	ReleaseTitle string // release_title
	Credit       string // credit
	Artist       string // artist
	ANV          string // anv
	Label        string // label
	Genre        string // genre
	Style        string // style
	Country      string // country
	Year         string // year
	Format       string // format
	Catno        string // catno
	Barcode      string // barcode
	Track        string // track
	Submitter    string // submitter
	Contributor  string // contributor
	PerPage      int    // per_page (single page only; page=1)
	Sort         string // sort
	SortOrder    string // sort_order ("asc" or "desc")
}

// SearchResult represents a single item returned by Discogs GET /database/search.
type SearchResult struct {
	ID          int32    `json:"id"`
	Title       string   `json:"title"`
	Year        string   `json:"year"`
	Formats     []string `json:"format"`
	Country     string   `json:"country"`
	URI         string   `json:"uri"`
	ResourceURL string   `json:"resource_url"`
}

// UnmarshalJSON unmarshals a SearchResult, accepting string, numeric, null, or omitted year values.
func (s *SearchResult) UnmarshalJSON(data []byte) error {
	type searchResultAlias struct {
		ID          int32           `json:"id"`
		Title       string          `json:"title"`
		Year        json.RawMessage `json:"year"`
		Formats     []string        `json:"format"`
		Country     string          `json:"country"`
		URI         string          `json:"uri"`
		ResourceURL string          `json:"resource_url"`
	}

	var raw searchResultAlias
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s.ID = raw.ID
	s.Title = raw.Title
	s.Formats = raw.Formats
	s.Country = raw.Country
	s.URI = raw.URI
	s.ResourceURL = raw.ResourceURL
	s.Year = ""

	trimmed := strings.TrimSpace(string(raw.Year))
	if trimmed == "" || trimmed == "null" {
		return nil
	}
	if strings.HasPrefix(trimmed, "\"") {
		var yr string
		if err := json.Unmarshal(raw.Year, &yr); err != nil {
			return err
		}
		s.Year = yr
		return nil
	}

	var num json.Number
	if err := json.Unmarshal(raw.Year, &num); err != nil {
		return err
	}
	s.Year = num.String()
	return nil
}

type searchResponse struct {
	Pagination Pagination      `json:"pagination"`
	Results    []*SearchResult `json:"results"`
}

func buildSearchURL(baseURL string, params SearchParams) string {
	if baseURL == "" {
		baseURL = defaultDiscogsBaseURL
	}
	endpoint := strings.TrimRight(baseURL, "/") + "/database/search"

	q := url.Values{}
	fields := []struct {
		key string
		val string
	}{
		{"q", params.Query},
		{"type", params.Type},
		{"title", params.Title},
		{"release_title", params.ReleaseTitle},
		{"credit", params.Credit},
		{"artist", params.Artist},
		{"anv", params.ANV},
		{"label", params.Label},
		{"genre", params.Genre},
		{"style", params.Style},
		{"country", params.Country},
		{"year", params.Year},
		{"format", params.Format},
		{"catno", params.Catno},
		{"barcode", params.Barcode},
		{"track", params.Track},
		{"submitter", params.Submitter},
		{"contributor", params.Contributor},
		{"sort", params.Sort},
		{"sort_order", params.SortOrder},
	}
	for _, f := range fields {
		if f.val != "" {
			q.Set(f.key, f.val)
		}
	}
	if params.PerPage > 0 {
		q.Set("per_page", strconv.Itoa(params.PerPage))
		q.Set("page", "1")
	}

	if encoded := q.Encode(); encoded != "" {
		return endpoint + "?" + encoded
	}
	return endpoint
}

func redactToken(s, token string) string {
	if token == "" {
		return s
	}
	return strings.ReplaceAll(s, token, "[REDACTED]")
}

// Search queries Discogs GET /database/search for a single page of matching results.
func (r *DiscogsRetriever) Search(ctx context.Context, params SearchParams) ([]*SearchResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	endpointURL := buildSearchURL(r.baseURL, params)

	t1 := time.Now()
	if _, err := r.throttleContext(ctx); err != nil {
		return nil, status.FromContextError(ctx.Err()).Err()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create discogs search request: %v", redactToken(err.Error(), r.userToken))
	}
	if r.userToken != "" {
		req.Header.Set("Authorization", "Discogs token="+r.userToken)
	}
	if r.userAgent != "" {
		req.Header.Set("User-Agent", r.userAgent)
	}

	client := r.httpClient
	if client == nil {
		client = http.DefaultClient
	}

	t2 := time.Now()
	DiscogsRequests.With(prometheus.Labels{"method": "GET", "path1": "database"}).Inc()
	resp, err := client.Do(req)
	RequestLatency.With(prometheus.Labels{"method": "GET", "path1": "database"}).Observe(float64(time.Since(t2).Milliseconds()))
	r.Log(ctx, fmt.Sprintf("%v in %v with %v", endpointURL, time.Since(t2), t2.Sub(t1)))

	if err != nil {
		if ctx.Err() != nil {
			return nil, status.FromContextError(ctx.Err()).Err()
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.FromContextError(context.DeadlineExceeded).Err()
		}
		if errors.Is(err, context.Canceled) {
			return nil, status.FromContextError(context.Canceled).Err()
		}
		return nil, status.Errorf(codes.Unavailable, "discogs search request failed: %v", redactToken(err.Error(), r.userToken))
	}
	defer resp.Body.Close()

	r.updateRateLimit(ctx, resp, "GET")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if ctx.Err() != nil {
			return nil, status.FromContextError(ctx.Err()).Err()
		}
		return nil, status.Errorf(codes.Unavailable, "failed to read discogs search response: %v", redactToken(err.Error(), r.userToken))
	}

	throttleMu.Lock()
	if now := time.Now(); now.After(lastTimeRetrieved) {
		lastTimeRetrieved = now
	}
	throttleMu.Unlock()

	if resp.StatusCode != http.StatusOK {
		safeBody := redactToken(string(body), r.userToken)
		r.Log(ctx, fmt.Sprintf("SEARCH (%v) %v -> %v", endpointURL, resp.StatusCode, safeBody))
		switch {
		case resp.StatusCode == http.StatusTooManyRequests:
			return nil, status.Errorf(codes.ResourceExhausted, "discogs search rate limited (%v): %v", resp.StatusCode, safeBody)
		case resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden:
			return nil, status.Errorf(codes.FailedPrecondition, "discogs search unauthorized (%v): %v", resp.StatusCode, safeBody)
		case resp.StatusCode >= 500 && resp.StatusCode < 600:
			return nil, status.Errorf(codes.Unavailable, "discogs search server error (%v): %v", resp.StatusCode, safeBody)
		case resp.StatusCode == http.StatusNotFound:
			return nil, status.Errorf(codes.NotFound, "discogs search not found (%v): %v", resp.StatusCode, safeBody)
		case resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnprocessableEntity:
			return nil, status.Errorf(codes.InvalidArgument, "discogs search invalid request (%v): %v", resp.StatusCode, safeBody)
		default:
			return nil, status.Errorf(codes.Internal, "discogs search unexpected status (%v): %v", resp.StatusCode, safeBody)
		}
	}

	unmarshaller := r.unmarshaller
	if unmarshaller == nil {
		unmarshaller = prodUnmarshaller{}
	}

	var sr searchResponse
	if err := unmarshaller.Unmarshal(body, &sr); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal discogs search response: %v", err)
	}

	if sr.Results == nil {
		return []*SearchResult{}, nil
	}
	return sr.Results, nil
}
