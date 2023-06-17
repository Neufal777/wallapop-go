package wallapop

func (p *Wallapop) SetWallapopItemsbyCategory(category string) string {
	return ""
}

// SetWallapopProfileInfo sets the WallapopProfileInfo struct
func (p *Wallapop) SetWallapopProfileInfo() *Wallapop {
	p.WallapopProfile = p.HttpWallapopProfileInfo()
	return p
}

// SetWallapopProfile sets the WallapopProfileInfo struct
func (p *Wallapop) SetWallapopItems() *Wallapop {
	p.WallapopItems = p.HttpWallapopProfileItems()
	return p
}

// SetWallapopReviews sets the WallapopReviews struct
func (p *Wallapop) SetWallapopReviews() *Wallapop {
	p.WallapopReviews = p.HttpWallapopProfileReviews()
	return p
}

// SetWallapopCompleteProfile sets the Wallapop struct
func (p *Wallapop) SetWallapopCompleteProfile() *Wallapop {
	p.SetWallapopProfileInfo()
	p.SetWallapopItems()
	p.SetWallapopReviews()
	return p
}
