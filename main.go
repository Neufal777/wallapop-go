package main

import (
	"log"

	"github.com/walla-chollo/src/wallapop"
)

func main() {
	wallapop := wallapop.New("vjrd4dm3786k")

	wallapop.SetWallapopCompleteProfile()

	log.Println(wallapop)
}
