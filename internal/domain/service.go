package domain

type Site struct {
	Site string `json:"site" example:"example.com"`
}

type Answer struct {
	Time    int64  `json:"time,omitempty" example:"1"`
	Message string `json:"message,omitempty" example:"n/a"`
}
