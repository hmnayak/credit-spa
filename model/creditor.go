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
	PaymentCycle  int       `db:"pay_cycle" json:"paycycle"`
	LatestCredit  float64   `json:"latestcredit"`
}

// Defaulter is the model of a person/business entity who has defaulted on their payment
type Defaulter struct {
	ID             int     `json:"id"`
	FullName       string  `json:"fullname"`
	SearchName     string  `json:"searchname"`
	DeliveryRoute  string  `json:"route"`
	DueAmount      float64 `json:"dueamount"`
	LatestCredit   float64 `json:"latestcredit"`
	DueFrom        string  `json:"duefrom"`
	PaymentInCycle float64 `json:"paymentincycle"`
	LastPaidOn     string  `json:"lastpaidon"`
	PayCycle       string  `json:"paycycle"`
	DefaultCause   string  `json:"defaultcause"`
}
