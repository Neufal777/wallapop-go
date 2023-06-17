package wallapop

import (
	"fmt"

	"github.com/walla-chollo/src/httpRequest"
)

const (
	API_V3_USERS = "https://api.wallapop.com/api/v3/users/"
)

// WallapopProfileInfo is the struct that contains the information of a wallapop user
func WallapopProfileInformation(userID string) WallapopProfileInfo {
	var (
		url                = API_V3_USERS + userID
		profileInformation = WallapopProfileInfo{}
	)

	err := httpRequest.GetAPIResponse(url, &profileInformation)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return WallapopProfileInfo{}
	}

	return profileInformation
}

// WallapopProfileItems is the struct that contains the items of a wallapop user
func WallapopProfileItems(userID string) WallapopItems {
	var (
		url       = API_V3_USERS + userID + "/items"
		itemsList = WallapopItems{}
	)

	err := httpRequest.GetAPIResponse(url, &itemsList)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return WallapopItems{}
	}

	return itemsList
}

// WallapopProfileReviews is the struct that contains the reviews of a wallapop user
func WallapopProfileReviews(userID string) WallapopReviews {
	var (
		url    = API_V3_USERS + userID + "/reviews"
		result = WallapopReviews{}
	)

	err := httpRequest.GetAPIResponse(url, &result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return WallapopReviews{}
	}

	return result
}
