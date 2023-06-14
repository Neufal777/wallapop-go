package profile

import (
	"github.com/walla-chollo/utils"
)

type Profile struct {
	WallapopProfile WallapopProfileInfo `json:"wallapop_profile"`
	WallapopItems   WallapopItems       `json:"items"`
	WallapopReviews WallapopReviews     `json:"reviews"`
	ReviewsAverage  WallapopReview      `json:"reviews_average"`
}

type WallapopItems struct {
	Data []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		CategoryID  string `json:"category_id"`
		Slug        string `json:"slug"`
		Images      []struct {
			ID           string `json:"id"`
			AverageColor string `json:"average_color"`
			Urls         struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
				Big    string `json:"big"`
			} `json:"urls"`
		} `json:"images"`
		Price struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"price"`
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
		Shipping struct {
			ItemIsShippable    bool `json:"item_is_shippable"`
			UserAllowsShipping bool `json:"user_allows_shipping"`
		} `json:"shipping"`
	} `json:"data"`
	Meta struct {
		Next string `json:"next"`
	} `json:"meta"`
}

type WallapopReviews []struct {
	Item struct {
		CategoryID int `json:"category_id"`
	} `json:"item,omitempty"`
	Review struct {
		ID                    string `json:"id"`
		Date                  int64  `json:"date"`
		Scoring               int    `json:"scoring"`
		IsShippingTransaction bool   `json:"is_shipping_transaction"`
		CanTranslate          bool   `json:"can_translate"`
	} `json:"review,omitempty"`
	User struct {
		ID        string `json:"id"`
		MicroName string `json:"micro_name"`
		WebSlug   string `json:"web_slug"`
	} `json:"user,omitempty"`
	Type string `json:"type"`
}

type WallapopProfileInfo struct {
	FormattedDate string `json:"formatted_date"`
	ID            string `json:"id"`
	MicroName     string `json:"micro_name"`
	Type          string `json:"type"`
	// Image     struct {
	// 	ID              string `json:"id"`
	// 	OriginalWidth   int    `json:"original_width"`
	// 	OriginalHeight  int    `json:"original_height"`
	// 	AverageHexColor string `json:"average_hex_color"`
	// 	UrlsBySize      struct {
	// 		Small    string `json:"small"`
	// 		Xmall    string `json:"xmall"`
	// 		Original string `json:"original"`
	// 		Large    string `json:"large"`
	// 		Xlarge   string `json:"xlarge"`
	// 		Medium   string `json:"medium"`
	// 	} `json:"urls_by_size"`
	// } `json:"image"`
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
		OpeningHours                      string  `json:"opening_hours"`
		NewChatNotification               bool    `json:"new_chat_notification"`
		OnlyChatPhoneNotification         bool    `json:"only_chat_phone_notification"`
		ConsentThirdPartiesUseData        bool    `json:"consent_third_parties_use_data"`
		NewsNotification                  bool    `json:"news_notification"`
		ModifiedDate                      int64   `json:"modified_date"`
	} `json:"extra_info"`
}

func NewProfile() *Profile {
	return &Profile{}
}
func (p *Profile) SetWallapopProfileInfo(userID string) *Profile {
	p.WallapopProfile = WallapopProfileInformation(userID)
	p.SetCreationDate()

	return p
}

func (p *Profile) SetWallapopItems(userID string) *Profile {
	p.WallapopItems = WallapopProfileItems(userID)
	return p
}

func (p *Profile) SetWallapopReviews(userID string) *Profile {
	p.WallapopReviews = WallapopProfileReviews(userID)
	return p
}

func (p *Profile) SetWallapopProfile(userID string) *Profile {
	p.SetWallapopProfileInfo(userID)
	p.SetWallapopItems(userID)
	p.SetWallapopReviews(userID)
	p.SetCreationDate()
	return p
}

type WallapopReview struct {
	Buy   float64
	Sell  float64
	Count int
}

func (p *Profile) SetReviewsAverage() *Profile {
	WallapopReviewsMap := make(map[string][]int)
	WallapopReviews := WallapopReview{}

	for _, review := range p.WallapopReviews {
		WallapopReviewsMap[review.Type] = append(WallapopReviewsMap[review.Type], review.Review.Scoring)
	}

	averages := make(map[string]float64)
	for key, values := range WallapopReviewsMap {
		sum := 0
		for _, value := range values {
			sum += value
		}
		average := float64(sum) / float64(len(values))
		averages[key] = average
	}

	WallapopReviews.Buy = averages["buy"]
	WallapopReviews.Sell = averages["sell"]
	WallapopReviews.Count = len(p.WallapopReviews)

	p.ReviewsAverage = WallapopReviews
	return p
}

func (p *Profile) GetReviewsAverage() WallapopReview {
	return p.ReviewsAverage
}

func (p *Profile) SetCreationDate() *Profile {
	p.WallapopProfile.FormattedDate = utils.TimeFormatter(p.WallapopProfile.RegisterDate)
	return p
}
