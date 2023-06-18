package main

import (
	"github.com/wallapop-go/src/wallapop"
)

func main() {
	wall := wallapop.New("npj9q77050ze")

	// Sets ALL the information for the user (profile info, items, reviews, etc.)
	// In case you:
	// just need the profile info, you can use wallapop.SetWallapopProfileInfo()
	// just need the items, you can use wallapop.SetWallapopItems()
	// just need the reviews, you can use wallapop.SetWallapopReviews()
	wall.SetWallapopAll()

	// In case you:
	// want all the information, leave the parameter empty
	// just need the items, you can use wallapop.CliWallapopPrint("profile")
	// just need the profile info, you can use wallapop.CliWallapopPrint("items")
	// just need the reviews, you can use wallapop.CliWallapopPrint("reviews")
	// need multiple information, you can set them all, wallapop.CliWallapopPrint("profile", "products", etc.)
	wall.CliWallapopPrint("items", "profile")

}
