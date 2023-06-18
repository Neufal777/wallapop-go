
# Documentation: wallapop-Go

This package provides a set of structures and methods for interacting with the Wallapop API and retrieving user profile information, items, and reviews.


```go

func main() {

	// Create a new Wallapop Object
	// Set the user ID, can be found in the URL of the user profile page
	// Example: https://es.wallapop.com/app/user/ocasionplusg-437879034-8j3y83q89169/published
	// The user ID in this case is: 8j3y83q89169

	// In case you want to retrieve specific information, you can use the following methods:
	wall := wallapop.New("8j3y83q89169")

	wall.SetWallapopItems()       // Optional
	wall.SetWallapopProfileInfo() // Optional
	wall.SetWallapopReviews()     // Optional

	// The code above is the same as:
	_ = wallapop.New("8j3y83q89169")
	wall.SetWallapopAll() // This method sets all the information (profile, items, reviews)

	// To retrieve the data set to the Wallapop Object, you can use the following methods:
	wall.GetWallapopItems()       // return a WallapopItems{...}
	wall.GetWallapopProfileInfo() // return WallapopProfileInfo{...}
	wall.GetWallapopReviews()     // return WallapopReviews{...}

	// Regarding the tables, you can display them in the CLI with the following methods:

	// Display the tables with the information, you can set the tables you want to display
	wall.CliWallapopPrint("profile", "items") // Options: profile, items, reviews

	// Display all, left empty, it will display all the tables
	wall.CliWallapopPrint()

	// Categories are identified by their ID, you can find the ID in Item object (CategoryID)
	// If you need to retrieve the category information, you can use the following function,
	// This will return a WallapopCategory{...} object with the information of the category
	category := wallapop.GetWallapopCategoryData(14000) // 14000 is the ID of the category

}

```

- **`Wallapop`**: Represents the main structure that holds user information, profile data, items, and reviews.

**Example:** 
```json
{
  "User": null,
  "WallapopUserProfile": null,
  "WallapopUserItems": null,
  "WallapopUserReviews": null 
}
```
## Methods

- **`New(UserID string) *Wallapop`**: Returns a new instance of the `Wallapop` object with the specified user ID (the ID will be used to get the different data we might need in the future).

**Example:**

```go
walla := wallapop.New("abc0982732")
```

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

- **`SetWallapopProfileInfo() *Wallapop`**: Sets the `WallapopProfileInfo` struct.

**Example:**

```go
wallapop.SetWallapopProfileInfo()
```

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

- **`SetWallapopItems(category ...string) *Wallapop`**: Sets the `WallapopItems` object for the specified category. If no category is provided, it sets the items for all categories. Currently, only one category can be set.

**Example:**

```go
wallapop.SetWallapopItems()
```


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

- **`SetWallapopReviews() *Wallapop`**: Sets the `WallapopReviews` struct.

**Example:**

```go
wallapop.SetWallapopReviews()
```


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
- **`SetWallapopAll() *Wallapop`**: Sets all the user information including profile, items, and reviews.

```go
wallapop.SetWallapopAll()
```

- **`GetWallapopProfileInfo() WallapopProfileInfo`**: Returns the `WallapopProfileInfo` containing user profile information.

   **Example:**
   
```go
wallapop.GetWallapopProfileInfo()
```

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

- **`GetWallapopItems() WallapopItems`**: Returns the `WallapopItems` object containing user items.

```go
wallapop.GetWallapopItems()
```

- **`GetWallapopReviews() WallapopReviews`**: Returns the `WallapopReviews` object containing user reviews.

```go
wallapop.GetWallapopReviews()
```

- **`CliWallapopPrint(displayMode ...string)`**: Prints user profile information, items, and reviews based on the provided display mode. The display mode determines which information to print. If no display mode is provided, it defaults to printing all information.

**Example(s):** Note: You can add as many fields as you need

```go
wallapop.CliWallapopPrint()
```

## profile:

```go
func main() {
    wall := wallapop.New("npj9q77050ze")

    // In case you:
    //      just need the profile info, you can use wallapop.SetWallapopProfileInfo()
    //      just need the items, you can use wallapop.SetWallapopItems()
    //      just need the reviews, you can use wallapop.SetWallapopReviews()
    wall.SetWallapopAll()

    // In case you:
    //      want all the information, leave the parameter empty
    //      just need the items, you can use wallapop.CliWallapopPrint("profile")
    //      just need the profile info, you can use wallapop.CliWallapopPrint("items")
    //      just need the reviews, you can use wallapop.CliWallapopPrint("reviews")
    //      need multiple information, you can set them, wallapop.CliWallapopPrint("profile", "products", etc.)
    wall.CliWallapopPrint("profile")
}

```
### *Output*


| ID           | NAME       | LOCATION CITY | URL SHARE                             | SINCE                   | ... | 
| :----------- | :--------- | :------------ | :------------------------------------ | :---------------------- | :--- |
| `abc012345`  | `Dracco`| `Barcelona`   | `http://p.wallapop.com/p/0123456789`  | `2018-07-17`| ... |


## Items:

```go
func main() {
    ...
    wall.CliWallapopPrint("items")
}

```
### *Output*

|           TITLE            |    PRICE    | CATEGORY |                                 URL             | ...|
| :------------------------: | :---------: | :------: | :---------------------------------------------: |  :-------: |
| Mercedes-Benz | 50000.00 €  |   100    | https://es.wallapop....-00000 | ...|
| Volkswagen      | 34000.00 €  |   100    | https://es.wallapop....-00000 | ... |


## Reviews:

```go
func main() {
    ...
    wall.CliWallapopPrint("reviews")
}

```
### *Output*

| CATEGORY ID |   USER ID    |         COMMENT         | SCORING | ... |
| :---------- | :---------- | :--------------------- | ------: | ----: |
|   12467     | xyz465789   |  Buena experiencia      |    80   | .... |
|   12467     | xyy789746   | Muy agradable y puntual |   100   | ... |

## All:

```go
func main() {
    ...
    wall.CliWallapopPrint() // leave empty
}

```

**Ouput: Display all the tables**
## Notes

- The package relies on the `httpRequest` package for making HTTP requests to the Wallapop API.
- The `Wallapop` object and its associated methods provide a convenient way to retrieve and manage user information, items, and reviews from Wallapop.

---
