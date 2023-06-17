package wallapop

import (
	"fmt"

	"github.com/walla-chollo/src/httpRequest"
)

const (
	APIV3Users = "https://api.wallapop.com/api/v3/users/"
)

func handleError(err error) {
	fmt.Printf("Error: %s\n", err)
}

// HttpWallapopProfileInfo retrieves the profile information of a Wallapop user
func (p *Wallapop) HttpWallapopProfileInfo() WallapopProfileInfo {
	var profileInfo WallapopProfileInfo
	url := fmt.Sprintf("%s%s", APIV3Users, p.UserID)

	err := httpRequest.GetAPIResponse(url, &profileInfo)
	if err != nil {
		handleError(err)
	}

	return profileInfo
}

// HttpWallapopProfileItems retrieves the items of a Wallapop user
func (p *Wallapop) HttpWallapopProfileItems() WallapopItems {
	var items WallapopItems
	url := fmt.Sprintf("%s%s/items", APIV3Users, p.UserID)

	err := httpRequest.GetAPIResponse(url, &items)
	if err != nil {
		handleError(err)
	}

	return items
}

// HttpWallapopProfileReviews retrieves the reviews of a Wallapop user
func (p *Wallapop) HttpWallapopProfileReviews() WallapopReviews {
	var reviews WallapopReviews
	url := fmt.Sprintf("%s%s/reviews", APIV3Users, p.UserID)

	err := httpRequest.GetAPIResponse(url, &reviews)
	if err != nil {
		handleError(err)
	}

	return reviews
}
