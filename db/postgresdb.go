package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

// Config stores the connection string to connect to a db instance
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func (c Config) String() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
			c.Host, c.Port, c.User, c.DBName, c.SSLMode)
	}
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password, c.SSLMode)
}

//InitDb creates a table in postgres using the configuration provided
func InitDb(cfg Config) (*PostgresDb, error) {
	dbConn, err := sqlx.Connect("postgres", fmt.Sprintf("%v", cfg))
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
	createUserTable := `
		CREATE TABLE IF NOT EXISTS useraccount (
			id			SERIAL	NOT NULL PRIMARY KEY,
			name 	 	VARCHAR NOT NULL,
			password 	VARCHAR NOT NULL,
			auth 		VARCHAR NOT NULL
		)
	`

	createUserSessionTable := `
		CREATE TABLE IF NOT EXISTS usersession (
			token		VARCHAR	NOT NULL PRIMARY KEY,
			username 	VARCHAR NOT NULL,
			auth		VARCHAR NOT NULL
		)
	`

	createCustomerTable := `
		CREATE TABLE IF NOT EXISTS customer (
			id 			  	SERIAL        NOT NULL PRIMARY KEY,
			full_name       VARCHAR	NOT NULL,
			search_name	  	VARCHAR NOT NULL,
			delivery_route 	VARCHAR,
			credit_limit	integer,
			UNIQUE (search_name, delivery_route)
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

	_, err := p.dbConn.Exec(createUserTable)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(createUserSessionTable)
	if err != nil {
		return err
	}

	_, err = p.dbConn.Exec(createCustomerTable)
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

// CreateCustomer creates a new customer in customer table
func (p *PostgresDb) CreateCustomer(c model.Customer) (int64, error) {
	var newID int64
	query := `
		INSERT INTO	customer (full_name, search_name, delivery_route, credit_limit) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := p.dbConn.QueryRow(query,
		c.FullName, c.SearchName, strings.ToLower(c.DeliveryRoute), c.CreditLimit).Scan(&newID)
	return newID, err
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

// GetCustomersOnRoute returns all customers on a routes
func (p *PostgresDb) GetCustomersOnRoute(r string) ([]*model.Customer, error) {
	creditors := []*model.Customer{}
	query := `
		SELECT id, full_name, search_name, delivery_route, credit_limit
		FROM customer
		WHERE delivery_route=$1
		ORDER BY search_name
	`
	rows, err := p.dbConn.Query(query, r)
	if err != nil {
		return creditors, err
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Customer
		if err := rows.Scan(&c.ID, &c.FullName, &c.SearchName, &c.DeliveryRoute, &c.CreditLimit); err != nil {
			log.Panicln("error scanning creditor row:", err)
		}
		creditors = append(creditors, &c)
	}

	for _, c := range creditors {
		due, err := p.GetDueAmount(*c)
		if err != nil {
			log.Println("Error getting due amount:", err)
		}
		c.DueAmount = due
	}

	return creditors, nil
}

// GetAllCustomers gets the list of all customers
func (p *PostgresDb) GetAllCustomers() ([]*model.Customer, error) {
	creditors := []*model.Customer{}
	query := `
		SELECT id, full_name, search_name, delivery_route, credit_limit 
		FROM customer
	`
	err := p.dbConn.Select(&creditors, query)
	if err != nil {
		return nil, err
	}

	for _, c := range creditors {
		due, err := p.GetDueAmount(*c)
		if err != nil {
			log.Println("Error getting due amount:", err)
		}
		c.DueAmount = due
	}

	return creditors, err
}

// GetCustomerByID gets a single customer with the id provided
func (p *PostgresDb) GetCustomerByID(id int64) (*model.Customer, error) {
	c := model.Customer{}
	query := `
		SELECT id, full_name, search_name, delivery_route, credit_limit  
		FROM customer 
		WHERE id=$1
	`
	err := p.dbConn.Get(&c, query, id)
	if err != nil {
		return &c, err
	}

	due, err := p.GetDueAmount(c)
	if err != nil {
		log.Println("Error getting due amount:", err)
		return &c, err
	}
	c.DueAmount = due

	credits, err := p.GetCreditsByCustomer(id)
	if err != nil {
		log.Println("Error getting latest credit:", err)
		return &c, err
	}
	c.LatestCredit = credits[len(credits)-1].Amount

	return &c, nil
}

// GetCustomerByNameRoute gets a single customer whose route and name is provided
func (p *PostgresDb) GetCustomerByNameRoute(route string, name string) (*model.Customer, error) {
	c := model.Customer{}
	query := `
		SELECT id, full_name, search_name, delivery_route, credit_limit  
		FROM customer 
		WHERE delivery_route=$1 AND search_name=$2
	`
	err := p.dbConn.Get(&c, query, route, name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no customer with delivery route: %v, search_name: %v \n", route, name)
		}
		return &c, err
	}

	due, err := p.GetDueAmount(c)
	if err != nil {
		log.Println("Error getting due amount:", err)
	}
	c.DueAmount = due

	return &c, nil
}

// CreateCredit creates a new credit entry
func (p *PostgresDb) CreateCredit(t model.Credit) error {
	query := `
		INSERT INTO credit (customer_id, amount, date, remarks) 
		VALUES ($1, $2, $3, $4) 
	`
	_, err := p.dbConn.Exec(query, t.CustomerID, t.Amount, t.Date, t.Remarks)
	if err != nil {
		return err
	}

	return nil
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
			log.Printf("no payments found for customer id: %v \n")
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
			log.Printf("no payments found for customer id: %v \n")
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
func (p *PostgresDb) GetDueAmount(c model.Customer) (float64, error) {
	var dueAmount float64
	paymentsQuery := `
		SELECT SUM(amount) 
		FROM payment 
		WHERE customer_id=$1 
		GROUP BY customer_id
	`
	var sumPayments float64
	err := p.dbConn.Get(&sumPayments, paymentsQuery, c.ID)
	if err != nil {
		log.Println("Error getting sum payments:", err)
	}

	creditsQuery := `	
		SELECT SUM(amount) 
		FROM credit 
		WHERE customer_id=$1 
		GROUP BY customer_id
	`
	var sumCredits float64
	err = p.dbConn.Get(&sumCredits, creditsQuery, c.ID)
	if err != nil {
		log.Println("Error getting sum credits:", err)
	}
	dueAmount = sumCredits - sumPayments

	return dueAmount, nil
}

// GetAllDefaulters gets a list of creditors that are defaulting on payments
func (p *PostgresDb) GetAllDefaulters() ([]*model.Customer, error) {
	defaulters := []*model.Customer{}

	all, err := p.GetAllCustomers()
	if err != nil {
		log.Println("Error getting all customers: ", err)
		return defaulters, err
	}

	for _, c := range all {
		due, err := p.GetDueAmount(*c)
		if err != nil {
			log.Println("Error getting due amount for customer:", err)
			return defaulters, err
		}
		if due > float64(c.CreditLimit) {
			c.DueAmount = due
			defaulters = append(defaulters, c)
		}
	}

	return defaulters, nil
}
