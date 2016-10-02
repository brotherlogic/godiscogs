package godiscogs

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type testFileGetter struct{}

func (httpGetter testFileGetter) Get(url string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata" + strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	response.Body = blah
	return response, nil
}

func (httpGetter testFileGetter) Post(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	return response, nil
}

func (httpGetter testFileGetter) Delete(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	return response, nil
}

func TestGetWantlist(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	wantlist, err := retr.GetWantlist()

	if err != nil {
		t.Errorf("Error retrieving want list: %v", err)
	}

	if len(wantlist) == 0 {
		t.Errorf("Wantlist has come back empty")
	}
}

func NewTestDiscogsRetriever() *DiscogsRetriever {
	retr := NewDiscogsRetriever("token")
	retr.getter = testFileGetter{}
	retr.getSleep = 0.0
	return retr
}

func TestPost(t *testing.T) {
	retr := NewDiscogsRetriever("token")
	retr.getter = prodHTTPGetter{}
	retr.post("blah", "madeup")
}

func TestRetrieveLimiting(t *testing.T) {
	//Ignore the get Sleep override
	retr := NewDiscogsRetriever("token")
	retr.getter = testFileGetter{}
	start := time.Now()

	for i := 0; i < 3; i++ {
		retr.retrieve("/releases/249504")
	}
	end := time.Now()

	// 6 requests should take more than 3 seconds
	if end.Sub(start) < time.Second {
		t.Errorf("Danger of being throttled by discogs API; 6 requests took %v ms", end.Sub(start).Seconds())
	}
}

func TestGetRelease(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(249504)
	if release.Title != "Never Gonna Give You Up" {
		t.Errorf("Wrong title: %v", release)
	}
	if release.Artists[0].Name != "Rick Astley" {
		t.Errorf("Wrong artist name: %v", release.Artists[0].Name)
	}
	if !strings.Contains(release.Images[0].Uri, "https") {
		t.Errorf("Image has not been retrieved: %v", release)
	}
}

func TestGetEarliestReleaseDate(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(668315)
	if release.Title != "Totale's Turns (It's Now Or Never)" {
		t.Errorf("Wrong title: %v", release)
	}
	if time.Unix(release.EarliestReleaseDate, 0).In(time.UTC).Year() != 1980 {
		t.Errorf("Release has wrong date: (%v->%v) %v", release.EarliestReleaseDate, time.Unix(release.EarliestReleaseDate, 0).Year(), release)
	}
}

func TestGetEarliestReleaseDateOrdering(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(603365)
	if release.Title != "Live At The Witch Trials" {
		t.Errorf("Wrong title: %v", release)
	}
	if time.Unix(release.EarliestReleaseDate, 0).In(time.UTC).Year() != 1979 {
		t.Errorf("Release has wrong date: (%v->%v) %v", release.EarliestReleaseDate, time.Unix(release.EarliestReleaseDate, 0).Year(), release)
	}
}

func TestAddToFolder(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.AddToFolder(10, 10)
}

func TestMoveToUncateogrized(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.MoveToFolder(10, 10, 10, 10)
}

func TestAddToWantlist(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.AddToWantlist(100)
}

func TestRemoveFromWantlist(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.RemoveFromWantlist(100)
}

func TestRetrieve(t *testing.T) {
	startCount := GetHTTPGetCount()
	retr := NewTestDiscogsRetriever()
	retr.getter = prodHTTPGetter{}
	body, _ := retr.retrieve("/releases/249504")
	if !strings.Contains(string(body), "Astley") {
		t.Errorf("Error in retrieving data")
	}

	endCount := GetHTTPGetCount()
	if startCount != endCount-1 {
		t.Errorf("Retrieve did not perform a http get request: %v -> %v", startCount, endCount)
	}
}

type testFailGetter struct{}

func (httpGetter testFailGetter) Get(url string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}

func (httpGetter testFailGetter) Post(url string, data string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}

func (httpGetter testFailGetter) Delete(url string, data string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}

func TestFailGet(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.getter = testFailGetter{}
	_, err := retr.retrieve("/releases/249504")
	if err == nil {
		t.Errorf("Get did not throw an error")
	}
}

type testFailUnmarshaller struct{}

func (jsonUnmarshaller testFailUnmarshaller) Unmarshal(inp []byte, v interface{}) error {
	return errors.New("Built To Fail")
}

func TestFailMarshal(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.unmarshaller = testFailUnmarshaller{}
	_, err := retr.GetRelease(249504)
	if err == nil {
		t.Errorf("Error handling failed to fail on unmarshal")
	}
}

func TestGetCollection(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	collection := retr.GetCollection()
	if len(collection) != 2371 {
		t.Errorf("Collection retrieve is short: %v", len(collection))
	}
	found := false
	var foundRecord Release
	for _, record := range collection {
		if record.Id == 679324 {
			found = true
			foundRecord = record
		}
	}

	if !found {
		t.Errorf("Collection does not contain Earth Rot")
	}

	if foundRecord.FolderId != 242017 {
		t.Errorf("Earth Rot is not in the right folder: %v", foundRecord.FolderId)
	}

	if foundRecord.InstanceId != 19867228 {
		t.Errorf("Instance ID is not right: %v", foundRecord.InstanceId)
	}

	if foundRecord.Rating != 5 {
		t.Errorf("Rating is not right: %v", foundRecord)
	}
}

func TestGetFolders(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	folders := retr.GetFolders()
	if len(folders) == 0 {
		t.Errorf("Folder retrieve is short: %v", len(folders))
	}
	found := false
	for _, folder := range folders {
		if folder.Name == "ListeningPile" {
			found = true
		}
	}

	if !found {
		t.Errorf("Collection does not have ListeningPile: %v", folders)
	}
}

func TestPostTiming(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.getSleep = 200

	start := time.Now()
	for i := 0; i < 3; i++ {
		//Insert 200 ms of sleep here
		if i == 2 {
			time.Sleep(time.Millisecond * 200)
		}

		retr.post("madeup", "")
	}
	end := time.Now()
	diff := end.Sub(start) / time.Millisecond

	if diff > 700 || diff < 500 {
		t.Errorf("Timing on posts is quite wrong: %v", diff)
	}
}

func TestMain(m *testing.M) {
	val := m.Run()
	if GetHTTPGetCount() > 2 {
		log.Printf("Too many http get calls: %v", GetHTTPGetCount())
		val = 2
	}
	os.Exit(val)
}
