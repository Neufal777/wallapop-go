// Path: src/downloader.go
package src

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

type UserImage struct {
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
	ID        string    `json:"id"`
	MicroName string    `json:"micro_name"`
	Image     UserImage `json:"image"`
	Online    bool      `json:"online"`
	Kind      string    `json:"kind"`
}

type Flags struct {
	Pending  bool `json:"pending"`
	Sold     bool `json:"sold"`
	Reserved bool `json:"reserved"`
	Banned   bool `json:"banned"`
	Expired  bool `json:"expired"`
	Onhold   bool `json:"onhold"`
}

type VisibilityFlags struct {
	Bumped        bool `json:"bumped"`
	Highlighted   bool `json:"highlighted"`
	Urgent        bool `json:"urgent"`
	CountryBumped bool `json:"country_bumped"`
	Boosted       bool `json:"boosted"`
}

type ItemLocation struct {
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

type Shipping struct {
	ItemIsShippable     bool        `json:"item_is_shippable"`
	UserAllowsShipping  bool        `json:"user_allows_shipping"`
	CostConfigurationID interface{} `json:"cost_configuration_id"`
}

type Content struct {
	ID               string          `json:"id"`
	Title            string          `json:"title"`
	Storytelling     string          `json:"storytelling"`
	Distance         float64         `json:"distance"`
	Images           []Image         `json:"images"`
	User             User            `json:"user"`
	Flags            Flags           `json:"flags"`
	VisibilityFlags  VisibilityFlags `json:"visibility_flags"`
	Price            float64         `json:"price"`
	Currency         string          `json:"currency"`
	WebSlug          string          `json:"web_slug"`
	CategoryID       int             `json:"category_id"`
	Brand            string          `json:"brand"`
	Model            string          `json:"model"`
	Year             int             `json:"year"`
	Version          string          `json:"version"`
	Km               int             `json:"km"`
	Engine           string          `json:"engine"`
	Gearbox          string          `json:"gearbox"`
	Horsepower       float64         `json:"horsepower"`
	Favorited        bool            `json:"favorited"`
	CreationDate     int64           `json:"creation_date"`
	ModificationDate int64           `json:"modification_date"`
	ItemLocation     ItemLocation    `json:"location"`
	Shipping         Shipping        `json:"shipping"`
	SupportsShipping bool            `json:"supports_shipping"`
}

type Data struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Content []Content `json:"content"`
}

type SearchObjects struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Content Content `json:"content"`
}

type WallapopRequestResponse struct {
	SearchObjects   []SearchObjects `json:"search_objects"`
	From            int             `json:"from"`
	To              int             `json:"to"`
	DistanceOrdered bool            `json:"distance_ordered"`
	Keywords        string          `json:"keywords"`
	SearchPoint     Location        `json:"search_point"`
}

func NewContent() *Data {
	return &Data{}
}
