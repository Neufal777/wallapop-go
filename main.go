package main

import (
	"strconv"

	"github.com/walla-chollo/src"
)

func main() {

	downloader := src.NewDownloader().
		SetLocation("Calle de la Princesa, 3, Madrid, Spain").
		SetSearch("ford focus").
		SetCategory("100")

	// engine := src.NewEngine().
	// 	SetDownloader(*downloader)
	items := downloader.GetWallapopContent()

	for _, item := range items.SearchObjects {
		price := strconv.FormatFloat(item.Content.Price, 'f', 2, 64)
		println("[ "+price+"€ ]", item.Content.Title, "https://es.wallapop.com/item/"+item.Content.WebSlug)
	}

}
