package godiscogs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
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

//Log out a value to the log function
func (r *DiscogsRetriever) Log(text string) {
	if r.logger != nil {
		r.logger(text)
	}
}

// GetRelease returns a release from the discogs system
func (r *DiscogsRetriever) GetRelease(id int32) (*Release, error) {
	jsonString, _, _ := r.retrieve("/releases/" + strconv.Itoa(int(id)) + "?token=" + r.userToken)
	var release *Release
	err := r.unmarshaller.Unmarshal(jsonString, &release)

	if err != nil {
		return release, err
	}

	var versions VersionsResponse
	if release.MasterId != 0 {
		// Now get the earliest release date
		jsonString, _, _ = r.retrieve("/masters/" + strconv.Itoa(int(release.MasterId)) + "/versions?per_page=500&token=" + r.userToken)
		r.unmarshaller.Unmarshal(jsonString, &versions)
	} else {
		tmpVersion := Version{Released: release.Released}
		versions.Versions = append(versions.Versions, tmpVersion)
	}
	bestDate := int64(-1)
	for _, version := range versions.Versions {
		if strings.Count(version.Released, "-") == 2 {
			//Check that the date is legit
			if strings.Split(version.Released, "-")[1] == "00" {
				dateV, _ := time.Parse("2006", strings.Split(version.Released, "-")[0])
				date := dateV.Unix()
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			} else {
				dateV, _ := time.Parse("2006-01-02", version.Released)
				date := dateV.Unix()
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			}
		} else if strings.Count(version.Released, "-") == 0 && len(version.Released) > 0 {
			dateV, _ := time.Parse("2006", version.Released)
			date := dateV.Unix()
			if bestDate < 0 || date < bestDate {
				bestDate = date
			}
		}
	}
	end := versions.Pagination.Pages == versions.Pagination.Page

	for !end {
		jsonString, _, _ = r.retrieve(versions.Pagination.Urls.Next[23:])
		r.unmarshaller.Unmarshal(jsonString, &versions)

		for _, version := range versions.Versions {

			if strings.Count(version.Released, "-") == 2 {
				//Check that the date is legit
				if strings.Split(version.Released, "-")[1] == "00" {
					dateV, _ := time.Parse("2006", strings.Split(version.Released, "-")[0])
					date := dateV.Unix()
					if bestDate < 0 || date < bestDate {
						bestDate = date
					}
				}
				dateV, _ := time.Parse("2006-02-01", version.Released)
				date := dateV.Unix()
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			} else if strings.Count(version.Released, "-") == 0 {
				dateV, _ := time.Parse("2006", version.Released)
				date := dateV.Unix()
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			}
		}
		end = versions.Pagination.Pages == versions.Pagination.Page
	}

	if bestDate > 0 {
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
func (r *DiscogsRetriever) GetWantlist() ([]*Release, error) {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/wants?per_page=100&token=" + r.userToken)

	var releases []*Release
	var response WantlistResponse
	r.unmarshaller.Unmarshal(jsonString, &response)

	releases = append(releases, response.Wants...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _, _ = r.retrieve(response.Pagination.Urls.Next[23:])
		r.unmarshaller.Unmarshal(jsonString, &response)

		releases = append(releases, response.Wants...)
		end = response.Pagination.Pages == response.Pagination.Page
	}

	return releases, nil
}

// GetInstanceID Gets the instance ID for this release
func (r *DiscogsRetriever) GetInstanceID(releaseID int) int32 {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/collection/releases/" + strconv.Itoa(releaseID) + "?token=" + r.userToken)
	var response CollectionResponse
	r.unmarshaller.Unmarshal(jsonString, &response)
	if len(response.Releases) > 0 {
		return response.Releases[0].InstanceId
	}

	return -1
}

// AddToFolder adds the release to the given folder
func (r *DiscogsRetriever) AddToFolder(folderID int32, releaseID int32) (int, error) {
	jsonString, _ := r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(int(folderID))+"/releases/"+strconv.Itoa(int(releaseID))+"?token="+r.userToken, "")
	var response AddToFolderResponse
	err := r.unmarshaller.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return -1, err
	}
	return response.InstanceID, nil
}

func (r *DiscogsRetriever) post(path string, data string) (string, error) {
	urlv := "https://api.discogs.com/" + path
	r.Log(fmt.Sprintf("Posting %v to %v", data, urlv))

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	lastTimeRetrieved = time.Now()
	response, err := r.getter.Post(urlv, data)
	if err != nil {
		return fmt.Sprintf("POST ERROR ON RUN: %v", err), err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 {
		return fmt.Sprintf("RETR %v -> %v given %v", response.StatusCode, string(body), path), fmt.Errorf("POST ERROR (STATUS CODE): %v", response.StatusCode)
	}

	return string(body), nil
}

func (r *DiscogsRetriever) delete(path string, data string) string {
	urlv := "https://api.discogs.com/" + path

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	lastTimeRetrieved = time.Now()
	response, err := r.getter.Delete(urlv, data)
	if err != nil {
		return fmt.Sprintf("POST ERROR: %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func (r *DiscogsRetriever) put(path string, data string) ([]byte, error) {
	urlv := "https://api.discogs.com/" + path

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	lastTimeRetrieved = time.Now()
	response, err := r.getter.Put(urlv, data)
	if err != nil {
		return make([]byte, 0), err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 && response.StatusCode != 201 {
		r.Log(fmt.Sprintf("RETR %v -> %v given %v", response.StatusCode, string(body), path))
	}
	if err != nil {
		return make([]byte, 0), err
	}
	return body, nil
}