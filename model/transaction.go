package model

// Credit is the model of credits provided to customers
type Credit struct {
	CustomerID int     `db:"customer_id" json:"customerid"`
	Date       string  `db:"date" json:"date"`
	Amount     float64 `db:"amount" json:"amount,string"`
	Remarks    *string `db:"remarks" json:"remarks"`
}

// Payment is the model of payments made by customers
type Payment struct {
	CustomerID int     `db:"customer_id" json:"customerid"`
	Date       string  `db:"date" json:"date"`
	Amount     float64 `db:"amount" json:"amount,string"`
	Mode       *string `db:"mode" json:"mode"`
	Remarks    *string `db:"remarks" json:"remarks"`
}
