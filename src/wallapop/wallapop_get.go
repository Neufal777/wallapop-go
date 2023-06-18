package wallapop

// GetWallapopProfileInfo returns the WallapopProfileInfo struct
func (w *Wallapop) GetWallapopProfileInfo() WallapopProfileInfo {
	return w.WallapopUserProfile
}

// GetWallapopItems returns the WallapopItems struct
func (w *Wallapop) GetWallapopItems() WallapopItems {
	return w.WallapopUserItems
}

// GetWallapopReviews returns the WallapopReviews struct
func (w *Wallapop) GetWallapopReviews() WallapopReviews {
	return w.WallapopUserReviews
}

// Get information for a specific item
// https://api.wallapop.com/api/v3/items/p61ywex3x265
