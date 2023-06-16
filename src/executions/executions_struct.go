package executions

type Result struct {
	UserID    string `json:"userId"`
	UserName  string `json:"userName"`
	ReviewAvg struct {
		Buy   float64 `json:"buy"`
		Sell  float64 `json:"sell"`
		Count int     `json:"count"`
	} `json:"reviewAvg"`
}

type UsersProcessed struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
