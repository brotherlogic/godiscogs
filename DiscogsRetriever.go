package godiscogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// DiscogsRequests request out to discogs
	DiscogsRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "godiscogs_requests",
		Help: "The number of server requests",
	}, []string{"method", "path1"})
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

// NewDiscogsRetriever Build a production retriever
func NewDiscogsRetriever(token string, logger func(string)) *DiscogsRetriever {
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}, userToken: token, getSleep: 6000, lastRetrieveTime: time.Now().Unix(), logger: logger}
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
	Releases   []*CollectionRelease
}

// CollectionRelease returned from collection pull
type CollectionRelease struct {
	ID         int `json:"id"`
	FolderID   int `json:"folder_id"`
	InstanceID int `json:"instance_id"`
	Rating     int `json:"rating"`
	Notes      []*Note
}

// WantlistResponse returned from discogs
type WantlistResponse struct {
	Pagination Pagination
	Wants      []*Release
}

// Version a version of a master release
type Version struct {
	Released     string
	Format       string
	MajorFormats []string `json:"major_formats"`
	ID           int32
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

// PriceResponse response from get sale details
type PriceResponse struct {
	Price  Pricing
	Status string
}

// SellResponse response from selling a record
type SellResponse struct {
	ListingID int `json:"listing_id"`
}

type OrderResponse struct {
	Status   string
	Created  string
	Items    []Item
	Archived bool
}

type Item struct {
	ID    int32
	Price Pricing
}

// GetCurrentSalePrice gets the current sale price
func (r *DiscogsRetriever) GetOrder(order string) (map[int32]int32, time.Time, error) {
	rMap := make(map[int32]int32)
	tRet := time.Now()

	jsonString, _, err := r.retrieve("/marketplace/orders/" + order + "?token=" + r.userToken)
	if err != nil {
		return rMap, tRet, err
	}
	var resp OrderResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)

	tRet, err = time.Parse("2006-01-02T15:04:05-07:00", resp.Created)
	if err != nil {
		return rMap, tRet, status.Errorf(codes.Internal, "Cannot parse time (%v): %v", resp.Created, err)
	}

	if resp.Status != "Shipped" {
		// If the order was cancelled, return an empty response
		if strings.HasPrefix(resp.Status, "Cancelled") {
			return rMap, tRet, nil
		}

		// Ignore orders over two years old and have been archived
		if time.Now().Sub(tRet) > time.Hour*24*365*2 && resp.Archived {
			return rMap, tRet, nil
		}

		return rMap, tRet, status.Errorf(codes.FailedPrecondition, "Cannot process order with status %v (made on date %v -> %v)", resp.Status, tRet, resp.Archived)
	}

	for _, item := range resp.Items {
		rMap[item.ID] = int32(item.Price.Value)
	}

	return rMap, tRet, nil
}

// GetRateLimit returns the rate limit
func (r *DiscogsRetriever) GetRateLimit() int {
	_, headers, _ := r.retrieve("/releases/249504?token=" + r.userToken)
	val, _ := strconv.Atoi(headers.Get("X-Discogs-Ratelimit"))
	return val
}

func (r *DiscogsRetriever) processCollectionRelease(re *CollectionRelease) *Release {
	release := &Release{}
	for _, note := range re.Notes {
		// Media condition
		if note.FieldId == 1 {
			release.RecordCondition = note.Value
		}

		// Sleeve condition
		if note.FieldId == 2 {
			release.SleeveCondition = note.Value
		}
	}

	release.Id = int32(re.ID)
	release.InstanceId = int32(re.InstanceID)
	release.FolderId = int32(re.FolderID)
	release.Rating = int32(re.Rating)

	return release
}

func (r *DiscogsRetriever) SetNotes(iid, fid, id int32, value string) error {
	//_, err := r.post(fmt.Sprintf("/users/brotherlogic/collection/folders/%v/releases/%v/instances/%v/fields/3?value=%v&token=", fid, id, iid, value, r.userToken), "")
	return fmt.Errorf("Not finished yet")
}

// GetCollection gets all the releases in the users collection
func (r *DiscogsRetriever) GetCollection() []*Release {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/collection/folders/0/releases?per_page=100&token=" + r.userToken)

	var releases []*CollectionRelease
	response := &CollectionResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	releases = append(releases, response.Releases...)
	end := response.Pagination.Pages == response.Pagination.Page

	r.Log(fmt.Sprintf("FOUND %v PAGES", response.Pagination.Pages))

	for !end {
		jsonString, _, _ = r.retrieve(response.Pagination.Urls.Next[23:])
		newResponse := &CollectionResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		releases = append(releases, newResponse.Releases...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	procReleases := []*Release{}
	for _, re := range releases {
		procReleases = append(procReleases, r.processCollectionRelease(re))
	}

	return procReleases
}

// InventoryResponse returned from discogs
type InventoryResponse struct {
	Pagination Pagination
	Listings   []*InventoryEntry
}

// BasicRelease is the release returned from the inventory pull
type BasicRelease struct {
	ID int
}

// InventoryEntry returned from invetory pull
type InventoryEntry struct {
	Price   Pricing
	ID      int `json:"id"`
	Release BasicRelease
}

// GetInventory gets all the releases that are for sale
func (r *DiscogsRetriever) GetInventory() ([]*ForSale, error) {
	jsonString, _, err := r.retrieve("/users/brotherlogic/inventory?status=For%20Sale&per_page=100&token=" + r.userToken)
	if err != nil {
		return []*ForSale{}, err
	}

	var items []*InventoryEntry
	response := &InventoryResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	items = append(items, response.Listings...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _, err = r.retrieve(response.Pagination.Urls.Next[23:])
		if err != nil {
			return []*ForSale{}, err
		}
		newResponse := &InventoryResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		items = append(items, newResponse.Listings...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	sales := []*ForSale{}
	for _, re := range items {
		sales = append(sales, &ForSale{Id: int32(re.Release.ID), SaleId: int32(re.ID), SalePrice: int32(math.Round(float64(re.Price.Value * 100)))})
	}

	return sales, nil
}

// GetSalePrice gets the sale price for a release
func (r *DiscogsRetriever) GetSalePrice(releaseID int) (float32, error) {
	jsonString, _, err := r.retrieve("/marketplace/price_suggestions/" + strconv.Itoa(releaseID) + "?token=" + r.userToken)
	var resp map[string]Pricing
	r.unmarshaller.Unmarshal(jsonString, &resp)
	value := resp["Mint (M)"].Value

	if value == 0 && err == nil {
		return float32(100), err
	}
	return value, err
}

// SellRecord sells a given release
func (r *DiscogsRetriever) SellRecord(releaseID int, price float32, state string, condition, sleeve string) int {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"" + condition + "\", \"sleeve_condition\":\"" + sleeve + "\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"" + state + "\",\"weight\":\"auto\"}"
	databack, _ := r.post("/marketplace/listings?token="+r.userToken, data)
	var resp SellResponse
	r.unmarshaller.Unmarshal([]byte(databack), &resp)
	return resp.ListingID
}

// GetCurrentSalePrice gets the current sale price
func (r *DiscogsRetriever) GetCurrentSalePrice(saleID int) float32 {
	jsonString, _, _ := r.retrieve("/marketplace/listings/" + strconv.Itoa(saleID) + "?curr_abbr=USD&token=" + r.userToken)
	var resp PriceResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)
	return resp.Price.Value
}

// GetCurrentSaleState gets the current sale state
func (r *DiscogsRetriever) GetCurrentSaleState(saleID int) SaleState {
	jsonString, _, _ := r.retrieve("/marketplace/listings/" + strconv.Itoa(saleID) + "?curr_abbr=USD&token=" + r.userToken)
	var resp PriceResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)

	if resp.Status == "For Sale" {
		return SaleState_FOR_SALE
	} else if resp.Status == "Expired" {
		return SaleState_EXPIRED
	} else if resp.Status == "Sold" || resp.Status == "Draft" || resp.Status == "Deleted" {
		return SaleState_SOLD
	}

	r.Log(fmt.Sprintf("Unknown sale status: %v", resp.Status))
	return SaleState_NOT_FOR_SALE
}

// UpdateSalePrice updates the sale price
func (r *DiscogsRetriever) UpdateSalePrice(saleID int, releaseID int, condition, sleeve string, price float32) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"" + condition + "\", \"sleeve_condition\":\"" + sleeve + "\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"For Sale\"}"
	_, err := r.post("/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// RemoveFromSale removes the listing from sale
func (r *DiscogsRetriever) RemoveFromSale(saleID int, releaseID int) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"Near Mint (NM or M-)\", \"price\":5.00, \"status\":\"Draft\"}"
	_, err := r.post("/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// ExpireSale removes the listing from sale
func (r *DiscogsRetriever) ExpireSale(saleID int, releaseID int, price float32) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"Near Mint (NM or M-)\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"Expired\"}"
	_, err := r.post("/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// AddToWantlist adds a record to the wantlist
func (r *DiscogsRetriever) AddToWantlist(releaseID int) error {
	_, err := r.put("/users/brotherlogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
	return err
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

// MoveToFolder Moves the given release to the new folder
func (r *DiscogsRetriever) MoveToFolder(folderID int, releaseID int, instanceID int, newFolderID int) (string, error) {
	return r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "{\"folder_id\": "+strconv.Itoa(newFolderID)+"}")
}

// DeleteInstance removes a record from the collection
func (r *DiscogsRetriever) DeleteInstance(folderID int, releaseID int, instanceID int) string {
	return r.delete("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "")
}

//ReleaseBack what we get for a single release
type ReleaseBack struct {
	DateAdded  string  `json:"date_added"`
	InstanceID int32   `json:"instance_id"`
	Notes      []*Note `json:"notes"`
}

//ReleaseResponse what we get back from release
type ReleaseResponse struct {
	Pagination Pagination
	Releases   []ReleaseBack
}

type Stats struct {
	NumHave int32 `json:"num_have"`
	NumWant int32 `json:"num_want"`
}

func (r *DiscogsRetriever) GetStats(rid int32) (*Stats, error) {
	jsonString, _, err := r.retrieve(fmt.Sprintf("/releases/%v/stats?token=%v", rid, r.userToken))
	if err != nil {
		return nil, err
	}

	var stats *Stats
	r.unmarshaller.Unmarshal(jsonString, stats)

	return stats, nil
}

//InstanceInfo some basic details about the instance
type InstanceInfo struct {
	DateAdded       int64
	RecordCondition string
	SleeveCondition string
}

//GetInstanceInfo gets the info for an instance
func (r *DiscogsRetriever) GetInstanceInfo(rid int32) (map[int32]*InstanceInfo, error) {
	jsonString, _, err := r.retrieve(fmt.Sprintf("/users/brotherlogic/collection/releases/%v?token=%v", rid, r.userToken))
	mapper := make(map[int32]*InstanceInfo)
	if err != nil {
		return mapper, err
	}

	var response ReleaseResponse
	r.unmarshaller.Unmarshal(jsonString, &response)

	for _, entry := range response.Releases {
		//2015-11-30T10:54:13-08:00
		p, err := time.Parse("2006-01-02T15:04:05-07:00", entry.DateAdded)
		if err != nil {
			return mapper, err
		}
		mapper[entry.InstanceID] = &InstanceInfo{DateAdded: p.Unix()}

		for _, note := range entry.Notes {
			// Media condition
			if note.FieldId == 1 {
				mapper[entry.InstanceID].RecordCondition = note.Value
			}

			// Sleeve condition
			if note.FieldId == 2 {
				mapper[entry.InstanceID].SleeveCondition = note.Value
			}
		}

	}
	return mapper, nil
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

func (r *DiscogsRetriever) throttle() time.Duration {
	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	val := diff
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		val = time.Duration(r.getSleep)*time.Millisecond - diff
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}
	lastTimeRetrieved = time.Now()
	return val
}

var (
	// DiscogsRequests request out to discogs
	rateLimit = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "godiscogs_rate_limit",
		Help: "The number of server requests",
	}, []string{"method"})
	rateLimitUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "godiscogs_rate_limit_used",
		Help: "The number of server requests",
	}, []string{"method"})
)

func (r *DiscogsRetriever) updateRateLimit(resp *http.Response, method string) {
	limit, err := strconv.Atoi(resp.Header.Get("X-Discogs-Ratelimit"))
	if err != nil {
		r.Log(fmt.Sprintf("Unable to parse rate limit: %v", err))
	}
	rateLimit.With(prometheus.Labels{"method": method}).Set(float64(limit))

	used, err := strconv.Atoi(resp.Header.Get("X-Discogs-Ratelimit-Used"))
	if err != nil {
		r.Log(fmt.Sprintf("Unable to parse rate limit used: %v", err))
	}
	rateLimitUsed.With(prometheus.Labels{"method": method}).Set(float64(used))
}

func (r *DiscogsRetriever) retrieve(path string) ([]byte, http.Header, error) {
	urlv := "https://api.discogs.com/" + path

	r.throttle()
	DiscogsRequests.With(prometheus.Labels{"method": "GET", "path1": strings.Split(path, "/")[0]}).Inc()
	response, err := r.getter.Get(urlv)
	if err != nil {
		return make([]byte, 0), make(http.Header), err
	}
	r.updateRateLimit(response, "GET")

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 {
		r.Log(fmt.Sprintf("RETR (%v) %v -> %v", path, response.StatusCode, string(body)))
		return body, response.Header, fmt.Errorf("Bad Read: %v", string(body))
	}

	lastTimeRetrieved = time.Now()

	return body, response.Header, nil
}
