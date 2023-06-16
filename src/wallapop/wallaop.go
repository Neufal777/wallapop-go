package wallapop

import (
	"github.com/walla-chollo/utils"
)

// ├── Setters  for 'wallapop'/
// │   ├── SetWallapopItems()
// │   ├── SetWallapopReviews()
// │   ├── SetReviewsAverage()
// │   ├── SetWallapopProfileInfo()
// │   ├── SetFormattedCreationDate()
// │   └── SetWallapopCompleteProfile()
// └── Getters  for 'wallapop'/
//     ├── GetReviewsAverage()
//     ├── GetWallapopProfileInfo()
//     ├── GetWallapopItems()
//     └── GetWallapopReviews()

type Wallapop struct {
	WallapopProfile WallapopProfileInfo
	WallapopItems   WallapopItems
	WallapopReviews WallapopReviews
	ReviewsAverage  WallapopReview
}

// NewProfile returns a new Wallapop struct
func NewProfile() *Wallapop {
	return &Wallapop{}
}

// SetWallapopProfile sets the WallapopProfileInfo struct
func (p *Wallapop) SetWallapopItems(userID string) *Wallapop {
	p.WallapopItems = WallapopProfileItems(userID)
	return p
}

// SetWallapopReviews sets the WallapopReviews struct
func (p *Wallapop) SetWallapopReviews(userID string) *Wallapop {
	p.WallapopReviews = WallapopProfileReviews(userID)
	return p
}

// SetReviewsAverage sets the WallapopReview struct
func (p *Wallapop) SetReviewsAverage() *Wallapop {
	WallapopReviewsMap := make(map[string][]int)

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

	p.ReviewsAverage = WallapopReview{
		Buy:   averages["buy"],
		Sell:  averages["sell"],
		Count: len(p.WallapopReviews),
	}
	return p
}

// SetWallapopProfileInfo sets the WallapopProfileInfo struct
func (p *Wallapop) SetWallapopProfileInfo(userID string) *Wallapop {
	p.WallapopProfile = WallapopProfileInformation(userID)
	p.SetFormattedCreationDate()

	return p
}

// SetCreationDate sets the FormattedDate field in WallapopProfileInfo struct
func (p *Wallapop) SetFormattedCreationDate() *Wallapop {
	// Turns the date timestamp into a human readable format
	p.WallapopProfile.FormattedDate = utils.TimeFormatter(p.WallapopProfile.RegisterDate)
	return p
}

// SetWallapopCompleteProfile sets the Wallapop struct
func (p *Wallapop) SetWallapopCompleteProfile(userID string) *Wallapop {
	p.SetWallapopProfileInfo(userID)
	p.SetWallapopItems(userID)
	p.SetWallapopReviews(userID)
	p.SetReviewsAverage()
	return p
}

// GetReviewsAverage returns the WallapopReview struct
func (p *Wallapop) GetReviewsAverage() WallapopReview {
	return p.ReviewsAverage
}

// GetWallapopProfileInfo returns the WallapopProfileInfo struct
func (p *Wallapop) GetWallapopProfileInfo() WallapopProfileInfo {
	return p.WallapopProfile
}

// GetWallapopItems returns the WallapopItems struct
func (p *Wallapop) GetWallapopItems() WallapopItems {
	return p.WallapopItems
}

// GetWallapopReviews returns the WallapopReviews struct
func (p *Wallapop) GetWallapopReviews() WallapopReviews {
	return p.WallapopReviews
}
