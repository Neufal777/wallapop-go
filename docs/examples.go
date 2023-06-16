package docs

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/walla-chollo/src/downloader"
	"github.com/walla-chollo/src/executions"
	"github.com/walla-chollo/src/verticals"
	"github.com/walla-chollo/src/wallapop"
	"github.com/walla-chollo/utils"
)

func GetLowReviews() {
	wallapopContent := downloader.NewDownloader().
		SetLocation("providencia 33 manresa 08241 Barcelona").
		SetCategory("100").
		SetSearch("").
		SetLimit(50).
		SetCarParams(verticals.Car{
			Km:    200000,
			Price: 500000,
			Year:  2023,
			// Gearbox: "manual",
			// Fuel:    "gasoil", //gasoline,gasoil
		}).
		SetOffset(0).
		GetWallapopContent()

		//Get users with low reviews
	mapItemsUsers := wallapopContent.SetUserToItems()

	phisy_accounts := []executions.Result{}
	users_processed := []executions.UsersProcessed{}

	total_items := 0
	for k := range mapItemsUsers {
		total_items += len(mapItemsUsers[k])

		profile := wallapop.NewProfile().
			SetWallapopCompleteProfile(k).
			SetReviewsAverage()

		users_processed = append(users_processed, executions.UsersProcessed{
			Name: profile.WallapopProfile.MicroName,
			ID:   profile.WallapopProfile.ID,
		})

		utils.SaveToJSON(users_processed, "users_processed_two.json")

		review_avg := profile.GetReviewsAverage()

		if review_avg.Buy < 80 && review_avg.Buy != 0 && review_avg.Count != 0 && profile.WallapopProfile.Type != "professional" ||
			review_avg.Sell < 80 && review_avg.Sell != 0 && review_avg.Count != 0 && profile.WallapopProfile.Type != "professional" {
			color.Red("User: %s | %s \n", profile.WallapopProfile.MicroName, profile.WallapopProfile.ID)

			//print phishy users
			fmt.Println(utils.PrettyPrintStruct(review_avg))

			phisy_accounts = append(phisy_accounts, executions.Result{
				UserID:   profile.WallapopProfile.ID,
				UserName: profile.WallapopProfile.MicroName,
				ReviewAvg: struct {
					Buy   float64 `json:"buy"`
					Sell  float64 `json:"sell"`
					Count int     `json:"count"`
				}{
					Buy:   review_avg.Buy,
					Sell:  review_avg.Sell,
					Count: review_avg.Count,
				},
			})

			utils.SaveToJSON(phisy_accounts, "phisy_accounts_two.json")

		} else {
			color.Green("User: %s | %s \n", profile.WallapopProfile.MicroName, profile.WallapopProfile.ID)
			fmt.Println(utils.PrettyPrintStruct(review_avg))
		}
	}

	// for i := 0; i < len(wallapopContentWallapopRequestResponse); i++ {
	// 	fmt.Println(wallapopContentWallapopRequestResponse[i].SearchObjects[i].Content.Title)
	// }

	// profile := profile.NewProfile().
	// SetWallapopProfile("kp61k1m852z5").
	// SetReviewsAverage().
	// GetReviewsAverage()

	// fmt.Println(utils.PrettyPrintStruct(profile))

	color.Magenta("Total items: %d", total_items)
}
