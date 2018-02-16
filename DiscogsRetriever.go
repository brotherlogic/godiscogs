package godiscogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type jsonUnmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
type prodUnmarshaller struct{}

func (jsonUnmarshaller prodUnmarshaller) Unmarshal(inp []byte, v interface{}) error {
	err := json.Unmarshal(inp, v)
	return err
}

var httpCount int

// GetHTTPGetCount The number of http gets performed
func GetHTTPGetCount() int {
	return httpCount
}

type httpGetter interface {
	Get(url string) (*http.Response, error)
	Post(url string, data string) (*http.Response, error)
	Put(url string, data string) (*http.Response, error)
	Delete(url string, data string) (*http.Response, error)
}
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

// DiscogsRetriever Main retriever type
type DiscogsRetriever struct {
	userAgent        string
	lastRetrieveTime int64
	userToken        string
	unmarshaller     jsonUnmarshaller
	getter           httpGetter
	getSleep         int
	logger           func(string)
}

//Log out a value to the log function
func (r *DiscogsRetriever) Log(text string) {
	if r.logger != nil {
		r.logger(text)
	}
}

// NewDiscogsRetriever Build a production retriever
func NewDiscogsRetriever(token string, logger func(string)) *DiscogsRetriever {
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}, userToken: token, getSleep: 1500, lastRetrieveTime: time.Now().Unix(), logger: logger}
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

var lastTimeRetrieved time.Time

// Urls list of urls in pagination
type Urls struct {
	Next string
}

// Pagination the pagination structure
type Pagination struct {
	Pages int
	Page  int
	Urls  Urls
}

// CollectionResponse returned from discogs
type CollectionResponse struct {
	Pagination Pagination
	Releases   []*Release
}

// WantlistResponse returned from discogs
type WantlistResponse struct {
	Pagination Pagination
	Wants      []*Release
}

// Version a version of a master release
type Version struct {
	Released string
}

// VersionsResponse returned from discogs
type VersionsResponse struct {
	Pagination Pagination
	Versions   []Version
}

// Pricing the single price
type Pricing struct {
	Currency string
	Value    float32
}

// GetRateLimit returns the rate limit
func (r *DiscogsRetriever) GetRateLimit() int {
	_, headers, _ := r.retrieve("/releases/249504?token=" + r.userToken)
	val, _ := strconv.Atoi(headers.Get("X-Discogs-Ratelimit"))
	return val
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

// GetCollection gets all the releases in the users collection
func (r *DiscogsRetriever) GetCollection() []*Release {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/collection/folders/0/releases?per_page=100&token=" + r.userToken)

	var releases []*Release
	response := &CollectionResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	releases = append(releases, response.Releases...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _, _ = r.retrieve(response.Pagination.Urls.Next[23:])
		newResponse := &CollectionResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		releases = append(releases, newResponse.Releases...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	return releases
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

// GetSalePrice gets the sale price for a release
func (r *DiscogsRetriever) GetSalePrice(releaseID int) float32 {
	jsonString, _, _ := r.retrieve("/marketplace/price_suggestions/" + strconv.Itoa(releaseID) + "?token=" + r.userToken)
	var resp map[string]Pricing
	r.unmarshaller.Unmarshal(jsonString, &resp)
	return resp["Very Good Plus (VG+)"].Value
}

// SellRecord sells a given release
func (r *DiscogsRetriever) SellRecord(releaseID int, price float32, state string) {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"Very Good Plus (VG+)\", \"sleeve_condition\":\"Very Good Plus (VG+)\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"" + state + "\",\"weight\":\"auto\", \"allow_offers\":\"true\"}"
	r.post("/marketplace/listings?token="+r.userToken, data)
}

// AddToWantlist adds a record to the wantlist
func (r *DiscogsRetriever) AddToWantlist(releaseID int) {
	r.put("/users/brotherlogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
}

// RemoveFromWantlist adds a record to the wantlist
func (r *DiscogsRetriever) RemoveFromWantlist(releaseID int) {
	r.delete("/users/brotherlogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
}

//AddToFolderResponse the response back from an add request
type AddToFolderResponse struct {
	InstanceID  int `json:"instance_id"`
	ResourceURL string
	Simple      int
}

// AddToFolder adds the release to the given folder
func (r *DiscogsRetriever) AddToFolder(folderID int32, releaseID int32) (int, error) {
	jsonString := r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(int(folderID))+"/releases/"+strconv.Itoa(int(releaseID))+"?token="+r.userToken, "")
	var response AddToFolderResponse
	err := r.unmarshaller.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return -1, err
	}
	return response.InstanceID, nil
}

// MoveToFolder Moves the given release to the new folder
func (r *DiscogsRetriever) MoveToFolder(folderID int, releaseID int, instanceID int, newFolderID int) string {
	return r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "{\"folder_id\": "+strconv.Itoa(newFolderID)+"}")
}

// DeleteInstance removes a record from the collection
func (r *DiscogsRetriever) DeleteInstance(folderID int, releaseID int, instanceID int) string {
	return r.delete("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "")
}

// FoldersResponse returned from discogs
type FoldersResponse struct {
	Pagination Pagination
	Folders    []Folder
}

// GetFolders gets all the folders for a given user
func (r *DiscogsRetriever) GetFolders() []Folder {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/collection/folders?token=" + r.userToken)

	var folders []Folder
	var response FoldersResponse
	r.unmarshaller.Unmarshal(jsonString, &response)

	folders = append(folders, response.Folders...)

	return folders
}

func (r *DiscogsRetriever) retrieve(path string) ([]byte, http.Header, error) {
	urlv := "https://api.discogs.com/" + path

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		r.Log(fmt.Sprintf("GET (%v) Sleeping for %v", urlv, time.Duration(r.getSleep)*time.Millisecond-diff))
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	response, err := r.getter.Get(urlv)

	if err != nil {
		return make([]byte, 0), make(http.Header), err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	lastTimeRetrieved = time.Now()
	return body, response.Header, nil
}

func (r *DiscogsRetriever) post(path string, data string) string {
	urlv := "https://api.discogs.com/" + path
	r.Log(fmt.Sprintf("Posting %v to %v", data, urlv))

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		r.Log(fmt.Sprintf("Post Sleeping for %v", time.Duration(r.getSleep)*time.Millisecond-diff))
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	lastTimeRetrieved = time.Now()
	response, err := r.getter.Post(urlv, data)
	if err != nil {
		return fmt.Sprintf("POST ERROR: %v", err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func (r *DiscogsRetriever) delete(path string, data string) string {
	urlv := "https://api.discogs.com/" + path

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		r.Log(fmt.Sprintf("Delete Sleeping for %v", time.Duration(r.getSleep)*time.Millisecond-diff))
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
		r.Log(fmt.Sprintf("Put Sleeping for %v", time.Duration(r.getSleep)*time.Millisecond-diff))
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	lastTimeRetrieved = time.Now()
	response, err := r.getter.Put(urlv, data)
	if err != nil {
		return make([]byte, 0), err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]byte, 0), err
	}
	return body, nil
}
