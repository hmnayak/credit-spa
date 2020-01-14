package model

// Transaction is the model of an transaction with an entity
type Transaction struct {
	CustomerID int     `db:"customer_id" json:"customerid"`
	Date       string  `db:"date" json:"date"`
	Amount     float64 `db:"amount" json:"amount"`
	Mode       *string  `db:"mode" json:"mode"`
	Remarks    *string `db:"remarks" json:"remarks"`
}
