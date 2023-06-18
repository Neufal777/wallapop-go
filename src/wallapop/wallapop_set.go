package wallapop

// SetWallapopItems sets the WallapopItems struct for a specific category.
// If no category is provided, it sets the items for all categories.
// Categories options: Cars: "100"
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

// SetWallapopCompleteProfile sets the Wallapop struct
func (w *Wallapop) SetWallapopCompleteProfile() *Wallapop {
	w.SetWallapopProfileInfo()
	w.SetWallapopItems()
	w.SetWallapopReviews()
	return w
}
