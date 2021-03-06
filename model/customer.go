package model

// Customer is the model of a person/business entity who receives credits and makes payments
type Customer struct {
	ID             int    `db:"id"`
	CustomerID     string `db:"customer_id" json:"customerid"`
	OrganisationID string `db:"organisation_id" json:"organisationid"`
	Name           string `db:"name" json:"name"`
	Email          string `db:"email" json:"email"`
	PhoneNumber    string `db:"phone_no" json:"phone"`
	GSTIN          string `db:"gstin" json:"gstin"`
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
