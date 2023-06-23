package models

type WallapopUserResponse struct {
	Data struct {
		ID          string              `json:"id"`
		Type        string              `json:"type"`
		ProfileInfo WallapopProfileInfo `json:"profile_info"`
	} `json:"data"`
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
