package model

// Customer is the model of a person/business entity who receives credits and makes payments
type Customer struct {
	ID            int       `db:"id" json:"id"`
	FullName      string    `db:"full_name" json:"fullname"`
	SearchName    string    `db:"search_name" json:"searchname"`
	DeliveryRoute string    `db:"delivery_route" json:"route"`
	Credits       []Credit  `db:"credits" json:"-"`
	Payments      []Payment `db:"payments" json:"-"`
	DueAmount     float64   `db:"due_amount" json:"dueamount"`
	CreditLimit   int       `db:"credit_limit" json:"creditlimit,string"`
	LatestCredit  float64   `json:"latestcredit"`
}
