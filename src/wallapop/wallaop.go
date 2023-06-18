package wallapop

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
	WallapopUserProfile WallapopProfileInfo
	WallapopUserItems   WallapopItems
	WallapopUserReviews WallapopReviews
}

// New returns a new Wallapop struct
func New(UserID string) *Wallapop {
	return &Wallapop{
		User: struct{ ID string }{ID: UserID},
	}
}
