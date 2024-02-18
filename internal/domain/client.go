package domain

type Client struct {
	ID      int `json:"id"`
	Limit   int `json:"limit"`
	Balance int `json:"balance"`
}
