package godiscogs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	pb "github.com/brotherlogic/godiscogs/proto"

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

	// DiscogsRequests request out to discogs
	RequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "godiscogs_request_latency",
		Help:    "The number of server requests",
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2000, 4000, 8000, 16000, 32000, 64000, 128000, 256000, 1024000},
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
	logger           func(context.Context, string)
}

// NewDiscogsRetriever Build a production retriever
func NewDiscogsRetriever(token string, logger func(context.Context, string)) *DiscogsRetriever {
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}, userToken: token, getSleep: 500, lastRetrieveTime: time.Now().Unix(), logger: logger}
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
	Notes      []*pb.Note
}

// WantlistResponse returned from discogs
type WantlistResponse struct {
	Pagination Pagination
	Wants      []*pb.Release
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
	ListingID int64 `json:"listing_id"`
}

type OrderResponse struct {
	Status   string
	Created  string
	Items    []Item
	Archived bool
}

type Item struct {
	ID    int64
	Price Pricing
}

// GetCurrentSalePrice gets the current sale price
func (r *DiscogsRetriever) GetOrder(ctx context.Context, order string) (map[int64]int32, time.Time, error) {
	rMap := make(map[int64]int32)
	tRet := time.Now()

	jsonString, _, err := r.retrieve(ctx, "/marketplace/orders/"+order+"?token="+r.userToken)
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
			return rMap, tRet, status.Errorf(codes.DataLoss, "Order was cancelled")
		}

		if strings.HasPrefix(resp.Status, "Payment Received") {
			return rMap, tRet, status.Errorf(codes.FailedPrecondition, "This order needs to be processed fully")
		}

		// Ignore orders over two years old and have been archived
		if time.Now().Sub(tRet) > time.Hour*24*90 && resp.Archived {
			return rMap, tRet, nil
		}

		if strings.HasPrefix(resp.Status, "Merged") {
			return rMap, tRet, status.Errorf(codes.DataLoss, "Order was merged")
		}

		r.Log(ctx, fmt.Sprintf("Unable to process order: %v", resp))

		return rMap, tRet, status.Errorf(codes.FailedPrecondition, "Cannot process order with status of %v (made on date %v -> %v)", resp.Status, tRet, resp.Archived)
	}

	for _, item := range resp.Items {
		rMap[item.ID] = int32(item.Price.Value * 100)
	}

	return rMap, tRet, nil
}

// GetRateLimit returns the rate limit
func (r *DiscogsRetriever) GetRateLimit(ctx context.Context) int {
	_, headers, _ := r.retrieve(ctx, "/releases/249504?token="+r.userToken)
	val, _ := strconv.Atoi(headers.Get("X-Discogs-Ratelimit"))
	return val
}

func (r *DiscogsRetriever) processCollectionRelease(re *CollectionRelease) *pb.Release {
	release := &pb.Release{}
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
	//_, err := r.post(fmt.Sprintf("/users/BrotherLogic/collection/folders/%v/releases/%v/instances/%v/fields/3?value=%v&token=", fid, id, iid, value, r.userToken), "")
	return fmt.Errorf("Not finished yet")
}

// GetCollection gets all the releases in the users collection
func (r *DiscogsRetriever) GetCollection(ctx context.Context) []*pb.Release {
	jsonString, _, _ := r.retrieve(ctx, "/users/BrotherLogic/collection/folders/0/releases?per_page=100&token="+r.userToken)

	var releases []*CollectionRelease
	response := &CollectionResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	releases = append(releases, response.Releases...)
	end := response.Pagination.Pages == response.Pagination.Page

	r.Log(ctx, fmt.Sprintf("FOUND %v PAGES", response.Pagination.Pages))

	for !end {
		jsonString, _, _ = r.retrieve(ctx, response.Pagination.Urls.Next[23:])
		newResponse := &CollectionResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		releases = append(releases, newResponse.Releases...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	procReleases := []*pb.Release{}
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
	ID      int64 `json:"id"`
	Release BasicRelease
	Posted  string `json:"posted"`
}

// GetInventory gets all the releases that are for sale
func (r *DiscogsRetriever) GetInventory(ctx context.Context) ([]*pb.ForSale, error) {
	jsonString, _, err := r.retrieve(ctx, "/users/BrotherLogic/inventory?status=For%20Sale&per_page=100&token="+r.userToken)
	if err != nil {
		return []*pb.ForSale{}, err
	}
	//log.Printf("JSON %v", string(jsonString))

	var items []*InventoryEntry
	response := &InventoryResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	items = append(items, response.Listings...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _, err = r.retrieve(ctx, response.Pagination.Urls.Next[23:])
		if err != nil {
			return []*pb.ForSale{}, err
		}
		newResponse := &InventoryResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		items = append(items, newResponse.Listings...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	sales := []*pb.ForSale{}
	for _, re := range items {
		p, err := time.Parse("2006-01-02T15:04:05-07:00", re.Posted)
		if err != nil {
			return nil, err
		}
		sales = append(sales,
			&pb.ForSale{
				Id:         int32(re.Release.ID),
				SaleId:     re.ID,
				DatePosted: int64(p.Unix()),
				SalePrice:  int32(math.Round(float64(re.Price.Value * 100)))})
	}

	return sales, nil
}

// GetSalePrice gets the sale price for a release
func (r *DiscogsRetriever) GetSalePrice(ctx context.Context, releaseID int) (float32, error) {
	jsonString, _, err := r.retrieve(ctx, "/marketplace/price_suggestions/"+strconv.Itoa(releaseID)+"?token="+r.userToken)
	var resp map[string]Pricing
	r.unmarshaller.Unmarshal(jsonString, &resp)
	value := resp["Mint (M)"].Value

	if value == 0 && err == nil {
		return float32(100), err
	}
	return value, err
}

// SellRecord sells a given release
func (r *DiscogsRetriever) SellRecord(ctx context.Context, releaseID int, price float32, state string, condition, sleeve string, weight int) (int64, error) {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"" + strings.TrimSpace(condition) + "\", \"sleeve_condition\":\"" + strings.TrimSpace(sleeve) + "\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"" + state + "\",\"weight\":\"" + fmt.Sprintf("%v", weight) + "\"}"
	databack, err := r.post(ctx, "/marketplace/listings?token="+r.userToken, data)
	if err != nil {
		return -1, fmt.Errorf("Bad return %w (%v)", err, string(databack))
	}
	var resp SellResponse
	err = r.unmarshaller.Unmarshal([]byte(databack), &resp)
	r.Log(ctx, fmt.Sprintf("Receive Sale Response: %v -> %v (%v)", resp, string(databack), err))
	return resp.ListingID, nil
}

// GetCurrentSalePrice gets the current sale price
func (r *DiscogsRetriever) GetCurrentSalePrice(ctx context.Context, saleID int64) float32 {
	jsonString, _, _ := r.retrieve(ctx, fmt.Sprintf("/marketplace/listings/%v?curr_abbr=USD&token="+r.userToken, saleID))
	var resp PriceResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)
	return resp.Price.Value
}

// GetCurrentSaleState gets the current sale state
func (r *DiscogsRetriever) GetCurrentSaleState(ctx context.Context, saleID int64) (pb.SaleState, error) {
	jsonString, _, err := r.retrieve(ctx, fmt.Sprintf("/marketplace/listings/%v?curr_abbr=USD&token="+r.userToken, saleID))
	if err != nil {
		return pb.SaleState_NOT_FOR_SALE, err
	}

	var resp PriceResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)

	if resp.Status == "For Sale" {
		return pb.SaleState_FOR_SALE, nil
	} else if resp.Status == "Expired" {
		return pb.SaleState_EXPIRED, nil
	} else if resp.Status == "Sold" || resp.Status == "Draft" || resp.Status == "Deleted" {
		return pb.SaleState_SOLD, nil
	}

	r.Log(ctx, fmt.Sprintf("Unknown sale status: %v", resp.Status))
	return pb.SaleState_NOT_FOR_SALE, nil
}

// UpdateSalePrice updates the sale price
func (r *DiscogsRetriever) UpdateSalePrice(ctx context.Context, saleID int, releaseID int, condition, sleeve string, price float32) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"" + condition + "\", \"sleeve_condition\":\"" + sleeve + "\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"For Sale\"}"
	_, err := r.post(ctx, "/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// RemoveFromSale removes the listing from sale
func (r *DiscogsRetriever) RemoveFromSale(ctx context.Context, saleID int, releaseID int) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"Near Mint (NM or M-)\", \"price\":5.00, \"status\":\"Draft\"}"
	_, err := r.post(ctx, "/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// ExpireSale removes the listing from sale
func (r *DiscogsRetriever) ExpireSale(ctx context.Context, saleID int64, releaseID int, price float32) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"Near Mint (NM or M-)\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"Expired\"}"
	_, err := r.post(ctx, fmt.Sprintf("/marketplace/listings/%v?curr_abr=USD&token=%v", saleID, r.userToken), data)
	return err
}

// AddToWantlist adds a record to the wantlist
func (r *DiscogsRetriever) AddToWantlist(ctx context.Context, releaseID int) error {
	_, err := r.put(ctx, "/users/BrotherLogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
	return err
}

// RemoveFromWantlist adds a record to the wantlist
func (r *DiscogsRetriever) RemoveFromWantlist(ctx context.Context, releaseID int) error {
	err := r.delete(ctx, "/users/BrotherLogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
	return err
}

// AddToFolderResponse the response back from an add request
type AddToFolderResponse struct {
	InstanceID  int `json:"instance_id"`
	ResourceURL string
	Simple      int
}

// MoveToFolder Moves the given release to the new folder
func (r *DiscogsRetriever) MoveToFolder(ctx context.Context, folderID int, releaseID int, instanceID int, newFolderID int) (string, error) {
	return r.post(ctx, "/users/BrotherLogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "{\"folder_id\": "+strconv.Itoa(newFolderID)+"}")
}

// DeleteInstance removes a record from the collection
func (r *DiscogsRetriever) DeleteInstance(ctx context.Context, folderID int, releaseID int, instanceID int) error {
	return r.delete(ctx, "/users/BrotherLogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "")
}

// ReleaseBack what we get for a single release
type ReleaseBack struct {
	DateAdded  string     `json:"date_added"`
	InstanceID int32      `json:"instance_id"`
	FolderId   int32      `json:"folder_id"`
	Notes      []*pb.Note `json:"notes"`
	Rating     int32      `json:"rating"`
}

// ReleaseResponse what we get back from release
type ReleaseResponse struct {
	Pagination Pagination
	Releases   []ReleaseBack
}

type Stats struct {
	NumHave int32 `json:"num_have"`
	NumWant int32 `json:"num_want"`
}

func (r *DiscogsRetriever) GetStats(ctx context.Context, rid int32) (*Stats, error) {
	jsonString, _, err := r.retrieve(ctx, fmt.Sprintf("/releases/%v/stats?token=%v", rid, r.userToken))
	if err != nil {
		return nil, err
	}

	var stats *Stats
	r.unmarshaller.Unmarshal(jsonString, stats)

	return stats, nil
}

// InstanceInfo some basic details about the instance
type InstanceInfo struct {
	DateAdded        int64
	RecordCondition  string
	SleeveCondition  string
	LastCleanDate    string
	Width            string
	Weight           string
	Sleeve           string
	Keep             string
	Arrived          int64
	FolderId         int32
	Rating           int32
	LastListenTime   int64
	PurchaseLocation string
	PurchasePrice    int32
}

// GetInstanceInfo gets the info for an instance
func (r *DiscogsRetriever) GetInstanceInfo(ctx context.Context, rid int32) (map[int32]*InstanceInfo, error) {
	jsonString, _, err := r.retrieve(ctx, fmt.Sprintf("/users/BrotherLogic/collection/releases/%v?token=%v", rid, r.userToken))
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

		mapper[entry.InstanceID].Rating = entry.Rating

		mapper[entry.InstanceID].FolderId = entry.FolderId

		for _, note := range entry.Notes {
			// Media condition
			if note.FieldId == 1 {
				mapper[entry.InstanceID].RecordCondition = note.Value
			}

			// Sleeve condition
			if note.FieldId == 2 {
				mapper[entry.InstanceID].SleeveCondition = note.Value
			}

			// Last clean date
			if note.FieldId == 5 {
				mapper[entry.InstanceID].LastCleanDate = note.Value
			}

			if note.FieldId == 4 {
				mapper[entry.InstanceID].Width = note.Value
			}

			if note.FieldId == 7 {
				mapper[entry.InstanceID].Weight = note.Value
			}
			if note.FieldId == 9 {
				mapper[entry.InstanceID].Sleeve = note.Value
			}
			if note.FieldId == 10 {
				mapper[entry.InstanceID].Keep = note.Value
			}
			if note.FieldId == 14 {
				mapper[entry.InstanceID].PurchaseLocation = note.Value
			}
			if note.FieldId == 13 {
				// Remove decimal point - rc stores prices in cents
				val, err := strconv.ParseInt(strings.ReplaceAll(note.Value, ".", ""), 10, 32)
				if err != nil {
					return mapper, err
				}
				mapper[entry.InstanceID].PurchasePrice = int32(val)
			}

			if note.FieldId == 12 && note.Value != "" {
				pv, err := time.Parse("2006-01-02", note.Value)
				if err != nil {
					return mapper, err
				}
				mapper[entry.InstanceID].Arrived = pv.Unix()
			}

			if note.FieldId == 6 && note.Value != "" {
				pv, err := time.Parse("2006-01-02", note.Value)
				if err != nil {
					return mapper, err
				}
				mapper[entry.InstanceID].LastListenTime = pv.Unix()
			}
		}

	}
	return mapper, nil
}

// FoldersResponse returned from discogs
type FoldersResponse struct {
	Pagination Pagination
	Folders    []pb.Folder
}

// GetFolders gets all the folders for a given user
func (r *DiscogsRetriever) GetFolders(ctx context.Context) []pb.Folder {
	jsonString, _, _ := r.retrieve(ctx, "/users/BrotherLogic/collection/folders?token="+r.userToken)

	var folders []pb.Folder
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

func (r *DiscogsRetriever) updateRateLimit(ctx context.Context, resp *http.Response, method string) {
	limit, err := strconv.Atoi(resp.Header.Get("X-Discogs-Ratelimit"))
	if err != nil {
		r.Log(ctx, fmt.Sprintf("Unable to parse rate limit: %v", err))
	}
	rateLimit.With(prometheus.Labels{"method": method}).Set(float64(limit))

	used, err := strconv.Atoi(resp.Header.Get("X-Discogs-Ratelimit-Used"))
	if err != nil {
		r.Log(ctx, fmt.Sprintf("Unable to parse rate limit used: %v", err))
	}
	rateLimitUsed.With(prometheus.Labels{"method": method}).Set(float64(used))
}

func (r *DiscogsRetriever) retrieve(ctx context.Context, path string) ([]byte, http.Header, error) {
	urlv := "https://api.discogs.com/" + path
	if strings.HasPrefix(path, "/") {
		urlv = "https://api.discogs.com/" + path[1:]
	}

	t1 := time.Now()
	r.throttle()
	t2 := time.Now()
	t := time.Now()
	DiscogsRequests.With(prometheus.Labels{"method": "GET", "path1": strings.Split(path, "/")[0]}).Inc()
	response, err := r.getter.Get(urlv)
	RequestLatency.With(prometheus.Labels{"method": "GET", "path1": strings.Split(path, "/")[0]}).Observe(float64(time.Now().Sub(t).Milliseconds()))

	r.Log(ctx, fmt.Sprintf("%v in %v with %v", urlv, time.Since(t2), t2.Sub(t1)))
	if err != nil {
		return make([]byte, 0), make(http.Header), err
	}
	r.updateRateLimit(ctx, response, "GET")

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 && response.StatusCode != 0 {
		r.Log(ctx, fmt.Sprintf("RETR (%v) %v -> %v", path, response.StatusCode, string(body)))
		if response.StatusCode == 404 {
			return body, response.Header, status.Errorf(codes.NotFound, "Bad Read: (%v) %v", response.StatusCode, string(body))
		}
		if response.StatusCode == 429 {
			return body, response.Header, status.Errorf(codes.ResourceExhausted, "Bad Read: (%v) %v", response.StatusCode, string(body))
		}
		return body, response.Header, fmt.Errorf("Bad Read: (%v) %v", response.StatusCode, string(body))
	}

	lastTimeRetrieved = time.Now()

	return body, response.Header, nil
}
