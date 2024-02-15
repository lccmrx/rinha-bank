package dto

type StatementOutput struct {
	Balance      StatementBalanceOutput       `json:"saldo"`
	Transactions []StatementTransactionOutput `json:"ultimas_transacoes"`
}

type StatementBalanceOutput struct {
	Total         int    `json:"total"`
	StatementDate string `json:"data_extrato"`
	Limit         int    `json:"limite"`
}

type StatementTransactionOutput struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	Timestamp   string `json:"realizada_em"`
}
