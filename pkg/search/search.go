// Work in progress
package search

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/wallapop-go/pkg/http"
	"github.com/wallapop-go/pkg/models"
	"github.com/wallapop-go/pkg/services"
)

const (
	WALLAPOP_BASE_SEARCH = "https://api.wallapop.com/api/v3/cars/search"
)

type Car struct {
	Brand   string `json:"brand"`   //Brand of the car
	Model   string `json:"model"`   //Model of the car
	Year    int    `json:"year"`    //Year of the car
	Km      int    `json:"km"`      //Kilometers of the car
	Gearbox string `json:"gearbox"` //Gearbox of the car
	Price   int    `json:"price"`   //Price of the car
	Fuel    string `json:"fuel"`    //Fuel of the car
}

type Downloader struct {
	Search                  string                           `json:"search"`     //Search query
	Category                string                           `json:"category"`   //Categories: cars (100), etc..
	Order                   string                           `json:"order"`      //Order by: most_relevance, etc
	Limit                   int                              `json:"limit"`      //Number of results to return
	Offset                  int                              `json:"offset"`     //Offset of the results
	Location                models.Location                  `json:"location"`   //Location of the search
	CarFields               Car                              `json:"car_fields"` //Car to search
	WallapopRequestResponse []models.WallapopRequestResponse `json:"wallapop_request_response"`
}

func (dow *Downloader) SetLocation(address string) *Downloader {
	searchGeo := services.HereGeoCoordinates(address)
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

func (dow *Downloader) SetCarParams(search Car) *Downloader {
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
			fmt.Printf("[ %s €] [%s] Km, %s https://es.wallapop.com/item/%s \n", price, km, item.Content.Title, item.Content.WebSlug)
		}
	}
}

// GetWallapopContent returns the content of the Wallapop search
func (dow *Downloader) GetWallapopContentPage() *Downloader {
	url := fmt.Sprintf("%s?keywords=%s&filters_source=search_box&latitude=%f&start=%d&order_by=most_relevance&step=2&category_ids=%s&longitude=%f&max_sale_price=%d&max_year=%d&max_km=%d&engine=%s&gearbox=%s",
		WALLAPOP_BASE_SEARCH, dow.Search, dow.Location.Latitude, dow.Offset, dow.Category, dow.Location.Longitude, dow.CarFields.Price, dow.CarFields.Year, dow.CarFields.Km, dow.CarFields.Fuel, dow.CarFields.Gearbox)

	content := models.WallapopRequestResponse{}
	err := http.GetAPIResponse(url, &content, nil)

	log.Println(url)
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
