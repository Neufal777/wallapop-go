package models

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

type Flags struct {
	Pending  bool `json:"pending"`
	Sold     bool `json:"sold"`
	Reserved bool `json:"reserved"`
	Banned   bool `json:"banned"`
	Expired  bool `json:"expired"`
	Onhold   bool `json:"onhold"`
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

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type WallapopItems struct {
	Data []struct {
		ID             string  `json:"id"`
		Title          string  `json:"title"`
		Description    string  `json:"description"`
		CategoryID     string  `json:"category_id"`
		Slug           string  `json:"slug"`
		Images         []Image `json:"images"`
		Price          Price   `json:"price"`
		TypeAttributes struct {
			BodyType struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"body_type"`
			HorsePower struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"horse_power"`
			Km struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"km"`
			Year struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"year"`
			Version struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"version"`
			Seats struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"seats"`
			Doors struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"doors"`
			Condition struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"condition"`
			Engine struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"engine"`
			Model struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"model"`
			GearBox struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"gear_box"`
			Brand struct {
				Value    string `json:"value"`
				Title    string `json:"title"`
				Text     string `json:"text"`
				IconText string `json:"icon_text"`
			} `json:"brand"`
		} `json:"type_attributes,omitempty"`
		Shipping Shipping `json:"shipping"`
	} `json:"data"`
	Meta struct {
		Next string `json:"next"`
	} `json:"meta"`
}
