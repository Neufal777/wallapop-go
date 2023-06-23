package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/wallapop-go/pkg/models"
)

const HERE_API_KEY = "AezPb11S9-e0_qFOWr7UaYnzzBCG0M2doKE05XbZXew"

func HereGeoCoordinates(q string) models.Location {
	apiKey := HERE_API_KEY

	query := url.QueryEscape(q)
	url := fmt.Sprintf("https://geocode.search.hereapi.com/v1/geocode?q=%s&apiKey=%s", query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return models.Location{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return models.Location{}
	}

	// Check if the response body is empty
	if len(body) == 0 {
		fmt.Println("Empty response body")
		return models.Location{}
	}

	// Define a variable to store the response data
	var data map[string]interface{}

	// Unmarshal the JSON response into the data variable
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return models.Location{}
	}

	// Access the desired information from the response data
	if items, ok := data["items"].([]interface{}); ok && len(items) > 0 {
		item := items[0].(map[string]interface{})
		address := item["address"].(map[string]interface{})
		access := item["access"].([]interface{})[0].(map[string]interface{})

		return models.Location{
			Address:   address["label"].(string),
			City:      address["city"].(string),
			ZipCode:   address["postalCode"].(string),
			Longitude: access["lng"].(float64),
			Latitude:  access["lat"].(float64),
		}
	} else {
		fmt.Println("No results found.")
	}

	return models.Location{}
}
