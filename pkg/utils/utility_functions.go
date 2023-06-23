package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/wallapop-go/pkg/api"
	"github.com/wallapop-go/pkg/http"
	"github.com/wallapop-go/pkg/models"
)

const (
	BaseClientURL = "https://es.wallapop.com/item/"
)

func CliWallapopPrint(w *api.Wallapop, displayMode ...string) {
	// Set default display modes if displayMode is empty
	if len(displayMode) == 0 {
		displayMode = []string{"profile", "items", "reviews"}
	}

	// Print the user profile information
	if Contains(displayMode, "profile") {
		printUserProfile(w)
	}

	// Print the listed items
	if Contains(displayMode, "items") {
		printUserItems(w)
	}

	// Print the reviews
	if Contains(displayMode, "reviews") {
		printUserReviews(w)
	}
}

func printUserProfile(w *api.Wallapop) {
	profileTable := tablewriter.NewWriter(os.Stdout)
	profileTable.SetHeader([]string{"ID", "Name", "Location", "URL Share", "Since", "Type"})

	userProfile := w.WallapopUserProfile
	profileTable.Append([]string{
		userProfile.ID,
		userProfile.MicroName,
		userProfile.Location.City,
		userProfile.URLShare,
		TimeFormatter(userProfile.RegisterDate),
		userProfile.Type,
	})

	profileTable.Render()
}

func printUserItems(w *api.Wallapop) {
	itemsTable := tablewriter.NewWriter(os.Stdout)
	itemsTable.SetHeader([]string{"Title", "Price", "Category", "URL"})

	for _, item := range w.WallapopUserItems.Data {
		itemsTable.Append([]string{
			item.Title,
			fmt.Sprintf("%.2f €", item.Price.Amount),
			item.CategoryID,
			BaseClientURL + item.Slug,
		})
	}

	itemsTable.Render()
}

func printUserReviews(w *api.Wallapop) {
	reviewsTable := tablewriter.NewWriter(os.Stdout)
	reviewsTable.SetHeader([]string{"Category ID", "User ID", "Comment", "Scoring"})

	for _, review := range w.WallapopUserReviews {
		reviewsTable.Append([]string{
			fmt.Sprintf("%d", review.Item.CategoryID),
			review.User.ID,
			review.Review.Comments,
			fmt.Sprintf("%d", review.Review.Scoring),
		})
	}

	reviewsTable.Render()
}

func GetWallapopCategoryData(categories ...int) []models.WallapopCategory {
	var wallacat models.WallapopCategories
	requestedCategories := []models.WallapopCategory{}

	err := http.GetAPIResponse(models.CategoriesBaseURL, &wallacat, nil)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	for _, category := range wallacat.Categories {
		for _, requestedCategory := range categories {
			if category.ID == requestedCategory {
				requestedCategories = append(requestedCategories, category)
			}
		}
	}

	return requestedCategories
}
