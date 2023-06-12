package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/walla-chollo/src/verticals"
)

type Downloader struct {
	Search    string        `json:"search"`     //Search query
	Category  string        `json:"category"`   //Categories: cars (100), etc..
	Order     string        `json:"order"`      //Order by: most_relevance, etc
	Limit     int           `json:"limit"`      //Number of results to return
	Offset    int           `json:"offset"`     //Offset of the results
	Location  Location      `json:"location"`   //Location of the search
	CarFields verticals.Car `json:"car_fields"` //Car to search
}

type Location struct {
	Address   string  `json:"address"`   // Street address of the location
	City      string  `json:"city"`      // City where the location is situated
	ZipCode   string  `json:"zipCode"`   // Postal/ZIP code of the location
	Longitude float64 `json:"longitude"` // Longitude of the location
	Latitude  float64 `json:"latitude"`  // Latitude of the location
}

func (dow *Downloader) SetLocation(address string) *Downloader {
	searchGeo := HereGeoCoordinates(address)
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
	dow.CarFields.Brand = search.Brand
	dow.CarFields.Model = search.Model
	dow.CarFields.Year = search.Year
	dow.CarFields.Version = search.Version
	dow.CarFields.Km = search.Km
	dow.CarFields.Engine = search.Engine
	dow.CarFields.Gearbox = search.Gearbox
	dow.CarFields.Price = search.Price

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

func (dow *Downloader) GetWallapopContent() WallapopRequestResponse {
	url := fmt.Sprintf("https://api.wallapop.com/api/v3/cars/search?keywords=%s&filters_source=search_box&latitude=%f&start=%d&order_by=most_relevance&step=2&category_ids=%s&longitude=%f&max_sale_price=%d&max_year=%d&max_km=%d",
		dow.Search, dow.Location.Latitude, 20, dow.Category, dow.Location.Longitude, dow.CarFields.Price, dow.CarFields.Year, dow.CarFields.Km)

	log.Println(url)
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make the HTTP request: %s\n", err)
		return WallapopRequestResponse{}
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read the response body: %s\n", err)
		return WallapopRequestResponse{}
	}

	// Create a new instance of the Content struct
	content := WallapopRequestResponse{}

	// Unmarshal the JSON response into the Content struct
	err = json.Unmarshal(body, &content)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %s\n", err)
		return WallapopRequestResponse{}
	}

	return content
}
