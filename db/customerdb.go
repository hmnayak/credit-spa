package db

import (
	"database/sql"
	"log"

	"github.com/hmnayak/credit/model"
)

// UpsertCustomer updates customer, if not found inserts a new record
func (p *PostgresDb) UpsertCustomer(c model.Customer) (err error) {
	query :=
		`
			INSERT INTO customers (customer_id, org_id, name, email, phone_no, gstin) 
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (customer_id, org_id)
			DO UPDATE SET
				name = EXCLUDED.name,
				email = EXCLUDED.email,
				phone_no = EXCLUDED.phone_no,
				gstin = EXCLUDED.gstin
		`

	_, err = p.dbConn.Exec(query, c.CustomerID, c.OrganisationID, c.Name, c.Email, c.PhoneNumber, c.GSTIN)
	if err != nil {
		log.Printf("Error inserting customer: %v\n", err.Error())
	}
	return
}

// GetLatestCustomerID gets the id assigned to last inserted customer
func (p *PostgresDb) GetLatestCustomerID(orgID string) (customerID string, err error) {
	var nullableID sql.NullString
	query :=
		`
			SELECT MAX(customer_id)
			FROM customers
			WHERE org_id = $1
		`
	err = p.dbConn.Get(&nullableID, query, orgID)
	if err != nil {
		log.Printf("Error getting latest customerID: %v", err.Error())
	}
	customerID = nullableID.String
	return
}

// GetCustomersCount gets the total number customers
func (p *PostgresDb) GetCustomersCount(orgID string) (count int, err error) {
	query :=
		`
			SELECT COUNT(*)
			FROM customers
			WHERE org_id = $1
		`

	err = p.dbConn.Get(&count, query, orgID)
	if err != nil {
		log.Printf("Error getting customers count: %v", err.Error())
	}
	return
}

// GetCustomersPaginated gets a paginated list of customers
func (p *PostgresDb) GetCustomersPaginated(orgID string, pageToken int, pageSize int) (customers []*model.Customer, err error) {
	customers = []*model.Customer{}
	query :=
		`
			SELECT customer_id, name, email, phone_no, gstin
			FROM customers
			WHERE org_id = $1
			ORDER BY customer_id ASC
			LIMIT $2 OFFSET $3
		`
	offset := 0
	if pageToken > 0 {
		offset = (pageToken - 1) * pageSize
	}
	err = p.dbConn.Select(&customers, query, orgID, pageSize, offset)
	if err != nil {
		log.Printf("Error getting paginated list of customers: %v", err.Error())
	}
	return
}

// GetAllCustomers gets a list of all customers
func (p *PostgresDb) GetAllCustomers(orgID string) (customers []*model.Customer, err error) {
	customers = []*model.Customer{}
	query :=
		`
			SELECT customer_id, name, email, phone_no, gstin
			FROM customers
			WHERE org_id = $1
			ORDER BY customer_id ASC
		`
	err = p.dbConn.Select(&customers, query, orgID)
	if err != nil {
		log.Printf("Error getting all customers: %v", err.Error())
	}
	return
}

// GetCustomer gets customer with specified customer id
func (p *PostgresDb) GetCustomer(customerID string, orgID string) (customer model.Customer, err error) {
	query :=
		`
			SELECT customer_id, name, email, phone_no, gstin
			FROM customers
			WHERE customer_id = $1 AND org_id = $2
		`
	err = p.dbConn.Get(&customer, query, customerID, orgID)
	if err != nil {
		log.Printf("Error getting customer with ID - %v: %v", customerID, err.Error())
	}
	return
}
