package main

import (
	"github.com/wallapop-go/src/wallapop"
)

func main() {

	// Create a new Wallapop Object
	// Set the user ID, can be found in the URL of the user profile page
	// Example: https://es.wallapop.com/app/user/ocasionplusg-437879034-8j3y83q89169/published
	// The user ID in this case is: 8j3y83q89169

	// In case you want to retrieve specific information, you can use the following methods:
	wall := wallapop.New("8j3y83q89169")

	wall.SetWallapopItems()       // Optional
	wall.SetWallapopProfileInfo() // Optional
	wall.SetWallapopReviews()     // Optional

	// The code above is the same as:
	_ = wallapop.New("8j3y83q89169")
	wall.SetWallapopAll() // This method sets all the information (profile, items, reviews)

	// To retrieve the data set to the Wallapop Object, you can use the following methods:
	wall.GetWallapopItems()       // return a WallapopItems{...}
	wall.GetWallapopProfileInfo() // return WallapopProfileInfo{...}
	wall.GetWallapopReviews()     // return WallapopReviews{...}

	// Regarding the tables, you can display them in the CLI with the following methods:

	// Display the tables with the information, you can set the tables you want to display
	wall.CliWallapopPrint("profile", "items") // Options: profile, items, reviews

	// Display all, left empty, it will display all the tables
	wall.CliWallapopPrint()

	// Categories are identified by their ID, you can find the ID in Item object (CategoryID)
	// If you need to retrieve the category information, you can use the following function,
	// This will return a WallapopCategory{...} object with the information of the category
	_ = wallapop.GetWallapopCategoryData(14000) // 14000 is the ID of the category

}
