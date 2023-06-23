package models

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
