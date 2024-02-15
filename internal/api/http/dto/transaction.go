package dto

type TransactionInput struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type TransactionOutput struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}
