package wallapop

import "github.com/wallapop-go/src/location"

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
	SearchObjects   []SearchObjects   `json:"search_objects"`
	From            int               `json:"from"`
	To              int               `json:"to"`
	DistanceOrdered bool              `json:"distance_ordered"`
	Keywords        string            `json:"keywords"`
	SearchPoint     location.Location `json:"search_point"`
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

type WallapopProfileInfo struct {
	ID        string `json:"id"`
	MicroName string `json:"micro_name"`
	Type      string `json:"type"`
	Image     struct {
		ID              string `json:"id"`
		OriginalWidth   int    `json:"original_width"`
		OriginalHeight  int    `json:"original_height"`
		AverageHexColor string `json:"average_hex_color"`
		UrlsBySize      struct {
			Small    string `json:"small"`
			Xsmall   string `json:"xsmall"`
			Original string `json:"original"`
			Large    string `json:"large"`
			Xlarge   string `json:"xlarge"`
			Medium   string `json:"medium"`
		} `json:"urls_by_size"`
	} `json:"image"`
	Location struct {
		ApproximatedLatitude  float64 `json:"approximated_latitude"`
		ApproximatedLongitude float64 `json:"approximated_longitude"`
		City                  string  `json:"city"`
		Zip                   string  `json:"zip"`
		CountryCode           string  `json:"country_code"`
		ApproxRadius          int     `json:"approxRadius"`
		ApproximatedLocation  bool    `json:"approximated_location"`
	} `json:"location"`
	Gender       string `json:"gender"`
	WebSlug      string `json:"web_slug"`
	URLShare     string `json:"url_share"`
	RegisterDate int64  `json:"register_date"`
	Featured     bool   `json:"featured"`
	ExtraInfo    struct {
		UpdatedLatitude                   bool    `json:"updatedLatitude"`
		UpdatedLongitude                  bool    `json:"updatedLongitude"`
		UpdatedAddress                    bool    `json:"updatedAddress"`
		UpdatedLink                       bool    `json:"updatedLink"`
		UpdatedDescription                bool    `json:"updatedDescription"`
		UpdatedPhoneNumber                bool    `json:"updatedPhoneNumber"`
		UpdatedModifiedDate               bool    `json:"updatedModifiedDate"`
		UpdatedOpeningHours               bool    `json:"updatedOpeningHours"`
		UpdatedNewChatNotification        bool    `json:"updatedNewChatNotification"`
		UpdatedOnlyChatPhoneNotification  bool    `json:"updatedOnlyChatPhoneNotification"`
		UpdatedConsentThirdPartiesUseData bool    `json:"updatedConsentThirdPartiesUseData"`
		UpdatedNewsNotification           bool    `json:"updatedNewsNotification"`
		Description                       string  `json:"description"`
		Address                           string  `json:"address"`
		PhoneNumber                       string  `json:"phone_number"`
		Link                              string  `json:"link"`
		Latitude                          float64 `json:"latitude"`
		Longitude                         float64 `json:"longitude"`
		CountryCode                       string  `json:"country_code"`
		City                              string  `json:"city"`
		OpeningHours                      string  `json:"opening_hours"`
		SmsNotification                   bool    `json:"sms_notification"`
		ChatNotification                  bool    `json:"chat_notification"`
		NewsNotification                  bool    `json:"news_notification"`
		ShowOnline                        bool    `json:"show_online"`
		OnlyChatPhoneNotification         bool    `json:"only_chat_phone_notification"`
		ConsentThirdPartiesUseData        bool    `json:"consent_third_parties_use_data"`
	} `json:"extra_info"`
	LanguageCode    string `json:"language_code"`
	RatingAverage   int    `json:"rating_average"`
	RatingCounts    int    `json:"rating_counts"`
	IsVisible       bool   `json:"is_visible"`
	ReportedAsScam  bool   `json:"reported_as_scam"`
	LastActivity    int64  `json:"last_activity"`
	ScamReports     int    `json:"scam_reports"`
	HasScamReports  bool   `json:"has_scam_reports"`
	IsFeatured      bool   `json:"is_featured"`
	LastUpdate      int64  `json:"last_update"`
	TicketsCreated  int    `json:"tickets_created"`
	TicketsResolved int    `json:"tickets_resolved"`
}

type WallapopUserResponse struct {
	Data struct {
		ID          string              `json:"id"`
		Type        string              `json:"type"`
		ProfileInfo WallapopProfileInfo `json:"profile_info"`
	} `json:"data"`
}
