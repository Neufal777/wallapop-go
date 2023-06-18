
# Package Documentation: wallapop-Go

This package provides a set of structures and methods for interacting with the Wallapop API and retrieving user profile information, items, and reviews.

## Types

- `Wallapop`: Represents the main structure that holds user information, profile data, items, and reviews.

## Methods

- `New(UserID string) *Wallapop`: Returns a new instance of the `Wallapop` struct with the specified user ID (the ID will be used to get the different data we might need in the future).

**Example:**
```json
{
  "User": {
      "ID": "abc0123456"
  },
  "WallapopUserProfile": null,
  "WallapopUserItems": null,
  "WallapopUserReviews": null 
}
```

- `SetWallapopProfileInfo() *Wallapop`: Sets the `WallapopProfileInfo` struct.

**Example:**
```json
{
	"User": {
		"ID": "abc0123456"
	},
	"WallapopUserProfile": {
		"id": "abc0123456",
		"micro_name": "Dracco NF.",
		"type": "normal",
		"image": {...},
		"location": {...},
		"gender": "male",
		"web_slug": "dracco-0123456",
		"url_share": "http://p.wallapop.com/p/0123456?_pid=wi",
		"register_date": 1661424491000,
		"featured": false,
		"extra_info": {...},
		"language_code": "",
		"rating_average": 0,
		"rating_counts": 0,
		"is_visible": false,
		"reported_as_scam": false,
		"last_activity": 0,
		"scam_reports": 0,
		"has_scam_reports": false,
		"is_featured": false,
		"last_update": 0,
		"tickets_created": 0,
		"tickets_resolved": 0
	},
    "WallapopUserItems": null,
    "WallapopUserReviews": null
}
```

- `SetWallapopItems(category ...string) *Wallapop`: Sets the `WallapopItems` struct for the specified category. If no category is provided, it sets the items for all categories. Currently, only one category can be set.

**Example:**

```json
{
	"User": {
		"ID": "wzvy4v2l07zl"
	},
	"WallapopUserProfile": {...},
	"WallapopUserItems": {
		"data": [
			{
				"id": "qzmykegqqxzv",
				"title": "guitarra flamenca fustero ",
				"description": "clavijeros fustero! increíbles nuevos Clavijero Artesano, acabado en Negro Grafito, maquinaria remachada, palomilla metacrilato, imitación nacar.\nhago envíos por wallapop \núltimo precio no negociable!!!\nRegalo cejilla artesana ",
				"category_id": "12463",
				"slug": "guitarra-flamenca-909411232",
				"images": [...],
				"price": {...},
				"type_attributes": {...},
				"shipping": {...}
			},
            {...}
		],
	},
	"WallapopUserReviews": null
}
```

- `SetWallapopReviews() *Wallapop`: Sets the `WallapopReviews` struct.

**Example:**

```json
{
	"User": {
		"ID": "abcd0123456"
	},
	"WallapopUserProfile": {...},
	"WallapopUserItems": {...},
	"WallapopUserReviews": [
		{
			"item": {
				"category_id": 12467
			},
			"review": {
				"id": "mznge11vek6n",
				"date": 1652334358000,
				"scoring": 80,
				"is_shipping_transaction": false,
				"can_translate": false,
				"comments": "Buen comprador"
			},
			"user": {...},
			"type": "buy"
		},
        {...}
	]
}
```
- `SetWallapopAll() *Wallapop`: Sets all the user information including profile, items, and reviews.

- `GetWallapopProfileInfo() WallapopProfileInfo`: Returns the `WallapopProfileInfo` containing user profile information.

   **Example:**
   ```json
   {
     "id": "abcde0123456789",
     "micro_name": "Dracco NF.",
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

- `CliWallapopPrint(displayMode ...string)`: Prints user profile information, items, and reviews based on the provided display mode. The display mode determines which information to print. If no display mode is provided, it defaults to printing all information.


## Notes

- The package relies on the `httpRequest` package for making HTTP requests to the Wallapop API.
- The `Wallapop` struct and its associated methods provide a convenient way to retrieve and manage user information, items, and reviews from Wallapop.

---