package godiscogs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type jsonUnmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
type prodUnmarshaller struct{}

func (jsonUnmarshaller prodUnmarshaller) Unmarshal(inp []byte, v interface{}) error {
	return json.Unmarshal(inp, v)
}

var httpCount int

// GetHTTPGetCount The number of http gets performed
func GetHTTPGetCount() int {
	return httpCount
}

type httpGetter interface {
	Get(url string) (*http.Response, error)
}
type prodHTTPGetter struct{}

func (httpGetter prodHTTPGetter) Get(url string) (*http.Response, error) {
	httpCount++
	log.Printf("Retrieving %v", url)
	return http.Get(url)
}

// DiscogsRetriever Main retriever type
type DiscogsRetriever struct {
	userAgent        string
	lastRetrieveTime int64
	userToken        string
	unmarshaller     jsonUnmarshaller
	getter           httpGetter
	getSleep         int
}

// NewDiscogsRetriever Build a production retriever
func NewDiscogsRetriever(token string) *DiscogsRetriever {
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}, userToken: token, getSleep: 500}
}

// GetRelease returns a release from the discogs system
func (r *DiscogsRetriever) GetRelease(id int) (Release, error) {
	jsonString, _ := r.retrieve("/releases/" + strconv.Itoa(id) + "?token=" + r.userToken)
	var release Release
	err := r.unmarshaller.Unmarshal(jsonString, &release)

	if err != nil {
		return release, err
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
	Releases   []Release
}

// GetCollection gets all the releases in the users collection
func (r *DiscogsRetriever) GetCollection() []Release {
	jsonString, _ := r.retrieve("/users/brotherlogic/collection/folders/0/releases?per_page=100&token=" + r.userToken)

	var releases []Release
	var response CollectionResponse
	r.unmarshaller.Unmarshal(jsonString, &response)

	releases = append(releases, response.Releases...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _ = r.retrieve(response.Pagination.Urls.Next[23:])
		r.unmarshaller.Unmarshal(jsonString, &response)

		releases = append(releases, response.Releases...)
		end = response.Pagination.Pages == response.Pagination.Page
	}

	return releases
}

// FoldersResponse returned from discogs
type FoldersResponse struct {
	Pagination Pagination
	Folders    []Folder
}

// GetFolders gets all the folders for a given user
func (r *DiscogsRetriever) GetFolders() []Folder {
	jsonString, _ := r.retrieve("/users/brotherlogic/collection/folders?token=" + r.userToken)

	var folders []Folder
	var response FoldersResponse
	r.unmarshaller.Unmarshal(jsonString, &response)

	folders = append(folders, response.Folders...)

	return folders
}

func (r *DiscogsRetriever) retrieve(path string) ([]byte, error) {
	urlv := "https://api.discogs.com/" + path

	//Sleep here
	if lastTimeRetrieved.Second() > 0 {
		diff := lastTimeRetrieved.Sub(time.Now())
		if diff < time.Duration(r.getSleep)*time.Millisecond {
			time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
		}
	}

	response, err := r.getter.Get(urlv)

	if err != nil {
		return make([]byte, 0), err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	lastTimeRetrieved = time.Now()
	return body, nil
}
