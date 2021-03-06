package db

import (
	"log"

	"github.com/hmnayak/credit/model"
)

// UpsertCustomer updates customer, if not found inserts a new record
func (p *PostgresDb) UpsertCustomer(c model.Customer) {
	query :=
		`
		INSERT INTO customer (customer_id, org_id, name, email, phone_no, gstin) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (customer_id, org_id)
		DO UPDATE SET
			name = EXCLUDED.name,
			email = EXCLUDED.email,
			phone_no = EXCLUDED.phone_no,
			gstin = EXCLUDED.gstin
		`

	_, err := p.dbConn.Exec(query, c.CustomerID, c.OrganisationID, c.Name, c.Email, c.PhoneNumber, c.GSTIN)
	if err != nil {
		log.Printf("Error inserting customer: %v\n", err.Error())
	}
}

// GetCustomerCount gets the count of customer records
func (p *PostgresDb) GetCustomerCount() (count int, err error) {
	query :=
		`
		SELECT COUNT(*)
		FROM customer
		`
	err = p.dbConn.Get(&count, query)
	if err != nil {
		log.Printf("Error getting customers count: %v", err.Error())
	}
	return
}
