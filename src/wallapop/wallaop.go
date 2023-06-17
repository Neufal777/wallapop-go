package wallapop

// ├── Setters  for 'wallapop'/
// │   ├── SetWallapopItems()
// │   ├── SetWallapopReviews()
// │   ├── SetReviewsAverage()
// │   ├── SetWallapopProfileInfo()
// │   ├── SetFormattedCreationDate()
// │   └── SetWallapopCompleteProfile()
// └── Getters  for 'wallapop'/
//     ├── GetReviewsAverage()
//     ├── GetWallapopProfileInfo()
//     ├── GetWallapopItems()
//     └── GetWallapopReviews()

type Wallapop struct {
	WallapopProfile WallapopProfileInfo
	WallapopItems   WallapopItems
	WallapopReviews WallapopReviews
}

// New returns a new Wallapop struct
func New() *Wallapop {
	return &Wallapop{}
}

// SetWallapopProfileInfo sets the WallapopProfileInfo struct
func (p *Wallapop) SetWallapopProfileInfo(userID string) *Wallapop {
	p.WallapopProfile = WallapopProfileInformation(userID)
	return p
}

// SetWallapopProfile sets the WallapopProfileInfo struct
func (p *Wallapop) SetWallapopItems(userID string) *Wallapop {
	p.WallapopItems = WallapopProfileItems(userID)
	return p
}

// SetWallapopReviews sets the WallapopReviews struct
func (p *Wallapop) SetWallapopReviews(userID string) *Wallapop {
	p.WallapopReviews = WallapopProfileReviews(userID)
	return p
}

// SetWallapopCompleteProfile sets the Wallapop struct
func (p *Wallapop) SetWallapopCompleteProfile(userID string) *Wallapop {
	p.SetWallapopProfileInfo(userID)
	p.SetWallapopItems(userID)
	p.SetWallapopReviews(userID)
	return p
}

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
