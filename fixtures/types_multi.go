package fixtures

type Rect struct {
	// Name of the Rect
	Name string `json:"name"`
	Width int `json:"width"`
	Height string `json:"height"`
}

type Circle struct {
	// Name of the Circle
	Name string `json:"name"`
	Diameter int `json:"diameter"`
}
