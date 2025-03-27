package godiscogs

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	pb "github.com/brotherlogic/godiscogs/proto"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type prodHTTPGetter struct{}

func (httpGetter prodHTTPGetter) Get(url string) (*http.Response, error) {
	httpCount++
	return http.Get(url)
}

func (httpGetter prodHTTPGetter) Post(url string, data string) (*http.Response, error) {
	httpCount++
	return http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
}

func (httpGetter prodHTTPGetter) Put(url string, data string) (*http.Response, error) {
	httpCount++
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(req)
}

func (httpGetter prodHTTPGetter) Delete(url string, data string) (*http.Response, error) {
	httpCount++
	req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte(data)))
	return http.DefaultClient.Do(req)
}

// Log out a value to the log function
func (r *DiscogsRetriever) Log(ctx context.Context, text string) {
	if r.logger != nil {
		r.logger(ctx, text)
	}
}

func (r *DiscogsRetriever) setTrack(ctx context.Context, t *pb.Track) {
	switch t.Type_ {
	case "track":
		t.TrackType = pb.Track_TRACK
	case "heading":
		t.TrackType = pb.Track_HEADING
	case "index":
		// Pass
	default:
		r.Log(ctx, fmt.Sprintf("Unknown type: %v", t.Type_))
	}
}

// GetRelease returns a release from the discogs system
func (r *DiscogsRetriever) GetRelease(ctx context.Context, id int32) (*pb.Release, error) {
	jsonString, _, err := r.retrieve(ctx, "/releases/"+strconv.Itoa(int(id))+"?token="+r.userToken)
	if err != nil {
		return nil, err
	}
	var release *pb.Release
	err = r.unmarshaller.Unmarshal(jsonString, &release)

	if err != nil {
		return release, err
	}

	// Work the tracks
	for _, t := range release.GetTracklist() {
		r.setTrack(ctx, t)
		for _, st := range t.SubTracks {
			r.setTrack(ctx, st)
		}
	}

	var versions VersionsResponse
	if release.MasterId != 0 {
		// Now get the earliest release date
		jsonString, _, err = r.retrieve(ctx, "/masters/"+strconv.Itoa(int(release.MasterId))+"/versions?per_page=500&token="+r.userToken)
		if err != nil {
			return nil, err
		}
		r.unmarshaller.Unmarshal(jsonString, &versions)
	} else {
		tmpVersion := Version{Released: release.Released, Format: ""}
		versions.Versions = append(versions.Versions, tmpVersion)
	}
	bestDate := int64(0)
	release.DigitalVersions = []int32{}
	release.OtherVersions = []int32{}
	for _, version := range versions.Versions {
		if version.ID > 0 {
			release.OtherVersions = append(release.OtherVersions, version.ID)

			official := true
			if strings.Contains(version.Format, "Unofficial") {
				official = false
			}

			if strings.Contains(version.Format, "CD") || strings.Contains(version.Format, "File") {
				if official {
					release.DigitalVersions = append(release.DigitalVersions, version.ID)
				}
			} else {
				for _, format := range version.MajorFormats {
					if strings.Contains(format, "CD") || strings.Contains(format, "File") {
						if official {
							release.DigitalVersions = append(release.DigitalVersions, version.ID)
							break
						}
					}
				}
			}
			if version.Released != "0" {
				if strings.Count(version.Released, "-") == 2 {
					//Check that the date is legit
					if strings.Split(version.Released, "-")[1] == "00" {
						dateP, _ := time.Parse("2006", strings.Split(version.Released, "-")[0])
						dateV := time.Date(dateP.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
						date := dateV.Unix()
						if bestDate == 0 || date < bestDate {
							bestDate = date
						}
					} else {
						dateV, err := time.Parse("2006-01-02", version.Released)
						if err == nil {
							date := dateV.Unix()
							if bestDate == 0 || date < bestDate {
								bestDate = date
							}
						}
					}
				} else if strings.Count(version.Released, "-") == 0 && len(version.Released) > 0 {
					dateP, _ := time.Parse("2006", version.Released)
					dateV := time.Date(dateP.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
					date := dateV.Unix()
					if bestDate == 0 || date < bestDate {
						bestDate = date
					}
				} else {
					dateV, err := time.Parse("01 Feb 2006", version.Released)
					if err == nil {
						date := dateV.Unix()
						if bestDate < 0 || date < bestDate {
							bestDate = date
						}
					}
				}
			}
		}
	}
	end := versions.Pagination.Pages == versions.Pagination.Page

	for !end {
		jsonString, _, err = r.retrieve(ctx, versions.Pagination.Urls.Next[23:])
		if err != nil {
			return nil, err
		}
		r.unmarshaller.Unmarshal(jsonString, &versions)

		for _, version := range versions.Versions {
			if strings.Contains(version.Format, "CD") || strings.Contains(version.Format, "File") {
				release.DigitalVersions = append(release.DigitalVersions, version.ID)
			}

			if version.Released != "0" {
				if strings.Count(version.Released, "-") == 2 {
					//Check that the date is legit
					if strings.Split(version.Released, "-")[1] == "00" {
						dateP, _ := time.Parse("2006", strings.Split(version.Released, "-")[0])
						dateV := time.Date(dateP.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
						date := dateV.Unix()
						if bestDate < 0 || date < bestDate {
							bestDate = date
						}
					}
					dateV, _ := time.Parse("2006-02-01", version.Released)
					date := dateV.Unix()
					if bestDate == 0 || date < bestDate {
						bestDate = date
					}
				} else if strings.Count(version.Released, "-") == 0 {
					dateP, _ := time.Parse("2006", version.Released)
					dateV := time.Date(dateP.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
					date := dateV.Unix()
					if bestDate == 0 || date < bestDate {
						bestDate = date
					}
				} else {
					dateV, err := time.Parse("01 Feb 2006", version.Released)
					if err == nil {
						date := dateV.Unix()
						if bestDate == 0 || date < bestDate {
							bestDate = date
						}
					}
				}
			}

		}
		end = versions.Pagination.Pages == versions.Pagination.Page
	}

	if bestDate != 0 {
		release.EarliestReleaseDate = bestDate
	}

	//Set boolean fields
	for _, format := range release.GetFormats() {
		if format.Text == "Gatefold" {
			release.Gatefold = true
		} else if format.Text == "Boxset" || format.Name == "Box Set" {
			release.Boxset = true
		}
	}

	return release, nil
}

// GetWantlist returns the wantlist for the given user
func (r *DiscogsRetriever) GetWantlist(ctx context.Context) ([]*pb.Release, error) {
	jsonString, _, err := r.retrieve(ctx, "users/BrotherLogic/wants?per_page=100&token="+r.userToken)

	if err != nil {
		return nil, err
	}

	var releases []*pb.Release
	response := &WantlistResponse{}
	r.unmarshaller.Unmarshal(jsonString, &response)

	releases = append(releases, response.Wants...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _, _ = r.retrieve(ctx, response.Pagination.Urls.Next[23:])
		newResponse := &WantlistResponse{}
		err := r.unmarshaller.Unmarshal(jsonString, &newResponse)
		if err != nil {
			return nil, err
		}

		releases = append(releases, newResponse.Wants...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	return releases, nil
}

// GetInstanceID Gets the instance ID for this release
func (r *DiscogsRetriever) GetInstanceID(ctx context.Context, releaseID int) int32 {
	jsonString, _, _ := r.retrieve(ctx, "/users/BrotherLogic/collection/releases/"+strconv.Itoa(releaseID)+"?token="+r.userToken)
	var response CollectionResponse
	r.unmarshaller.Unmarshal(jsonString, &response)
	if len(response.Releases) > 0 {
		return int32(response.Releases[0].InstanceID)
	}

	return -1
}

// AddToFolder adds the release to the given folder
func (r *DiscogsRetriever) AddToFolder(ctx context.Context, folderID int32, releaseID int32) (int, error) {
	jsonString, _ := r.post(ctx, "/users/BrotherLogic/collection/folders/"+strconv.Itoa(int(folderID))+"/releases/"+strconv.Itoa(int(releaseID))+"?token="+r.userToken, "")
	var response AddToFolderResponse
	err := r.unmarshaller.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return -1, err
	}
	return response.InstanceID, nil
}

func (r *DiscogsRetriever) post(ctx context.Context, path string, data string) (string, error) {
	path = strings.TrimPrefix(path, "/")
	urlv := "https://api.discogs.com/" + path
	r.Log(ctx, fmt.Sprintf("Posting %v to %v", data, urlv))

	//Sleep here
	tv := r.throttle()
	t := time.Now()
	DiscogsRequests.With(prometheus.Labels{"method": "POST", "path1": strings.Split(path, "/")[0]}).Inc()
	response, err := r.getter.Post(urlv, data)
	RequestLatency.With(prometheus.Labels{"method": "POST", "path1": strings.Split(path, "/")[0]}).Observe(float64(time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return fmt.Sprintf("POST ERROR ON RUN: %v", err), err
	}
	r.updateRateLimit(ctx, response, "POST")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	r.Log(ctx, fmt.Sprintf("POST: %v", response.StatusCode))

	// This means we're trying to modify something we shouldn't
	if response.StatusCode == 422 {
		return "", status.Error(codes.FailedPrecondition, fmt.Sprintf("POST ERROR (STATUS CODE): %v, %v", response.StatusCode, string(body)))
	}

	if response.StatusCode == 429 {
		return "", status.Errorf(codes.ResourceExhausted, "%v / %v", response.Header.Get("X-Discogs-Ratelimit"), response.Header.Get("X-Discogs-Ratelimit-Used"))
	}

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 {
		return fmt.Sprintf("RETR %v -> %v given %v", response.StatusCode, string(body), path), fmt.Errorf("POST ERROR (STATUS CODE): %v, %v (%v, %v) throttled %v", response.StatusCode, string(body), response.Header.Get("X-Discogs-Ratelimit"), response.Header.Get("X-Discogs-Ratelimit-Used"), tv)
	}

	// Return Unavailable on a 502
	if response.StatusCode == 502 {
		return "", status.Error(codes.Unavailable, fmt.Sprintf("%v", string(body)))
	}

	return string(body), nil
}

func (r *DiscogsRetriever) delete(ctx context.Context, path string, data string) error {
	urlv := "https://api.discogs.com/" + path

	r.Log(ctx, fmt.Sprintf("Deleting %v", urlv))

	//Sleep here
	r.throttle()
	t := time.Now()
	DiscogsRequests.With(prometheus.Labels{"method": "DELETE", "path1": strings.Split(path, "/")[0]}).Inc()
	response, err := r.getter.Delete(urlv, data)
	RequestLatency.With(prometheus.Labels{"method": "DELETE", "path1": strings.Split(path, "/")[0]}).Observe(float64(time.Now().Sub(t).Milliseconds()))

	if err != nil {
		return fmt.Errorf("POST ERROR: %v", err)
	}
	r.updateRateLimit(ctx, response, "DELETE")
	defer response.Body.Close()

	r.logger(ctx, fmt.Sprintf("DELETE %v -> %v", path, response.StatusCode))

	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 204 {
		if response.StatusCode == 404 {
			return status.Errorf(codes.NotFound, "error on delete ( => %v): %v", response.StatusCode, string(body))
		}
		return fmt.Errorf("error on delete ( => %v): %v", response.StatusCode, string(body))
	}
	return nil
}

func (r *DiscogsRetriever) put(ctx context.Context, path string, data string) ([]byte, error) {
	urlv := "https://api.discogs.com/" + path

	//Sleep here
	r.throttle()

	t := time.Now()
	DiscogsRequests.With(prometheus.Labels{"method": "PUT", "path1": strings.Split(path, "/")[0]}).Inc()
	response, err := r.getter.Put(urlv, data)
	RequestLatency.With(prometheus.Labels{"method": "PUT", "path1": strings.Split(path, "/")[0]}).Observe(float64(time.Now().Sub(t).Milliseconds()))

	if err != nil {
		return make([]byte, 0), err
	}
	r.updateRateLimit(ctx, response, "PUT")

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 && response.StatusCode != 201 {
		r.Log(ctx, fmt.Sprintf("RETR %v -> %v given %v", response.StatusCode, string(body), path))
	}
	if err != nil {
		return make([]byte, 0), err
	}
	return body, nil
}
