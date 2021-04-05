package model

// Customer is the model of a person/business entity who receives credits and makes payments
type Customer struct {
	ID             int    `db:"id" json:"-"`
	CustomerID     string `db:"customer_id" json:"customerid"`
	OrganisationID string `db:"org_id" json:"-"`
	Name           string `db:"name" json:"name"`
	Email          string `db:"email" json:"email"`
	PhoneNumber    string `db:"phone_no" json:"phone"`
	GSTIN          string `db:"gstin" json:"gstin"`
}

type ListCustomersResponse struct {
	Customers     []*Customer `json:"customers"`
	NextPageToken int         `json:"nextpagetoken"`
	TotalSize     int         `json:"totalsize"`
}
