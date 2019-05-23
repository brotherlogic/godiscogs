package godiscogs

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	proto "github.com/golang/protobuf/proto"
)

type testFileGetter struct{}

func (httpGetter testFileGetter) Get(url string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata" + strippedURL)
	log.Printf("OPEN %v", "testdata"+strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	response.Body = blah

	// Add the header if it exists
	headers, err := os.Open("testdata" + strippedURL + ".headers")
	if err == nil {
		var t http.Header
		t = make(http.Header)
		response.Header = t

		defer headers.Close()
		scanner := bufio.NewScanner(headers)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, ":") {
				elems := strings.Split(line, ":")
				log.Printf("HEADER: '%v', '%v'", strings.TrimSpace(elems[0]), strings.TrimSpace(elems[1]))
				response.Header.Add(strings.TrimSpace(elems[0]), strings.TrimSpace(elems[1]))
			}
		}
	}

	return response, nil
}

func (httpGetter testFileGetter) Post(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata" + strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	response.Body = blah
	response.StatusCode = 204
	return response, nil
}

func (httpGetter testFileGetter) Put(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata" + strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	response.Body = blah
	return response, nil
}

func (httpGetter testFileGetter) Delete(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata" + strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	response.Body = blah
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
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}
	retr.getSleep = 0.0
	return retr
}

func TestGetImage(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}

	r, err := retr.GetRelease(4707982)

	if err != nil {
		t.Fatalf("Error getting release: %v", err)
	}

	found := false
	for _, i := range r.GetImages() {
		if i.Type == "primary" {
			if i.Uri == "" {
				t.Errorf("Unable to pick out image uri: %v", r)
			}
			found = true
		}
	}

	if !found {
		t.Errorf("No primary image: %v", r)
	}
}

func TestGetTracks(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}

	r, err := retr.GetRelease(1161277)

	if err != nil {
		t.Fatalf("Error getting release: %v", err)
	}

	if len(r.GetTracklist()) != 5 {
		t.Errorf("Wrong number of tracks retrieved: %v", len(r.GetTracklist()))
	}

	count := 0
	for _, t := range r.GetTracklist() {
		if t.TrackType == Track_TRACK {
			count++
		} else {
			log.Printf("%v", t.TrackType)
		}

		log.Printf("HERE %v", len(t.SubTracks))
		for _, st := range t.SubTracks {
			if st.TrackType == Track_TRACK {
				count++
			} else {
				log.Printf("%v", st.TrackType)
			}

		}
	}

	if count != 14 {
		t.Errorf("Bad track count: %v", count)
	}
}

func TestSellRecord(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}

	id := retr.SellRecord(2576104, 12.345, "Draft")

	if id != 567306424 {
		t.Errorf("Sale has failed: %v", id)
	}
}

func TestRemoveSale(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}

	err := retr.RemoveFromSale(805055435, 1145342)

	if err != nil {
		t.Errorf("Sale has failed: %v", err)
	}
}

func TestGetSuggestedPrice(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}

	salePrice := retr.GetSalePrice(2576104)

	if salePrice != 7.1619444 {
		t.Errorf("Failure to get sale price: %v", salePrice)
	}
}

func TestGetRateLimit(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = testFileGetter{}

	rateLimit := retr.GetRateLimit()
	if rateLimit != 60 {
		t.Errorf("Rate limit has come back wrong: %v", rateLimit)
	}
}

func TestPost(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)
	retr.getter = prodHTTPGetter{}
	retr.post("blah", "madeup")
}

func TestRetrieveLimiting(t *testing.T) {
	//Ignore the get Sleep override
	retr := NewDiscogsRetriever("token", nil)
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

func TestBuildRelease(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(1018055)
	data, _ := proto.Marshal(release)
	ioutil.WriteFile("1018055.file", data, 0644)
	release, _ = retr.GetRelease(565473)
	data, _ = proto.Marshal(release)
	ioutil.WriteFile("565473.file", data, 0644)

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

	if len(release.Formats) != 1 {
		t.Errorf("Formats has been pulled wrong: %v", release.Formats)
	}

	if len(release.Labels) != 1 {
		t.Errorf("Labels has been pulled wrong: %v", release.Labels)
	}

	if release.Labels[0].Id != 895 {
		t.Errorf("Label ID has not been pulled correctly: %v", release.Labels[0])
	}
}

func TestGetReleaseNoData(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(2425133)
	if release.Title != "Love, Love, Love, Love, Love" {
		t.Errorf("Wrong title: %v", release)
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
	v, err := retr.AddToFolder(812802, 10)

	if err != nil {
		t.Fatalf("Error running add: %v", err)
	}

	if v != 267910454 {
		t.Errorf("Error in returned instance ID: %v", v)
	}
}

func TestMoveToUncateogrized(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.MoveToFolder(10, 10, 10, 10)
}

func TestDelete(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.DeleteInstance(673768, 10866041, 280210978)
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
	body, _, _ := retr.retrieve("/releases/249504")
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

func (httpGetter testFailGetter) Put(url string, data string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}

func (httpGetter testFailGetter) Delete(url string, data string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}

func TestFailGet(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.getter = testFailGetter{}
	_, err, _ := retr.retrieve("/releases/249504")
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

func TestGetInstanceID(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	val := retr.GetInstanceID(11146958)
	if val != 261212718 {
		t.Errorf("Error in getting instance ID: %v", val)
	}
}

func TestGetCollection(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	collection := retr.GetCollection()
	if len(collection) != 3104 {
		t.Errorf("Collection retrieve is short: %v", len(collection))
	}
	found := false
	var foundRecord *Release
	log.Printf("Collection size: %v", len(collection))
	count := 0
	for _, record := range collection {
		if record.Id == 2180118 {
			count++
		}
		if record.Id == 679324 {
			found = true
			foundRecord = record
		}

		if record.Id == 0 {
			t.Errorf("Bad record found!: %v", record)
		}
	}

	if !found {
		t.Fatalf("Collection does not contain Earth Rot : %v", count)
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

func TestBoxSet(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(2370027)
	if !release.Boxset {
		t.Errorf("Boxset has not been marked as such: %v", release)
	}
}

func TestGatefold(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(9082405)
	if !release.Gatefold {
		t.Errorf("Gatefold has not been marked as such: %v", release)
	}
}

func TestGetSalePrice(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	price := retr.GetCurrentSalePrice(805377159)
	if price != 9.75 {
		t.Errorf("Price is incorrect: %v", price)
	}
}

func TestGetSaleState(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	state := retr.GetCurrentSaleState(805377159)
	if state != SaleState_FOR_SALE {
		t.Errorf("State is incorrect: %v", state)
	}
}

func TestGetSaleStateOnFail(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	state := retr.GetCurrentSaleState(805377158)
	if state != SaleState_NOT_FOR_SALE {
		t.Errorf("State is incorrect: %v", state)
	}
}

func TestGetSaleStateOnSold(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	state := retr.GetCurrentSaleState(805377157)
	if state != SaleState_SOLD {
		t.Errorf("State is incorrect: %v", state)
	}
}

func TestUpdateSalePrice(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	err := retr.UpdateSalePrice(805377159, 11403112, "Very Good Plus (VG+)", 9.50)
	if err != nil {
		t.Errorf("Update price failed!: %v", err)
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
