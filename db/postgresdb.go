package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/hmnayak/credit/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used to specify postgres driver
)

// PostgresDb stores a connection to a postgres db
// also implements Db interface
type PostgresDb struct {
	dbConn *sqlx.DB
}

//InitDb creates a table in postgres using the configuration provided
func InitDb(connStr string) (*PostgresDb, error) {
	dbConn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	p := &PostgresDb{
		dbConn: dbConn,
	}
	err = p.dbConn.Ping()
	if err != nil {
		return nil, err
	}
	err = p.createTable()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PostgresDb) createTable() error {
	user :=
		`
			CREATE TABLE IF NOT EXISTS user_accounts (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				user_id 	TEXT UNIQUE,
				id_type		TEXT 
			)
		`

	organisation :=
		`
			CREATE TABLE IF NOT EXISTS organisations (
				id			BIGSERIAL NOT NULL PRIMARY KEY,
				org_id 		TEXT UNIQUE,
				owner_id	TEXT REFERENCES user_account(user_id)
			)
		`

	customer :=
		`
			CREATE TABLE IF NOT EXISTS customers (
				id 				BIGSERIAL NOT NULL PRIMARY KEY,
				customer_id 	TEXT NOT NULL UNIQUE,
				org_id 			TEXT NOT NULL UNIQUE,
				name			TEXT NOT NULL,
				email	  		TEXT,
				phone_no 		TEXT,
				gstin			integer UNIQUE, 
				CONSTRAINT fk_org_id FOREIGN KEY(org_id) REFERENCES organisation(org_id),
				CONSTRAINT unique_customer_org UNIQUE (customer_id, org_id)
			)	
		`

	createUserSessionTable := `
		CREATE TABLE IF NOT EXISTS usersession (
			token		VARCHAR	NOT NULL PRIMARY KEY,
			username 	VARCHAR NOT NULL,
			auth		VARCHAR NOT NULL
		)
	`

	createCreditTable := `
		CREATE TABLE IF NOT EXISTS credit (
			id 			SERIAL,
			customer_id SERIAL REFERENCES customer(id),
			amount 		NUMERIC,
			date		DATE,
			remarks		VARCHAR
		)
	`

	createPaymentTable := `
		CREATE TABLE IF NOT EXISTS payment (
			id 			SERIAL,
			customer_id SERIAL REFERENCES customer(id),
			amount 		NUMERIC,
			date		DATE,
			mode	    VARCHAR,
			remarks		VARCHAR
		)
	`

	_, err := p.dbConn.Exec(user)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(organisation)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(createUserSessionTable)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(customer)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(createCreditTable)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(createPaymentTable)
	if err != nil {
		return err
	}

	return nil
}

// Login returns the authorization assigned to the provided login credentials.
func (p *PostgresDb) Login(username string, password string) (string, error) {
	var authorization string
	query := `
		SELECT auth
		FROM useraccount
		WHERE name=$1 AND password=$2
	`
	err := p.dbConn.Get(&authorization, query, username, password)
	return authorization, err
}

// CreateUserSession creates an entry for the provided authentication token in usersession table.
func (p *PostgresDb) CreateUserSession(token string, username string, auth string) error {
	query := `
		INSERT INTO usersession (token, username, auth)
		VALUES ($1, $2, $3)
	`
	_, err := p.dbConn.Exec(query, token, username, auth)
	return err
}

// GetUserSession returns an authentication associated with the provided user if one exists.
func (p *PostgresDb) GetUserSession(username string) (model.AuthToken, error) {
	var token model.AuthToken
	query := `
		SELECT token, username, auth
		FROM usersession
		WHERE username=$1
	`
	err := p.dbConn.Get(&token, query, username)
	if err != nil {
		log.Println("Error getting user session:", err)
		return token, err
	}
	return token, nil
}

// DeleteUserSession creates an entry for the provided authentication token in usersession table.
func (p *PostgresDb) DeleteUserSession(token string) error {
	query := `
		DELETE FROM usersession 
		WHERE token=$1
	`
	_, err := p.dbConn.Exec(query, token)
	return err
}

// ValidateUser looks for the provided authentication token in the user session table.
func (p *PostgresDb) ValidateUser(token string) (model.AuthToken, error) {
	var user model.AuthToken
	query := `
		SELECT username, auth
		FROM usersession
		WHERE token=$1
	`
	err := p.dbConn.Get(&user, query, token)
	return user, err
}

// GetRoutes returns all delivery routes.
func (p *PostgresDb) GetRoutes() ([]string, error) {
	r := []string{}
	query := `
		SELECT DISTINCT delivery_route 
		FROM customer
		ORDER BY delivery_route
	`
	err := p.dbConn.Select(&r, query)
	if err != nil {
		return r, err
	}
	return r, nil
}

// GetCreditsByCustomer gets all credits received by a customer
func (p *PostgresDb) GetCreditsByCustomer(customerID int64) ([]*model.Credit, error) {
	credits := []*model.Credit{}
	query := `
		SELECT id, customer_id, date, amount, remarks
		FROM credit
		WHERE customer_id=$1
		ORDER BY date
	`
	rows, err := p.dbConn.Query(query, customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no payments found for customer id: \n")
		}
		return credits, err
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Credit
		if err := rows.Scan(&c.ID, &c.CustomerID, &c.Date, &c.Amount, &c.Remarks); err != nil {
			log.Panicln("error scanning payment row:", err)
		}
		date, err := time.Parse(time.RFC3339, c.Date)
		if err != nil {
			log.Println("error parsing date: ", err)
		}
		c.Date = date.Format("02-01-2006")
		credits = append(credits, &c)
	}

	return credits, nil
}

// UpdateCredit updates an existing credit
func (p *PostgresDb) UpdateCredit(c model.Credit) error {
	query := `
		UPDATE credit 
		SET amount=$1
		WHERE id=$2
	`
	_, err := p.dbConn.Exec(query, c.Amount, c.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCredit deletes an existing payment
func (p *PostgresDb) DeleteCredit(id int) error {
	query := `
		DELETE FROM credit 
		WHERE id=$1
	`
	_, err := p.dbConn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// CreatePayment creates a new payment entry
func (p *PostgresDb) CreatePayment(t model.Payment) error {
	query := `
		INSERT INTO payment (customer_id, amount, date, mode, remarks) 
		VALUES ($1, $2, $3, $4, $5) 
	`
	_, err := p.dbConn.Exec(query, t.CustomerID, t.Amount, t.Date, t.Mode, t.Remarks)
	if err != nil {
		return err
	}

	return nil
}

// GetPaymentsByCustomer gets all payments made by a customer
func (p *PostgresDb) GetPaymentsByCustomer(customerID int64) ([]*model.Payment, error) {
	payments := []*model.Payment{}
	query := `
		SELECT id, customer_id, date, amount, mode, remarks
		FROM payment
		WHERE customer_id=$1
		ORDER BY date
	`

	rows, err := p.dbConn.Query(query, customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no payments found for customer id: \n")
		}
		return payments, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Payment
		if err := rows.Scan(&p.ID, &p.CustomerID, &p.Date, &p.Amount, &p.Mode, &p.Remarks); err != nil {
			log.Panicln("error scanning payment row:", err)
		}
		date, err := time.Parse(time.RFC3339, p.Date)
		if err != nil {
			log.Println("error parsing date: ", err)
		}
		p.Date = date.Format("02-01-2006")
		payments = append(payments, &p)
	}

	return payments, nil
}

// UpdatePayment updates an existing payment
func (p *PostgresDb) UpdatePayment(pmt model.Payment) error {
	query := `
		UPDATE payment 
		SET amount=$1 
		WHERE id=$2
	`
	_, err := p.dbConn.Exec(query, pmt.Amount, pmt.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeletePayment deletes an existing payment
func (p *PostgresDb) DeletePayment(id int) error {
	query := `
		DELETE FROM payment 
		WHERE id=$1
	`
	_, err := p.dbConn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// GetDueAmount returns the due amount for a customer
func (p *PostgresDb) GetDueAmount(customerID int) (float64, error) {
	var dueAmount float64
	paymentsQuery := `
		SELECT COALESCE(SUM(amount), 0) 
		FROM payment 
		WHERE customer_id=$1 
		GROUP BY customer_id
	`
	var sumPayments float64
	err := p.dbConn.Get(&sumPayments, paymentsQuery, customerID)
	if err != nil {
		log.Println("Error getting sum payments:", err)
	}

	creditsQuery := `	
		SELECT COALESCE(SUM(amount), 0) 
		FROM credit 
		WHERE customer_id=$1 
		GROUP BY customer_id
	`
	var sumCredits float64
	err = p.dbConn.Get(&sumCredits, creditsQuery, customerID)
	if err != nil {
		log.Println("Error getting sum credits:", err)
	}
	dueAmount = sumCredits - sumPayments

	return dueAmount, nil
}

func (p *PostgresDb) getLastCreditDateByCreditor(c model.Customer) (string, error) {
	var lastCreditDate string
	query := `
		SELECT date 
		FROM credit 
		WHERE customer_id=$1 
		ORDER BY date DESC 
		LIMIT 1
	`
	err := p.dbConn.Get(&lastCreditDate, query, c.ID)
	if err != nil {
		log.Println("Error getting last credit date:", err)
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return lastCreditDate, err
}

func (p *PostgresDb) getLastPayDateByCreditor(c model.Customer) (string, error) {
	var lastPayDate string
	query := `
		SELECT date 
		FROM payment 
		WHERE customer_id=$1 
		ORDER BY date DESC 
		LIMIT 1
	`
	err := p.dbConn.Get(&lastPayDate, query, c.ID)
	if err != nil {
		log.Println("Error getting last payment date:", err)
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return lastPayDate, err
}
