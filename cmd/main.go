package main

import (
	"github.com/Neufal777/wallapop-go/pkg/api"
	"github.com/Neufal777/wallapop-go/pkg/utils"
)

func main() {

	//New Wallapop client
	wall := api.New("9wzvlepe376l")

	//Set the Wallapop user profile, items, and reviews
	wall.SetWallapopAll()

	//Print the Wallapop user profile, items, and reviews
	utils.CliWallapopPrint(wall)
}
