# Wallapop Package

The `wallapop` package provides functionality for interacting with the Wallapop API and retrieving user information, items, and reviews.

## Types

### Wallapop

```go
type Wallapop struct {
    User                User
    WallapopUserProfile WallapopProfileInfo
    WallapopUserItems   WallapopItems
    WallapopUserReviews WallapopReviews
}
```

The `Wallapop` struct represents a Wallapop user and contains fields for storing user information, profile data, items, and reviews.

## Functions

### New

```go
func New(UserID string) *Wallapop
```

The `New` function creates a new instance of the `Wallapop` struct.

Parameters:
- `UserID` (string): The ID of the Wallapop user.

Returns:
- `*Wallapop`: A pointer to the newly created `Wallapop` struct.

## Setters

### SetWallapopItemsbyCategory

```go
func (p *Wallapop) SetWallapopItemsbyCategory(category string) string
```

The `SetWallapopItemsbyCategory` method sets the Wallapop items based on a specific category.

Parameters:
- `category` (string): The category of the items to retrieve.

Returns:
- `string`: An empty string.

### SetWallapopProfileInfo

```go
func (p *Wallapop) SetWallapopProfileInfo() *Wallapop
```

The `SetWallapopProfileInfo` method sets the Wallapop profile information.

Returns:
- `*Wallapop`: A pointer to the `Wallapop` struct.

### SetWallapopItems

```go
func (p *Wallapop) SetWallapopItems() *Wallapop
```

The `SetWallapopItems` method sets the Wallapop items.

Returns:
- `*Wallapop`: A pointer to the `Wallapop` struct.

### SetWallapopReviews

```go
func (p *Wallapop) SetWallapopReviews() *Wallapop
```

The `SetWallapopReviews` method sets the Wallapop reviews.

Returns:
- `*Wallapop`: A pointer to the `Wallapop` struct.

### SetWallapopCompleteProfile

```go
func (p *Wallapop) SetWallapopCompleteProfile() *Wallapop
```

The `SetWallapopCompleteProfile` method sets the complete Wallapop profile by retrieving profile information, items, and reviews.

Returns:
- `*Wallapop`: A pointer to the `Wallapop` struct.

## Getters

### GetWallapopProfileInfo

```go
func (p *Wallapop) GetWallapopProfileInfo() WallapopProfileInfo
```

The `GetWallapopProfileInfo` method returns the Wallapop profile information.

Returns:
- `WallapopProfileInfo`: The profile information of the Wallapop user.

### GetWallapopItems

```go
func (p *Wallapop) GetWallapopItems() WallapopItems
```

The `GetWallapopItems` method returns the Wallapop items.

Returns:
- `WallapopItems`: The items of the Wallapop user.

### GetWallapopReviews

```go
func (p *Wallapop) GetWallapopReviews() WallapopReviews
```

The `GetWallapopReviews` method returns the Wallapop reviews.

Returns:
- `Wall

apopReviews`: The reviews of the Wallapop user.

## Examples

### Creating a Wallapop instance and setting profile information

```go
package main

import (
	"fmt"

	"github.com/your-package-import-path/wallapop"
)

func main() {
	// Create a new Wallapop instance
	w := wallapop.New("your-user-id")

	// Set the Wallapop profile information
	w.SetWallapopProfileInfo()

	// Retrieve the Wallapop profile information
	profileInfo := w.GetWallapopProfileInfo()
	fmt.Println(profileInfo)
}
```

### Retrieving Wallapop items

```go
package main

import (
	"fmt"

	"github.com/your-package-import-path/wallapop"
)

func main() {
	// Create a new Wallapop instance
	w := wallapop.New("your-user-id")

	// Set the Wallapop items
	w.SetWallapopItems()

	// Retrieve the Wallapop items
	items := w.GetWallapopItems()
	fmt.Println(items)
}
```

### Retrieving Wallapop reviews

```go
package main

import (
	"fmt"

	"github.com/your-package-import-path/wallapop"
)

func main() {
	// Create a new Wallapop instance
	w := wallapop.New("your-user-id")

	// Set the Wallapop reviews
	w.SetWallapopReviews()

	// Retrieve the Wallapop reviews
	reviews := w.GetWallapopReviews()
	fmt.Println(reviews)
}
```