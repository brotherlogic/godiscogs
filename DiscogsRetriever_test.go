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

func NewTestDiscogsRetriever(token string) *DiscogsRetriever {
	retr := NewDiscogsRetriever(token)
	retr.getter = testFileGetter{}
	return retr
}

func TestRetrieveLimiting(t *testing.T) {
	retr := NewTestDiscogsRetriever("token")
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
	retr := NewTestDiscogsRetriever("token")
	release, _ := retr.GetRelease(249504)
	if release.Title != "Never Gonna Give You Up" {
		t.Errorf("Wrong title: %v", release)
	}
	if release.Artists[0].Name != "Rick Astley" {
		t.Errorf("Wrong artist name: %v", release.Artists[0].Name)
	}
}

func TestRetrieve(t *testing.T) {
	startCount := GetHTTPGetCount()
	retr := NewTestDiscogsRetriever("token")
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

func TestFailGet(t *testing.T) {
	retr := NewTestDiscogsRetriever("token")
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
	retr := NewTestDiscogsRetriever("token")
	retr.unmarshaller = testFailUnmarshaller{}
	_, err := retr.GetRelease(249504)
	if err == nil {
		t.Errorf("Error handling failed to fail on unmarshal")
	}
}

func TestGetCollection(t *testing.T) {
	retr := NewTestDiscogsRetriever("token")
	collection := retr.GetCollection()
	if len(collection) != 1918 {
		t.Errorf("Collection retrieve is short: %v", len(collection))
	}
	found := false
	for _, record := range collection {
		if record.Id == 679324 {
			found = true
		}
	}

	if !found {
		t.Errorf("Collection does not contain Earth Rot")
	}
}

func TestGetFolders(t *testing.T) {
	retr := NewTestDiscogsRetriever("token")
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

func TestMain(m *testing.M) {
	val := m.Run()
	if GetHTTPGetCount() > 1 {
		log.Printf("Too many http get calls: %v", GetHTTPGetCount())
		val = 2
	}
	os.Exit(val)
}
