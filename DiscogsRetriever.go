package godiscogs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
	return json.Unmarshal(inp, v)
}

var httpCount int

// GetHTTPGetCount The number of http gets performed
func GetHTTPGetCount() int {
	return httpCount
}

type httpGetter interface {
	Get(url string) (*http.Response, error)
	Post(url string, data string) (*http.Response, error)
}
type prodHTTPGetter struct{}

func (httpGetter prodHTTPGetter) Get(url string) (*http.Response, error) {
	httpCount++
	log.Printf("Retrieving %v", url)
	return http.Get(url)
}

func (httpGetter prodHTTPGetter) Post(url string, data string) (*http.Response, error) {
	httpCount++
	log.Printf("Posting %v", url)
	return http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
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
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}, userToken: token, getSleep: 500, lastRetrieveTime: time.Now().Unix()}
}

// GetRelease returns a release from the discogs system
func (r *DiscogsRetriever) GetRelease(id int) (Release, error) {
	jsonString, _ := r.retrieve("/releases/" + strconv.Itoa(id) + "?token=" + r.userToken)
	var release Release
	err := r.unmarshaller.Unmarshal(jsonString, &release)

	if err != nil {
		return release, err
	}

	var versions VersionsResponse
	if release.MasterId != 0 {
		// Now get the earliest release date
		jsonString, _ = r.retrieve("/masters/" + strconv.Itoa(int(release.MasterId)) + "/versions?per_page=500&token=" + r.userToken)
		r.unmarshaller.Unmarshal(jsonString, &versions)
	} else {
		tmpVersion := Version{Released: release.Released}
		versions.Versions = append(versions.Versions, tmpVersion)
	}
	bestDate := int64(-1)
	for _, version := range versions.Versions {
		log.Printf("VERSION RELEASE: %v (%v)", version.Released, strings.Count(version.Released, "-"))

		log.Printf("BEST_SO_FAR = %v", bestDate)

		if strings.Count(version.Released, "-") == 2 {
			//Check that the date is legit
			if strings.Split(version.Released, "-")[1] == "00" {
				dateV, _ := time.Parse("2006", strings.Split(version.Released, "-")[0])
				date := dateV.Unix()
				log.Printf("HERE_SPEC = %v", date)
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			} else {
				dateV, _ := time.Parse("2006-01-02", version.Released)
				date := dateV.Unix()
				log.Printf("HERE = %v (%v)", date, dateV)
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			}
		} else if strings.Count(version.Released, "-") == 0 && len(version.Released) > 0 {
			dateV, _ := time.Parse("2006", version.Released)
			date := dateV.Unix()
			log.Printf("HERE = %v (%v with %v)", date, dateV, dateV.Year())
			if bestDate < 0 || date < bestDate {
				bestDate = date
			}
		}
	}
	end := versions.Pagination.Pages == versions.Pagination.Page
	log.Printf("BEST_SO_FAR = %v", bestDate)

	for !end {
		jsonString, _ = r.retrieve(versions.Pagination.Urls.Next[23:])
		r.unmarshaller.Unmarshal(jsonString, &versions)

		for _, version := range versions.Versions {
			log.Printf("VERSION RELEASE: %v (%v)", version.Released, strings.Count(version.Released, "-"))

			if strings.Count(version.Released, "-") == 2 {
				//Check that the date is legit
				if strings.Split(version.Released, "-")[1] == "00" {
					dateV, _ := time.Parse("2006", strings.Split(version.Released, "-")[0])
					date := dateV.Unix()
					log.Printf("HERE_SPEC = %v", date)
					if bestDate < 0 || date < bestDate {
						bestDate = date
					}
				}
				dateV, _ := time.Parse("2006-02-01", version.Released)
				date := dateV.Unix()
				log.Printf("HERE = %v", date)
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			} else if strings.Count(version.Released, "-") == 0 {
				dateV, _ := time.Parse("2006", version.Released)
				date := dateV.Unix()
				log.Printf("HERE = %v", date)
				if bestDate < 0 || date < bestDate {
					bestDate = date
				}
			}
		}
		end = versions.Pagination.Pages == versions.Pagination.Page
	}

	log.Printf("BEST = %v -> %v", bestDate, time.Unix(bestDate, 0))

	if bestDate > 0 {
		release.EarliestReleaseDate = bestDate
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

// Version a version of a master release
type Version struct {
	Released string
}

// VersionsResponse returned from discogs
type VersionsResponse struct {
	Pagination Pagination
	Versions   []Version
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

// AddToFolder adds the release to the given folder
func (r *DiscogsRetriever) AddToFolder(folderID int, releaseID int) {
	r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
}

// MoveToFolder Moves the given release to the new folder
func (r *DiscogsRetriever) MoveToFolder(folderID int, releaseID int, instanceID int, newFolderID int) {
	r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "{\"folder_id\": "+strconv.Itoa(newFolderID)+"}")
}

// SetRating sets the rating on the specified releases
func (r *DiscogsRetriever) SetRating(folderID int, releaseID int, instanceID int, rating int) {
	r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "{\"rating\": "+strconv.Itoa(rating)+"}")
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
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
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

func (r *DiscogsRetriever) post(path string, data string) {
	urlv := "https://api.discogs.com/" + path

	log.Printf("POST = %v", urlv)

	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	log.Printf("DIFF = %v", diff)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		log.Printf("Sleeping for %v from %v, %v", time.Duration(r.getSleep)*time.Millisecond-diff, diff, time.Duration(r.getSleep)*time.Millisecond)
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}

	lastTimeRetrieved = time.Now()
	r.getter.Post(urlv, data)
}
