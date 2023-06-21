package wallapop

import (
	"fmt"

	"github.com/wallapop-go/src/httpRequest"
)

// Get different data from wallapop
// Category information
const (
	CategoriesBaseURL = "https://api.wallapop.com/api/v3/categories"
)

type WallapopCategory struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	VerticalID string `json:"vertical_id"`
	Icon       string `json:"icon"`
	Attributes struct {
	} `json:"attributes,omitempty"`
	Presentation struct {
		ImageURL        string `json:"image_url"`
		TitleColor      string `json:"title_color"`
		BackgroundColor string `json:"background_color"`
	} `json:"presentation"`
	Subcategories []any `json:"subcategories"`
}

type WallapopCategories struct {
	Categories []WallapopCategory `json:"categories"`
}

func GetWallapopCategoryData(categories ...int) []WallapopCategory {
	var wallacat WallapopCategories
	requestedCategories := []WallapopCategory{}

	err := httpRequest.GetAPIResponse(CategoriesBaseURL, &wallacat, nil)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	for _, category := range wallacat.Categories {
		for _, requestedCategory := range categories {
			if category.ID == requestedCategory {
				requestedCategories = append(requestedCategories, category)
			}
		}
	}

	return requestedCategories
}
