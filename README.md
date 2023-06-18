**Package Documentation: wallapop-Go**

This package provides a set of structures and methods for interacting with the Wallapop API and retrieving user profile information, items, and reviews.

**Types:**

- `Wallapop`: Represents the main structure that holds user information, profile data, items, and reviews.

**Methods:**

- `New(UserID string) *Wallapop`: Returns a new instance of the `Wallapop` struct with the specified user ID.

- `CliWallapopPrint(displayMode ...string)`: Prints user profile information, items, and reviews based on the provided display mode. The display mode determines which information to print. If no display mode is provided, it defaults to printing all information.

**Get information**
- `GetWallapopProfileInfo() WallapopProfileInfo`: Returns the `WallapopProfileInfo` struct containing user profile information.
```json
{
  "id": "abcde0123456789",
  "micro_name": "Dracco .",
  "type": "normal",
  "image": {
    "id": "xyz987654321",
    "original_width": 1600,
    "original_height": 1600,
    "average_hex_color": "f1f4f6",
    "urls_by_size": {...}
  },
  "location": {...},
  "gender": "male",
  "web_slug": "example-0123456789",
  "url_share": "http://p.wallapop.com/p/0123456789?_pid=wi",
  "register_date": 1637501205000,
  "featured": true,
  "extra_info": {...}
}

```
- `GetWallapopItems() WallapopItems`: Returns the `WallapopItems` struct containing user items.

- `GetWallapopReviews() WallapopReviews`: Returns the `WallapopReviews` struct containing user reviews.

**Set information**
- `SetWallapopItems(category ...string) *Wallapop`: Sets the `WallapopItems` struct for the specified category. If no category is provided, it sets the items for all categories. Currently, only one category can be set.

- `SetWallapopProfileInfo() *Wallapop`: Sets the `WallapopProfileInfo` struct.

- `SetWallapopReviews() *Wallapop`: Sets the `WallapopReviews` struct.

- `SetWallapopAll() *Wallapop`: Sets all the user information including profile, items, and reviews.

**Notes:**

- The package relies on the `httpRequest` package for making HTTP requests to the Wallapop API.
- The `Wallapop` struct and its associated methods provide a convenient way to retrieve and manage user information, items, and reviews from Wallapop.