package api

import (
	"fmt"

	"github.com/wallapop-go/pkg/http"
	"github.com/wallapop-go/pkg/models"
)

type WallapopAPI struct {
	UsersEndpoint string
	ItemsEndpoint string
}

var API = WallapopAPI{
	UsersEndpoint: "https://api.wallapop.com/api/v3/users/",
	ItemsEndpoint: "https://api.wallapop.com/api/v3/items/",
}

func handleError(err error) {
	fmt.Printf("Error: %s\n", err)
}

// HttpWallapop retrieves the data from the Wallapop API
func (w *Wallapop) HttpWallapop(endpoint string, result interface{}) {
	url := fmt.Sprintf("%s%s%s", API.UsersEndpoint, w.User.ID, endpoint)

	fmt.Println(url)
	err := http.GetAPIResponse(url, result, nil)
	if err != nil {
		handleError(err)
	}
}

// HttpWallapopProfileInfo retrieves the profile information of a Wallapop user
func (w *Wallapop) HttpWallapopProfileInfo() models.WallapopProfileInfo {
	var profileInfo models.WallapopProfileInfo
	w.HttpWallapop("/", &profileInfo)
	return profileInfo
}

// HttpWallapopProfileReviews retrieves the reviews of a Wallapop user
func (w *Wallapop) HttpWallapopProfileReviews() models.WallapopReviews {
	var reviews models.WallapopReviews
	route := "/reviews"

	w.HttpWallapop(route, &reviews)
	return reviews
}

// HttpWallapopProfileItems retrieves the items of a Wallapop user for a specific category.
// If no category is provided, it retrieves all items.
func (w *Wallapop) HttpWallapopProfileItems(category ...string) models.WallapopItems {
	var items models.WallapopItems
	route := "/items"

	if len(category) > 0 {
		route = fmt.Sprintf("/items?category_id=%s", category[0])
	}

	w.HttpWallapop(route, &items)
	return items
}
