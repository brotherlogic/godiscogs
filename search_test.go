package godiscogs

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newSearchTestRetriever(serverURL, token string, logger func(context.Context, string)) *DiscogsRetriever {
	r := NewDiscogsRetriever(token, logger)
	r.baseURL = serverURL
	r.getSleep = 0
	return r
}

func TestSearchQueryEncodingAllFields(t *testing.T) {
	var gotPath string
	var gotQuery url.Values
	var gotAuth string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQuery = r.URL.Query()
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("X-Discogs-Ratelimit", "60")
		w.Header().Set("X-Discogs-Ratelimit-Used", "1")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"pagination":{"pages":1,"page":1},"results":[]}`))
	}))
	defer srv.Close()

	retr := newSearchTestRetriever(srv.URL, "secret-token-123", nil)
	params := SearchParams{
		Query:        "kind of blue",
		Type:         "release",
		Title:        "Kind Of Blue",
		ReleaseTitle: "Kind Of Blue (Mono)",
		Credit:       "Bill Evans",
		Artist:       "Miles Davis",
		ANV:          "Miles",
		Label:        "Columbia",
		Genre:        "Jazz",
		Style:        "Modal",
		Country:      "US",
		Year:         "1959",
		Format:       "Vinyl",
		Catno:        "CL 1355",
		Barcode:      "07464081632",
		Track:        "So What",
		Submitter:    "user1",
		Contributor:  "user2",
		PerPage:      25,
		Sort:         "year",
		SortOrder:    "asc",
	}

	results, err := retr.Search(context.Background(), params)
	if err != nil {
		t.Fatalf("Search returned unexpected error: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}

	if gotPath != "/database/search" {
		t.Errorf("unexpected path: got %q, want %q", gotPath, "/database/search")
	}
	if gotAuth != "Discogs token=secret-token-123" {
		t.Errorf("unexpected Authorization header: got %q", gotAuth)
	}
	if gotQuery.Has("token") {
		t.Errorf("token must not be sent as query parameter, got %q", gotQuery.Get("token"))
	}

	expected := map[string]string{
		"q":             "kind of blue",
		"type":          "release",
		"title":         "Kind Of Blue",
		"release_title": "Kind Of Blue (Mono)",
		"credit":        "Bill Evans",
		"artist":        "Miles Davis",
		"anv":           "Miles",
		"label":         "Columbia",
		"genre":         "Jazz",
		"style":         "Modal",
		"country":       "US",
		"year":          "1959",
		"format":        "Vinyl",
		"catno":         "CL 1355",
		"barcode":       "07464081632",
		"track":         "So What",
		"submitter":     "user1",
		"contributor":   "user2",
		"per_page":      "25",
		"page":          "1",
		"sort":          "year",
		"sort_order":    "asc",
	}
	if len(gotQuery) != len(expected) {
		t.Errorf("unexpected number of query params: got %d (%v), want %d", len(gotQuery), gotQuery, len(expected))
	}
	for k, want := range expected {
		if got := gotQuery.Get(k); got != want {
			t.Errorf("query param %q = %q, want %q", k, got, want)
		}
	}
}

func TestSearchOmitsEmptyFields(t *testing.T) {
	var gotRawQuery string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotRawQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"results":[]}`))
	}))
	defer srv.Close()

	retr := newSearchTestRetriever(srv.URL, "token", nil)
	if _, err := retr.Search(context.Background(), SearchParams{}); err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if gotRawQuery != "" {
		t.Errorf("expected empty query string for empty SearchParams, got %q", gotRawQuery)
	}
}

func TestSearchSpecialCharactersAndUnicode(t *testing.T) {
	var gotQuery url.Values
	var gotRawQuery string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query()
		gotRawQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"results":[]}`))
	}))
	defer srv.Close()

	retr := newSearchTestRetriever(srv.URL, "token", nil)
	params := SearchParams{
		Query:  `Björk & "Friends" / 100% + 日本語`,
		Artist: "AC/DC & Co.",
		Title:  "Track #1? [Live]",
	}

	if _, err := retr.Search(context.Background(), params); err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if gotQuery.Get("q") != params.Query {
		t.Errorf("decoded q = %q, want %q", gotQuery.Get("q"), params.Query)
	}
	if gotQuery.Get("artist") != params.Artist {
		t.Errorf("decoded artist = %q, want %q", gotQuery.Get("artist"), params.Artist)
	}
	if gotQuery.Get("title") != params.Title {
		t.Errorf("decoded title = %q, want %q", gotQuery.Get("title"), params.Title)
	}
	if strings.Contains(gotRawQuery, "Björk") || strings.Contains(gotRawQuery, "日本語") {
		t.Errorf("expected raw query to be percent-encoded, got %q", gotRawQuery)
	}
}

func TestSearchJSONUnmarshalingResilience(t *testing.T) {
	payload := `{
		"pagination": {"pages": 1, "page": 1},
		"results": [
			{
				"id": 101,
				"title": "String Year Album",
				"year": "2020",
				"format": ["Vinyl", "LP", "Album"],
				"country": "UK",
				"uri": "/release/101",
				"resource_url": "https://api.discogs.com/releases/101"
			},
			{
				"id": 202,
				"title": "Numeric Year Album",
				"year": 1994,
				"format": ["CD"],
				"country": "US",
				"uri": "/release/202",
				"resource_url": "https://api.discogs.com/releases/202"
			},
			{
				"id": 303,
				"title": "Null Year And Missing Format",
				"year": null
			},
			{
				"id": 404,
				"title": "Omitted Year"
			}
		]
	}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(payload))
	}))
	defer srv.Close()

	retr := newSearchTestRetriever(srv.URL, "token", nil)
	results, err := retr.Search(context.Background(), SearchParams{Query: "test"})
	if err != nil {
		t.Fatalf("Search returned unexpected error: %v", err)
	}
	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}

	if results[0].ID != 101 || results[0].Year != "2020" || len(results[0].Formats) != 3 || results[0].Country != "UK" {
		t.Errorf("unexpected result[0]: %+v", results[0])
	}
	if results[1].ID != 202 || results[1].Year != "1994" || len(results[1].Formats) != 1 || results[1].Formats[0] != "CD" {
		t.Errorf("unexpected result[1]: %+v", results[1])
	}
	if results[2].ID != 303 || results[2].Year != "" || results[2].Formats != nil {
		t.Errorf("unexpected result[2]: %+v", results[2])
	}
	if results[3].ID != 404 || results[3].Year != "" || results[3].Title != "Omitted Year" {
		t.Errorf("unexpected result[3]: %+v", results[3])
	}
}

func TestSearchZeroResults(t *testing.T) {
	cases := []string{
		`{"pagination":{"pages":0,"page":1},"results":[]}`,
		`{}`,
	}

	for i, body := range cases {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(body))
		}))
		retr := newSearchTestRetriever(srv.URL, "token", nil)
		results, err := retr.Search(context.Background(), SearchParams{Query: "nothing"})
		srv.Close()

		if err != nil {
			t.Fatalf("case %d: expected nil error, got %v", i, err)
		}
		if results == nil || len(results) != 0 {
			t.Fatalf("case %d: expected non-nil empty slice, got %#v", i, results)
		}
	}
}

func TestSearchErrorMappingsAndTokenRedaction(t *testing.T) {
	const secretToken = "super-secret-discogs-token-98765"

	cases := []struct {
		name       string
		statusCode int
		body       string
		wantCode   codes.Code
	}{
		{"rate_limited", http.StatusTooManyRequests, "rate limit hit with " + secretToken, codes.ResourceExhausted},
		{"unauthorized", http.StatusUnauthorized, "invalid token " + secretToken, codes.FailedPrecondition},
		{"forbidden", http.StatusForbidden, "forbidden for " + secretToken, codes.FailedPrecondition},
		{"internal_500", http.StatusInternalServerError, "upstream boom " + secretToken, codes.Unavailable},
		{"bad_gateway_502", http.StatusBadGateway, "bad gateway", codes.Unavailable},
		{"unavailable_503", http.StatusServiceUnavailable, "unavailable", codes.Unavailable},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var mu sync.Mutex
			var logs []string
			logger := func(_ context.Context, msg string) {
				mu.Lock()
				logs = append(logs, msg)
				mu.Unlock()
			}

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer srv.Close()

			retr := newSearchTestRetriever(srv.URL, secretToken, logger)
			_, err := retr.Search(context.Background(), SearchParams{Query: "test"})
			if err == nil {
				t.Fatalf("expected error for status %d, got nil", tc.statusCode)
			}
			if gotCode := status.Code(err); gotCode != tc.wantCode {
				t.Errorf("status.Code(err) = %v, want %v (err: %v)", gotCode, tc.wantCode, err)
			}
			if strings.Contains(err.Error(), secretToken) {
				t.Errorf("error string leaked userToken: %q", err.Error())
			}
			for _, l := range logs {
				if strings.Contains(l, secretToken) {
					t.Errorf("log entry leaked userToken: %q", l)
				}
			}
		})
	}

	t.Run("network_error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		urlStr := srv.URL
		srv.Close()

		retr := newSearchTestRetriever(urlStr, secretToken, nil)
		_, err := retr.Search(context.Background(), SearchParams{Query: "test"})
		if err == nil {
			t.Fatalf("expected network error, got nil")
		}
		if gotCode := status.Code(err); gotCode != codes.Unavailable {
			t.Errorf("status.Code(err) = %v, want %v (err: %v)", gotCode, codes.Unavailable, err)
		}
		if strings.Contains(err.Error(), secretToken) {
			t.Errorf("network error leaked userToken: %q", err.Error())
		}
	})
}

func TestSearchContextCancellationDuringThrottleWait(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"results":[]}`))
	}))
	defer srv.Close()

	// Seed lastTimeRetrieved using a fast request, then test cancellation on a long throttle wait.
	retr := newSearchTestRetriever(srv.URL, "token", nil)
	if _, err := retr.Search(context.Background(), SearchParams{Query: "seed"}); err != nil {
		t.Fatalf("seed Search failed: %v", err)
	}

	retr.getSleep = 5000
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := retr.Search(ctx, SearchParams{Query: "cancelled"})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatalf("expected deadline exceeded error during throttle wait, got nil")
	}
	if gotCode := status.Code(err); gotCode != codes.DeadlineExceeded {
		t.Errorf("status.Code(err) = %v, want %v (err: %v)", gotCode, codes.DeadlineExceeded, err)
	}
	if elapsed >= time.Second {
		t.Errorf("Search took too long (%v) instead of aborting throttle wait on context cancellation", elapsed)
	}

	// Ensure cancelled throttle did not leave lastTimeRetrieved far in the future.
	throttleMu.Lock()
	reservedDelta := time.Until(lastTimeRetrieved)
	throttleMu.Unlock()
	if reservedDelta > time.Second {
		t.Errorf("cancelled throttle leaked future lastTimeRetrieved reservation: %v ahead", reservedDelta)
	}
}

func TestSearchContextCancellationDuringHTTPRequest(t *testing.T) {
	started := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		<-r.Context().Done()
	}))
	defer srv.Close()

	retr := newSearchTestRetriever(srv.URL, "token", nil)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-started
		cancel()
	}()

	_, err := retr.Search(ctx, SearchParams{Query: "slow"})
	if err == nil {
		t.Fatalf("expected canceled error, got nil")
	}
	if gotCode := status.Code(err); gotCode != codes.Canceled {
		t.Errorf("status.Code(err) = %v, want %v (err: %v)", gotCode, codes.Canceled, err)
	}
}

func TestSearchConcurrentRaceSafety(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		w.Header().Set("X-Discogs-Ratelimit", "60")
		w.Header().Set("X-Discogs-Ratelimit-Used", "1")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, `{"results":[{"id":1,"title":%q,"year":2024}]}`, q)
	}))
	defer srv.Close()

	for _, sleepMs := range []int{0, 5} {
		retr := newSearchTestRetriever(srv.URL, "concurrent-token", nil)
		retr.getSleep = sleepMs

		const workers = 16
		var wg sync.WaitGroup
		errs := make(chan error, workers)

		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				query := fmt.Sprintf("query-%d", idx)
				res, err := retr.Search(context.Background(), SearchParams{Query: query, PerPage: 5})
				if err != nil {
					errs <- err
					return
				}
				if len(res) != 1 || res[0].Title != query || res[0].Year != "2024" {
					errs <- fmt.Errorf("unexpected result for %s: %+v", query, res)
				}
			}(i)
		}

		wg.Wait()
		close(errs)
		for err := range errs {
			t.Errorf("sleepMs=%d concurrent Search error: %v", sleepMs, err)
		}
	}
}
