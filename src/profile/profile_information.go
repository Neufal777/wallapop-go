package profile

import (
	"fmt"

	"github.com/walla-chollo/src/httpRequest"
)

func WallapopProfileInformation(userID string) WallapopProfileInfo {
	url := "https://api.wallapop.com/api/v3/users/" + userID
	result := WallapopProfileInfo{}

	err := httpRequest.GetAPIResponse(url, &result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return WallapopProfileInfo{}
	}

	// fmt.Printf("User: %s | %s \n", result.MicroName, result.ID)
	return result
}

func WallapopProfileItems(userID string) WallapopItems {
	url := "https://api.wallapop.com/api/v3/users/" + userID + "/items"
	result := WallapopItems{}

	err := httpRequest.GetAPIResponse(url, &result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return WallapopItems{}
	}

	return result
}

func WallapopProfileReviews(userID string) WallapopReviews {
	url := "https://api.wallapop.com/api/v3/users/" + userID + "/reviews"
	result := WallapopReviews{}

	err := httpRequest.GetAPIResponse(url, &result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return WallapopReviews{}
	}

	return result
}
