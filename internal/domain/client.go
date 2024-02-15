package domain

type Client struct {
	ID           int `json:"id"`
	AccountLimit int `json:"limit"`
	Balance      int `json:"balance"`
}
