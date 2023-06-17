package main

import (
	"log"

	"github.com/walla-chollo/src/wallapop"
	"github.com/walla-chollo/tools"
)

func main() {
	// docs.GetLowReviews()

	wallapop := wallapop.New("kp61k1m852z5")

	wallapop.SetWallapopCompleteProfile()

	log.Println(
		tools.PrettyPrintStruct(
			wallapop.GetWallapopReviews(),
		),
	)
}
