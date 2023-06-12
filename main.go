package main

import (
	"strconv"

	"github.com/walla-chollo/src"
	"github.com/walla-chollo/src/verticals"
)

func main() {

	downloader := src.NewDownloader().
		SetLocation("Calle de la Princesa, 3, Madrid, Spain").
		SetCategory("100").
		SetSearch("citroen c4").
		SetCarParams(verticals.Car{
			Km:    200000,
			Price: 5000,
			Year:  2012,
		})

	// engine := src.NewEngine().
	// 	SetDownloader(*downloader)
	items := downloader.GetWallapopContent()

	for _, item := range items.SearchObjects {
		//convert float to string
		price := strconv.FormatFloat(item.Content.Price, 'f', 2, 64)
		km := strconv.Itoa(item.Content.Km)

		println("[ "+price+" € ] [ "+km+" km]", item.Content.Title, "https://es.wallapop.com/item/"+item.Content.WebSlug)
	}

}
