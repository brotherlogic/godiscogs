package discogsgo

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetRelease(t *testing.T) {
	retr := NewDiscogsRetriever()
	release, _ := retr.GetRelease(249504)
	if release.Title != "Never Gonna Give You Up" {
		t.Errorf("Wrong title: %v", release)
	}
}

func TestRetrieve(t *testing.T) {
	retr := NewDiscogsRetriever()
	body,_ := retr.retrieve("/releases/249504")
	if !strings.Contains(string(body), "Astley") {
		t.Errorf("Error in retrieving data")
	}
}

type testFailGetter struct{}
func (httpGetter testFailGetter) Get(url string) (*http.Response, error) {
     return nil, errors.New("Built To Fail")
}

func TestFailGet(t *testing.T) {
     retr := NewDiscogsRetriever()
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
	retr := NewDiscogsRetriever()
	retr.unmarshaller = testFailUnmarshaller{}
	_, err := retr.GetRelease(249504)
	if err == nil {
		t.Errorf("Error handling failed to fail on unmarshal")
	}
}
