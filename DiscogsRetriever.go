package github.com/brotherlogic/godiscogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// DiscogsRetriever Main retriever type
type DiscogsRetriever struct {
	userAgent        string
	lastRetrieveTime int64
	userToken        string
}

// Release a release in the discogs sense
type Release struct {
	ID    int
	Title string
}

// GetRelease returns a release from the discogs system
func (r *DiscogsRetriever) GetRelease(id int) Release {
	jsonString := r.retrieve("/releases/" + strconv.Itoa(id))
	var release Release
	err := json.Unmarshal(jsonString, &release)

	if err != nil {
		panic(err)
	}

	return release
}

func (r *DiscogsRetriever) retrieve(path string) []byte {
	urlv := "https://api.discogs.com/" + path
	response, err := http.Get(urlv)

	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		return body
	}

	return make([]byte, 0)
}

func main() {
	retr := DiscogsRetriever{}
	fmt.Printf("%v\n", retr.GetRelease(249504))
}
