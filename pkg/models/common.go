package models

type Image struct {
	Original       string `json:"original"`
	Xsmall         string `json:"xsmall"`
	Small          string `json:"small"`
	Large          string `json:"large"`
	Medium         string `json:"medium"`
	Xlarge         string `json:"xlarge"`
	OriginalWidth  int    `json:"original_width"`
	OriginalHeight int    `json:"original_height"`
}

type User struct {
	ID        string `json:"id"`
	MicroName string `json:"micro_name"`
	Image     Image  `json:"image"`
	Online    bool   `json:"online"`
	Kind      string `json:"kind"`
	WebSlug   string `json:"web_slug"`
}

type WallapopRequestResponse struct {
	SearchObjects   []SearchObjects `json:"search_objects"`
	From            int             `json:"from"`
	To              int             `json:"to"`
	DistanceOrdered bool            `json:"distance_ordered"`
	Keywords        string          `json:"keywords"`
	SearchPoint     Location        `json:"search_point"`
}

type Location struct {
	Address   string  `json:"address"`   // Street address of the location
	City      string  `json:"city"`      // City where the location is situated
	ZipCode   string  `json:"zipCode"`   // Postal/ZIP code of the location
	Longitude float64 `json:"longitude"` // Longitude of the location
	Latitude  float64 `json:"latitude"`  // Latitude of the location
}
