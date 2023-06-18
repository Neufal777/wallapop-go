package wallapop

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/wallapop-go/tools"
)

const (
	BaseClientURL = "https://es.wallapop.com/item/"
)

func (w *Wallapop) CliWallapopPrint(displayMode ...string) {
	// Set default display modes if displayMode is empty
	if len(displayMode) == 0 {
		displayMode = []string{"profile", "items", "reviews"}
	}

	// Print the user profile information (name, location, etc.)
	if tools.Contains(displayMode, "profile") {
		profileinfo_table := tablewriter.NewWriter(os.Stdout)
		profileinfo_table.SetHeader([]string{"ID", "Name", "Location", "url_share", "since"})
		profileinfo_table.Append([]string{
			w.WallapopUserProfile.ID,
			w.WallapopUserProfile.MicroName,
			w.WallapopUserProfile.Location.City,
			w.WallapopUserProfile.URLShare,
			tools.TimeFormatter(w.WallapopUserProfile.RegisterDate),
		})
		profileinfo_table.Render()
	}

	// Print the listed items (title, price, etc.)
	if tools.Contains(displayMode, "items") {
		Products_table := tablewriter.NewWriter(os.Stdout)
		Products_table.SetHeader([]string{"Title", "Price", "Category", "URL"})

		for _, item := range w.WallapopUserItems.Data {
			Products_table.Append([]string{item.Title, fmt.Sprintf("%.2f €", item.Price.Amount), item.CategoryID, BaseClientURL + item.Slug})
		}
		Products_table.Render()
	}

	// Print the reviews (category, user, score)
	if tools.Contains(displayMode, "reviews") {
		reviews_table := tablewriter.NewWriter(os.Stdout)
		reviews_table.SetHeader([]string{"Category ID", "User ID", "Comment", "Scoring"})

		for _, review := range w.WallapopUserReviews {
			reviews_table.Append([]string{fmt.Sprintf("%d", review.Item.CategoryID), review.User.ID, review.Review.Comments, fmt.Sprintf("%d", review.Review.Scoring)})
		}
		reviews_table.Render()
	}
}
