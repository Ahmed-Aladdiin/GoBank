package types

type Transaction struct {
	AccountFrom int     `json:"accountFrom"`
	AccountTo   int     `json:"accountTo"`
	Amount      float64 `json:"amount"`
}
