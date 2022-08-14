package godiscogs

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	proto "github.com/golang/protobuf/proto"
)

type testFileGetter struct {
	count int
}

func (httpGetter *testFileGetter) Get(url string) (*http.Response, error) {
	if httpGetter.count == 0 {
		return nil, fmt.Errorf("Built to fail")
	}
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata" + strippedURL)

	log.Printf("Opened %v", "testdata"+strippedURL)
	if err != nil {
		return nil, err
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
				response.Header.Add(strings.TrimSpace(elems[0]), strings.TrimSpace(elems[1]))
			}
		}
	}

	httpGetter.count--
	return response, nil
}

func (httpGetter *testFileGetter) Post(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, _ := os.Open("testdata" + strippedURL)
	response.Body = blah
	response.StatusCode = 204
	return response, nil
}

func (httpGetter *testFileGetter) Put(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, _ := os.Open("testdata" + strippedURL)
	response.Body = blah
	return response, nil
}

func (httpGetter *testFileGetter) Delete(url string, data string) (*http.Response, error) {
	response := &http.Response{}
	strippedURL := strings.Replace(strings.Replace(url[24:], "?", "_", -1), "&", "_", -1)
	blah, _ := os.Open("testdata" + strippedURL)
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
	retr.getter = &testFileGetter{count: -1}

	retr.getSleep = 0.0
	return retr
}

func TestGetImage(t *testing.T) {
	retr := NewTestDiscogsRetriever()

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

func TestGetReleaseDate(t *testing.T) {
	retr := NewTestDiscogsRetriever()

	_, err := retr.GetRelease(2535152)

	if err != nil {
		t.Fatalf("Error getting release: %v", err)
	}

}

func TestGetTracks(t *testing.T) {
	retr := NewTestDiscogsRetriever()

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
		}

		for _, st := range t.SubTracks {
			if st.TrackType == Track_TRACK {
				count++
			}

		}
	}

	if count != 14 {
		t.Errorf("Bad track count: %v", count)
	}
}

func TestGetStats(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	stats, err := retr.GetStats(18121840)
	if err != nil {
		t.Fatalf("Bad retr: %v", err)
	}

	if stats.NumHave != 7 {
		t.Errorf("Bad stats: %+v", stats)
	}
}

func TestSellRecord(t *testing.T) {
	retr := NewTestDiscogsRetriever()

	id := retr.SellRecord(2576104, 12.345, "Draft", "blah", "blah", 12)

	if id != 567306424 {
		t.Errorf("Sale has failed: %v", id)
	}
}

func TestRemoveSale(t *testing.T) {
	retr := NewTestDiscogsRetriever()

	err := retr.ExpireSale(1079257117, 1473369, float32(4.99))

	if err != nil {
		t.Errorf("Sale has failed: %v", err)
	}
}

func TestExpireSale(t *testing.T) {
	retr := NewTestDiscogsRetriever()

	err := retr.RemoveFromSale(805055435, 1145342)

	if err != nil {
		t.Errorf("Sale has failed: %v", err)
	}
}

func TestGetSuggestedPrice(t *testing.T) {
	retr := NewTestDiscogsRetriever()

	salePrice, err := retr.GetSalePrice(2576104)
	if err != nil {
		t.Fatalf("Error in retrieve: %v", err)
	}

	if salePrice != 7.1619444 {
		t.Errorf("Failure to get sale price: %v", salePrice)
	}
}

func TestGetSuggestedPriceWithEmpty(t *testing.T) {
	retr := NewTestDiscogsRetriever()

	salePrice, err := retr.GetSalePrice(2576105)

	if err != nil || salePrice != 100 {
		t.Errorf("Failure to get sale price: %v -> %v", salePrice, err)
	}
}

func TestGetRateLimit(t *testing.T) {
	retr := NewDiscogsRetriever("token", nil)

	rateLimit := retr.GetRateLimit()
	if rateLimit != 60 {
		t.Errorf("Rate limit has come back wrong: %v", rateLimit)
	}
}

func TestPost(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.post("blah", "madeup")
}

func TestRetrieveLimiting(t *testing.T) {
	//Ignore the get Sleep override
	retr := NewDiscogsRetriever("token", nil)

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
		t.Errorf("Wrong title: %v", release.Title)
	}
	if time.Unix(release.EarliestReleaseDate, 0).In(time.UTC).Year() != 1980 {
		t.Errorf("Release has wrong date: (%v->%v) %v", release.EarliestReleaseDate, time.Unix(release.EarliestReleaseDate, 0).Year(), time.Unix(release.EarliestReleaseDate, 0))
	}
}

func TestGetOtherVersions(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(668315)
	if release.Title != "Totale's Turns (It's Now Or Never)" {
		t.Errorf("Wrong title: %v", release.Title)
	}
	if len(release.DigitalVersions) != 5 {
		t.Errorf("Wrong digital versions: %v", release.DigitalVersions)
	}
}

func TestGetOtherVersionsHuh(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, err := retr.GetRelease(372019)
	if err != nil {
		t.Fatalf("Unable to read version: %v", err)
	}
	if release.Title != "Totale's Turns (It's Now Or Never)" {
		t.Errorf("Wrong title: %v", release.Title)
	}
	if len(release.DigitalVersions) != 5 {
		t.Errorf("Wrong digital versions: %v", release.DigitalVersions)
	}
}

func TestGetEarliestReleaseDateOrdering(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(603365)
	if release.Title != "Live At The Witch Trials" {
		t.Errorf("Wrong title: %v", release.Title)
	}
	if time.Unix(release.EarliestReleaseDate, 0).In(time.UTC).Year() != 1979 {
		t.Errorf("Release has wrong date: (%v->%v) %v", release.EarliestReleaseDate, time.Unix(release.EarliestReleaseDate, 0).Year(), release.Title)
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
	body, _, err := retr.retrieve("/releases/249504")
	if !strings.Contains(string(body), "Astley") {
		t.Errorf("Error in retrieving data: %v, %v", err, string(body))
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

func TestGetInventory(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	inventory, err := retr.GetInventory()
	if err != nil {
		t.Fatalf("Bad pull: %v", err)
	}
	if len(inventory) != 322 {
		t.Errorf("Inventory pull is short: %v", len(inventory))
	}

	for _, entry := range inventory {
		if entry.Id == 551582 {
			if entry.SalePrice != 1035 {
				t.Errorf("Bad sale price: %v", entry)
			}
			if entry.DatePosted <= 0 {
				t.Errorf("Bad Date: %v", entry)
			}
		}
	}
}

func TestGetInventoryInitialFail(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.getter = &testFileGetter{count: 0}
	_, err := retr.GetInventory()
	if err == nil {
		t.Fatalf("Did not fail")
	}
}

func TestGetInventorySecondFail(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.getter = &testFileGetter{count: 1}
	_, err := retr.GetInventory()
	if err == nil {
		t.Fatalf("Did not fail")
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
	var bothHandsFree *Release
	count := 0
	for _, record := range collection {
		if record.Id == 2180118 {
			count++
		}
		if record.Id == 679324 {
			found = true
			foundRecord = record
		}

		if record.Id == 2901518 {
			bothHandsFree = record
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

	if bothHandsFree == nil {
		t.Fatalf("Both Hands Free not found")
	}

	if bothHandsFree.SleeveCondition != "Very Good Plus (VG+)" ||
		bothHandsFree.RecordCondition != "Near Mint (NM or M-)" {
		t.Errorf("Poor condition retrieve: %v", bothHandsFree)
	}

}

func TestGetInstanceInfo(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	info, err := retr.GetInstanceInfo(323005)
	if err != nil {
		t.Fatalf("Unable to pull iid: %v", err)
	}

	if info[19867048].DateAdded != 1351323375 {
		t.Errorf("Bad date: %v", info)
	}

	if info[19867048].RecordCondition != "Very Good Plus (VG+)" {
		t.Errorf("Bad condition: %+v", info[19867048])
	}

}

func TestGetInstanceInfoFailRet(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	_, err := retr.GetInstanceInfo(12)
	if err == nil {
		t.Fatalf("Bad pull did not fail: %v", err)
	}
}

func TestGetInstanceInfoFailParse(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	_, err := retr.GetInstanceInfo(323006)
	if err == nil {
		t.Fatalf("Bad pull did not fail: %v", err)
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

func TestGetSaleStateExpired(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	state := retr.GetCurrentSaleState(805377156)
	if state != SaleState_EXPIRED {
		t.Errorf("State is incorrect: %v", state)
	}
}

func TestUpdateSalePrice(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	err := retr.UpdateSalePrice(805377159, 11403112, "Very Good Plus (VG+)", "Very Good Plus (VG+)", 9.50)
	if err != nil {
		t.Errorf("Update price failed!: %v", err)
	}
}

func TestMain(m *testing.M) {
	val := m.Run()
	if GetHTTPGetCount() > 5 {
		val = 2
	}
	os.Exit(val)
}
