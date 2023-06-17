package wallapop

// ├── Setters  for 'wallapop'/
// │   ├── SetWallapopItems()
// │   ├── SetWallapopReviews()
// │   ├── SetReviewsAverage()
// │   ├── SetWallapopProfileInfo()
// │   ├── SetFormattedCreationDate()
// │   └── SetWallapopCompleteProfile()
// └── Getters  for 'wallapop'/
//
//	├── GetReviewsAverage()
//	├── GetWallapopProfileInfo()
//	├── GetWallapopItems()
//	└── GetWallapopReviews()

type Wallapop struct {
	UserID          string
	WallapopProfile WallapopProfileInfo
	WallapopItems   WallapopItems
	WallapopReviews WallapopReviews
}

// New returns a new Wallapop struct
func New(UserID string) *Wallapop {
	return &Wallapop{
		UserID: UserID,
	}
}
