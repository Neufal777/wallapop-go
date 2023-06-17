package docs

import (
	"log"

	"github.com/walla-chollo/src/downloader"
	"github.com/walla-chollo/src/verticals"
	"github.com/walla-chollo/tools"
)

func GetLowReviews() {
	wallapop := downloader.NewDownloader().
		SetLocation("providencia 33 manresa 08241 Barcelona").
		SetCategory("100").
		SetSearch("").
		SetLimit(1).
		SetCarParams(verticals.Car{
			Km:    200000,
			Price: 500000,
			Year:  2023,
			// Gearbox: "manual",
			// Fuel:    "gasoil", //gasoline,gasoil
		}).
		SetOffset(0).
		GetWallapopContent()

	log.Println(
		tools.PrettyPrintStruct(wallapop),
	)
	//Get users with low reviews
	// _ = wallapopContent.SetUserToItems()

	// for i := 0; i < len(wallapopContentWallapopRequestResponse); i++ {
	// 	fmt.Println(wallapopContentWallapopRequestResponse[i].SearchObjects[i].Content.Title)
	// }

	// profile := profile.NewProfile().
	// SetWallapopProfile("kp61k1m852z5").
	// SetReviewsAverage().
	// GetReviewsAverage()

}
