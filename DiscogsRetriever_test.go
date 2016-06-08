package godiscogs

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

type testFileGetter struct{}
func (httpGetter testFileGetter) Get(url string) (*http.Response, error) {
     response := &http.Response{}
     strippedURL := url[24:]
     blah,err := os.Open("testdata" + strippedURL)
     if err != nil {
     	log.Printf("Error opening test file %v", err)
     }
     response.Body = blah
     return response,nil
}

func NewTestDiscogsRetriever() *DiscogsRetriever {
     retr := NewDiscogsRetriever()
     retr.getter = testFileGetter{}
     return retr
}

func TestGetRelease(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(249504)
	if release.Title != "Never Gonna Give You Up" {
		t.Errorf("Wrong title: %v", release)
	}
}

func TestRetrieve(t *testing.T) {
     startCount := GetHTTPGetCount()
	retr := NewTestDiscogsRetriever()
	retr.getter = prodHTTPGetter{}
	body,_ := retr.retrieve("/releases/249504")
	if !strings.Contains(string(body), "Astley") {
		t.Errorf("Error in retrieving data")
	}

	endCount := GetHTTPGetCount()
	if startCount != endCount -1 {
	   t.Errorf("Retrieve did not perform a http get request: %v -> %v", startCount, endCount)
	}
}

type testFailGetter struct{}
func (httpGetter testFailGetter) Get(url string) (*http.Response, error) {
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

func TestMain(m *testing.M) {
     val := m.Run()
     if GetHTTPGetCount() > 1 {
     	log.Printf("Too many http get calls: %v", GetHTTPGetCount())
	val = 2
     }
     os.Exit(val)
}