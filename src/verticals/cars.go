package verticals

type Car struct {
	Brand   string `json:"brand"`   //Brand of the car
	Model   string `json:"model"`   //Model of the car
	Year    int    `json:"year"`    //Year of the car
	Km      int    `json:"km"`      //Kilometers of the car
	Gearbox string `json:"gearbox"` //Gearbox of the car
	Price   int    `json:"price"`   //Price of the car
	Fuel    string `json:"fuel"`    //Fuel of the car
}
