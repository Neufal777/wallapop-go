package downloader

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/walla-chollo/src/httpRequest"
	"github.com/walla-chollo/src/location"
	"github.com/walla-chollo/src/verticals"
	"github.com/walla-chollo/src/wallapop"
)

const (
	WALLAPOP_BASE_SEARCH = "https://api.wallapop.com/api/v3/cars/search"
)

type Downloader struct {
	Search                  string                             `json:"search"`     //Search query
	Category                string                             `json:"category"`   //Categories: cars (100), etc..
	Order                   string                             `json:"order"`      //Order by: most_relevance, etc
	Limit                   int                                `json:"limit"`      //Number of results to return
	Offset                  int                                `json:"offset"`     //Offset of the results
	Location                location.Location                  `json:"location"`   //Location of the search
	CarFields               verticals.Car                      `json:"car_fields"` //Car to search
	WallapopRequestResponse []wallapop.WallapopRequestResponse `json:"wallapop_request_response"`
}

func (dow *Downloader) SetLocation(address string) *Downloader {
	searchGeo := location.HereGeoCoordinates(address)
	dow.Location.Longitude = searchGeo.Longitude
	dow.Location.Latitude = searchGeo.Latitude
	dow.Location.Address = searchGeo.Address
	dow.Location.City = searchGeo.City
	dow.Location.ZipCode = searchGeo.ZipCode

	return dow
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (dow *Downloader) SetSearch(search string) *Downloader {
	dow.Search = strings.ReplaceAll(search, " ", "%20")
	return dow
}

func (dow *Downloader) SetCarParams(search verticals.Car) *Downloader {
	if search.Brand != "" {
		dow.CarFields.Brand = search.Brand
	}

	if search.Model != "" {
		dow.CarFields.Model = search.Model
	}

	if search.Year != 0 {
		dow.CarFields.Year = search.Year
	}

	if search.Km != 0 {
		dow.CarFields.Km = search.Km
	}

	if search.Gearbox != "" {
		dow.CarFields.Gearbox = search.Gearbox
	}

	if search.Price != 0 {
		dow.CarFields.Price = search.Price
	}

	if search.Fuel != "" {
		dow.CarFields.Fuel = search.Fuel
	}

	return dow
}

func (dow *Downloader) SetCategory(Category string) *Downloader {
	dow.Category = Category
	return dow
}

func (dow *Downloader) SetOrder(Order string) *Downloader {
	dow.Order = Order
	return dow
}

func (dow *Downloader) SetLimit(Limit int) *Downloader {
	dow.Limit = Limit
	return dow
}

func (dow *Downloader) SetOffset(Offset int) *Downloader {
	dow.Offset = Offset
	return dow
}

// Print prints the results of the search
func (dow *Downloader) Print() {
	// Print the search query
	for _, response := range dow.WallapopRequestResponse {
		for _, item := range response.SearchObjects {

			// Convert float to string
			price := strconv.FormatFloat(item.Content.Price, 'f', 2, 64)
			km := strconv.Itoa(item.Content.Km)

			// Print the results
			fmt.Println("[ "+price+" € ] [ "+km+" km]", item.Content.Title, "https://es.wallapop.com/item/"+item.Content.WebSlug)
		}
	}
}

// GetWallapopContent returns the content of the Wallapop search
func (dow *Downloader) GetWallapopContentPage() *Downloader {
	url := fmt.Sprintf("%s?keywords=%s&filters_source=search_box&latitude=%f&start=%d&order_by=most_relevance&step=2&category_ids=%s&longitude=%f&max_sale_price=%d&max_year=%d&max_km=%d&engine=%s&gearbox=%s",
		WALLAPOP_BASE_SEARCH, dow.Search, dow.Location.Latitude, dow.Offset, dow.Category, dow.Location.Longitude, dow.CarFields.Price, dow.CarFields.Year, dow.CarFields.Km, dow.CarFields.Fuel, dow.CarFields.Gearbox)

	content := wallapop.WallapopRequestResponse{}
	err := httpRequest.GetAPIResponse(url, &content)

	if err != nil {
		fmt.Printf("Failed to make the HTTP request: %s\n", err)
	}

	dow.WallapopRequestResponse = append(dow.WallapopRequestResponse, content)

	return dow
}

// GetWallapopContent gets the content of the Wallapop page
func (dow *Downloader) GetWallapopContent() *Downloader {
	for i := 0; i < dow.Limit; i++ {
		dow.Offset += 20
		dow.GetWallapopContentPage()
		color.Green("downloading page... %d of a total %d", i, dow.Limit)

		time.Sleep(1 * time.Second)
	}

	return dow
}

// // SetUserToItems maps the results with the user ID and the Post ID
// func (dow *Downloader) SetUserToItems() map[string][]string {
// 	// Map the results with the user ID and the Post ID
// 	mapResults := make(map[string][]string)

// 	for _, response := range dow.WallapopRequestResponse {
// 		for _, item := range response.SearchObjects {
// 			// Convert float to string
// 			UserID := item.Content.User.ID
// 			PostID := item.Content.ID

// 			mapResults[UserID] = append(mapResults[UserID], PostID)
// 		}
// 	}

// 	return mapResults
// }
