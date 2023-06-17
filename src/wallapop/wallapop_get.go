package wallapop

// GetWallapopProfileInfo returns the WallapopProfileInfo struct
func (p *Wallapop) GetWallapopProfileInfo() WallapopProfileInfo {
	return p.WallapopProfile
}

// GetWallapopItems returns the WallapopItems struct
func (p *Wallapop) GetWallapopItems() WallapopItems {
	return p.WallapopItems
}

// GetWallapopReviews returns the WallapopReviews struct
func (p *Wallapop) GetWallapopReviews() WallapopReviews {
	return p.WallapopReviews
}
