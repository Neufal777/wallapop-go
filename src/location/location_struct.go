package location

type Location struct {
	Address   string  `json:"address"`   // Street address of the location
	City      string  `json:"city"`      // City where the location is situated
	ZipCode   string  `json:"zipCode"`   // Postal/ZIP code of the location
	Longitude float64 `json:"longitude"` // Longitude of the location
	Latitude  float64 `json:"latitude"`  // Latitude of the location
}
