package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wallapop-go/pkg/models"
)

func TestGetWallapopCategoryData(t *testing.T) {
	// Create a mock HTTP server to simulate the API response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the response body to a sample JSON
		response := `{
			"categories": [
				{
					"id": 100,
					"name": "Cars",
					"vertical_id": "cars",
					"icon": "car",
					"attributes": {},
					"presentation": {
						"image_url": "https://category-bg-images.wallapop.com/100_Cars.png",
						"title_color": "04413C",
						"background_color": "13C1AC"
					},
					"subcategories": []
				},
				{
					"id": 200,
					"name": "Real Estate",
					"vertical_id": "real_estate",
					"icon": "house",
					"attributes": {},
					"presentation": {
						"image_url": "https://category-bg-images.wallapop.com/200_Real-estate.png",
						"title_color": "670A6F",
						"background_color": "FF89C2"
					},
					"subcategories": []
				}
			]
		}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Perform the test
	categories := GetWallapopCategoryData(100, 200)

	// Define the expected categories
	expectedCategories := []models.WallapopCategory{
		{
			ID:         100,
			Name:       "Cars",
			VerticalID: "cars",
			Icon:       "car",
			Attributes: struct{}{},
			Presentation: struct {
				ImageURL        string `json:"image_url"`
				TitleColor      string `json:"title_color"`
				BackgroundColor string `json:"background_color"`
			}{
				ImageURL:        "https://category-bg-images.wallapop.com/100_Cars.png",
				TitleColor:      "04413C",
				BackgroundColor: "13C1AC",
			},
			Subcategories: []interface{}{},
		},
		{
			ID:         200,
			Name:       "Real Estate",
			VerticalID: "real_estate",
			Icon:       "house",
			Attributes: struct{}{},
			Presentation: struct {
				ImageURL        string `json:"image_url"`
				TitleColor      string `json:"title_color"`
				BackgroundColor string `json:"background_color"`
			}{
				ImageURL:        "https://category-bg-images.wallapop.com/200_Real-estate.png",
				TitleColor:      "670A6F",
				BackgroundColor: "FF89C2",
			},
			Subcategories: []interface{}{},
		},
	}

	// Assert that the returned categories match the expected categories
	assert.Equal(t, expectedCategories, categories)
}
