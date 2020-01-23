package db

import (
	"fmt"
	"log"

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
	createCustomerTable := `
		CREATE TABLE IF NOT EXISTS customer (
			id 			  	SERIAL        NOT NULL PRIMARY KEY,
			name          	VARCHAR(40)	NOT NULL,
			short_name	  	VARCHAR(5),
			delivery_route 	VARCHAR(5),
			contact			integer[],
			UNIQUE (short_name, delivery_route)
		);
	`

	createCreditTable := `
		CREATE TABLE IF NOT EXISTS credit (
			id 			SERIAL,
			customer_id SERIAL REFERENCES customer(id),
			amount 		NUMERIC,
			date		DATE,
			mode	    VARCHAR,
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

	_, err := p.dbConn.Exec(createCustomerTable)
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

// GetAllCustomers gets the list of all customers
func (p *PostgresDb) GetAllCustomers() ([]*model.Customer, error) {
	customers := make([]*model.Customer, 0)
	query := `
		SELECT id, name, short_name, delivery_route 
		FROM customer
	`
	err := p.dbConn.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	for _, c := range customers {
		credits := []model.Transaction{}
		creditsStmt := `
			SELECT amount, date
			FROM credit
			WHERE customer_id=$1
		`
		err = p.dbConn.Select(&credits, creditsStmt, c.ID)
		if err != nil {
			log.Println("Error getting credits:", err)
		}
		c.Credits = credits

		payments := []model.Transaction{}
		paymentsStmt := `
			SELECT amount, date
			FROM payment
			WHERE customer_id=$1
		`
		err = p.dbConn.Select(&payments, paymentsStmt, c.ID)
		if err != nil {
			log.Println("Error getting credits:", err)
		}
		c.Payments = payments
	}

	return customers, nil
}

// GetCustomerByID gets a single customer with the id provided
func (p *PostgresDb) GetCustomerByID(id int64) (*model.Customer, error) {
	c := model.Customer{}
	query := `
		SELECT id, name, short_name, delivery_route 
		FROM customer 
		WHERE id=$1
	`
	err := p.dbConn.Get(&c, query, id)
	if err != nil {
		return &c, err
	}

	return &c, nil
}

// GetCustomerByNameRoute gets a single customer whose route and name is provided
func (p *PostgresDb) GetCustomerByNameRoute(route string, name string) (*model.Customer, error) {
	c := model.Customer{}
	query := `
		SELECT id, name, short_name, delivery_route 
		FROM customer 
		WHERE delivery_route=$1 AND short_name=$2
	`
	err := p.dbConn.Get(&c, query, route, name)
	if err != nil {
		return &c, err
	}

	due, err := p.GetDueAmount(c)
	if err != nil {
		log.Println("Error getting due amount:", err)
	}
	c.DueAmount = due

	return &c, nil
}

// CreateCustomer creates a new customer in customer table
func (p *PostgresDb) CreateCustomer(c model.Customer) (int64, error) {
	var newID int64
	query := `
		INSERT INTO	customer (name, short_name, delivery_route) 
		VALUES ($1, $2, $3)
	`
	res, err := p.dbConn.Exec(query, c.Name, c.ShortName, c.DeliveryRoute)
	if err != nil {
		return newID, err
	}
	newID, err = res.LastInsertId()
	return newID, nil
}

// CreateCredit creates a new credit entry
func (p *PostgresDb) CreateCredit(t model.Transaction) error {
	query := `
		INSERT INTO credit (customer_id, amount, date, mode, remarks) 
		VALUES ($1, $2, $3, $4, $5) 
	`
	_, err := p.dbConn.Exec(query, t.CustomerID, t.Amount, t.Date, t.Mode, t.Remarks)
	if err != nil {
		return err
	}

	return nil
}

// CreatePayment creates a new payment entry
func (p *PostgresDb) CreatePayment(t model.Transaction) error {
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

// GetRoutes returns all delivery routes
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
	c := []*model.Customer{}
	query := `
		SELECT id, name, short_name, delivery_route
		FROM customer
		WHERE delivery_route=$1
		ORDER BY short_name
	`
	err := p.dbConn.Select(&c, query, r)
	if err != nil {
		return c, err
	}
	// for i := range c {
	// 	due, err := p.GetDueAmount(*c[i])
	// 	if err != nil {
	// 		log.Println("Error getting due amount:", err)
	// 	}
	// 	c[i].DueAmount = due
	// }

	return c, nil
}

// GetCreditsByCustomer gets all credits received by a customer
func (p *PostgresDb) GetCreditsByCustomer(customerID int64) ([]*model.Transaction, error) {
	credits := []*model.Transaction{}
	query := `
		SELECT customer_id, amount, date, mode, remarks
		FROM credit
		WHERE customer_id=$1
	`
	err := p.dbConn.Select(&credits, query, customerID)
	if err != nil {
		return credits, err
	}

	return credits, nil
}

// GetPaymentsByCustomer gets all payments made by a customer
func (p *PostgresDb) GetPaymentsByCustomer(customerID int64) ([]*model.Transaction, error) {
	payments := []*model.Transaction{}
	query := `
		SELECT customer_id, amount, date, mode, remarks
		FROM payment
		WHERE customer_id=$1
	`
	err := p.dbConn.Select(&payments, query, customerID)
	if err != nil {
		return payments, err
	}

	return payments, nil
}
