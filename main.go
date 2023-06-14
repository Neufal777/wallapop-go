package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/walla-chollo/src/downloader"
	"github.com/walla-chollo/src/profile"
	"github.com/walla-chollo/src/verticals"
	"github.com/walla-chollo/utils"
)

type Result struct {
	UserID    string `json:"userId"`
	UserName  string `json:"userName"`
	ReviewAvg struct {
		Buy   float64 `json:"buy"`
		Sell  float64 `json:"sell"`
		Count int     `json:"count"`
	} `json:"reviewAvg"`
}

func main() {
	wallapopContent := downloader.NewDownloader().
		SetLocation("providencia 33 manresa 08241 Barcelona").
		SetCategory("100").
		SetSearch("mercedes").
		SetLimit(20).
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

	result := []Result{}
	for k := range mapItemsUsers {
		profile := profile.NewProfile().
			SetWallapopProfile(k).
			SetReviewsAverage()

		review_avg := profile.GetReviewsAverage()

		if review_avg.Buy < 80 && review_avg.Buy != 0 && review_avg.Count != 0 && profile.WallapopProfile.Type != "professional" ||
			review_avg.Sell < 80 && review_avg.Sell != 0 && review_avg.Count != 0 && profile.WallapopProfile.Type != "professional" {
			color.Red("User: %s | %s \n", profile.WallapopProfile.MicroName, profile.WallapopProfile.ID)

			//print phishy users
			fmt.Println(utils.PrettyPrintStruct(review_avg))

			result = append(result, Result{
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

			SaveJSON(result)
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
}

func SaveJSON(result []Result) {
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("phisy_accounts.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
