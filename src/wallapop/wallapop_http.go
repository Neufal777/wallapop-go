package wallapop

import (
	"fmt"

	"github.com/wallapop-go/src/httpRequest"
)

const (
	API_V3_USERS = "https://api.wallapop.com/api/v3/users/"
	API_V3_ITEM  = "https://api.wallapop.com/api/v3/items/"
)

func handleError(err error) {
	fmt.Printf("Error: %s\n", err)
}

// HttpWallapop retrieves the data of a Wallapop user
func (w *Wallapop) HttpWallapop(path string, result interface{}) {
	url := fmt.Sprintf("%s%s%s", API_V3_USERS, w.User.ID, path)

	err := httpRequest.GetAPIResponse(url, result, nil)
	if err != nil {
		handleError(err)
	}
}

// HttpWallapopProfileInfo retrieves the profile information of a Wallapop user
func (w *Wallapop) HttpWallapopProfileInfo() WallapopProfileInfo {
	var profileInfo WallapopProfileInfo
	w.HttpWallapop("/", &profileInfo)
	return profileInfo
}

// HttpWallapopProfileReviews retrieves the reviews of a Wallapop user
func (w *Wallapop) HttpWallapopProfileReviews() WallapopReviews {
	var reviews WallapopReviews
	w.HttpWallapop("/reviews", &reviews)
	return reviews
}

// HttpWallapopProfileItems retrieves the items of a Wallapop user for a specific category.
// If no category is provided, it retrieves all items.
func (w *Wallapop) HttpWallapopProfileItems(category ...string) WallapopItems {
	var items WallapopItems
	url := "/items"

	if len(category) > 0 {
		url = fmt.Sprintf("/items?category_id=%s", category[0])
	}

	w.HttpWallapop(url, &items)
	return items
}
