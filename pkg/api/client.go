package api

import "github.com/wallapop-go/pkg/models"

// ├── Setters  for 'wallapop'/
// │   ├── SetWallapopItems()
// │   ├── SetWallapopItemsbyCategory()
// │   ├── SetWallapopProfileInfo()
// │   ├── SetWallapopReviews()
// │   ├── SetWallapopCompleteProfile()
// └── Getters  for 'wallapop'/
//	├── GetWallapopProfileInfo()
//	├── GetWallapopItems()
//	├── GetWallapopReviews()

type Wallapop struct {
	User                struct{ ID string }
	WallapopUserProfile models.WallapopProfileInfo
	WallapopUserItems   models.WallapopItems
	WallapopUserReviews models.WallapopReviews
}

// New returns a new Wallapop struct
func New(UserID string) *Wallapop {
	return &Wallapop{
		User: struct{ ID string }{ID: UserID},
	}
}

// SetWallapopItems sets the WallapopItems struct for a specific category.
// If no category is provided, it sets the items for all categories.
// Currently, only one category can be set. future: multiple categories
func (w *Wallapop) SetWallapopItems(category ...string) *Wallapop {
	if len(category) > 0 {
		w.WallapopUserItems = w.HttpWallapopProfileItems(category[0])
	} else {
		w.WallapopUserItems = w.HttpWallapopProfileItems()
	}

	return w
}

// SetWallapopProfileInfo sets the WallapopProfileInfo struct
func (w *Wallapop) SetWallapopProfileInfo() *Wallapop {
	w.WallapopUserProfile = w.HttpWallapopProfileInfo()
	return w
}

// SetWallapopReviews sets the WallapopReviews struct
func (w *Wallapop) SetWallapopReviews() *Wallapop {
	w.WallapopUserReviews = w.HttpWallapopProfileReviews()
	return w
}

// SetWallapopAll sets all the information
func (w *Wallapop) SetWallapopAll() *Wallapop {
	w.SetWallapopProfileInfo()
	w.SetWallapopItems()
	w.SetWallapopReviews()
	return w
}

// GetWallapopProfileInfo returns the WallapopProfileInfo struct
func (w *Wallapop) GetWallapopProfileInfo() models.WallapopProfileInfo {
	return w.WallapopUserProfile
}

// GetWallapopItems returns the WallapopItems struct
func (w *Wallapop) GetWallapopItems() models.WallapopItems {
	return w.WallapopUserItems
}

// GetWallapopReviews returns the WallapopReviews struct
func (w *Wallapop) GetWallapopReviews() models.WallapopReviews {
	return w.WallapopUserReviews
}
