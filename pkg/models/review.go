package models

type WallapopReviews []struct {
	Item   ItemReview `json:"item,omitempty"`
	Review ReviewInfo `json:"review,omitempty"`
	User   User       `json:"user,omitempty"`
	Type   string     `json:"type"`
}

type ItemReview struct {
	CategoryID int `json:"category_id"`
}

type ReviewInfo struct {
	ID                    string `json:"id"`
	Date                  int64  `json:"date"`
	Scoring               int    `json:"scoring"`
	IsShippingTransaction bool   `json:"is_shipping_transaction"`
	CanTranslate          bool   `json:"can_translate"`
	Comments              string `json:"comments"`
}
