package controller

import (
	"log"

	"github.com/hmnayak/credit/db"
	"github.com/hmnayak/credit/model"
)

// Controller processes http requests.
// It contains a reference to the data store to perform CRUD operations.
type Controller struct {
	model *model.Model
}

// Init sets up a connection to database with configuration provided
func (c *Controller) Init(dbConfig db.Config) error {
	db, err := db.InitDb(dbConfig)
	if err != nil {
		log.Panicln("Error InitDb: %v", err)
	}

	c.model = model.New(db)

	return err
}

// GetAllRoutes returns a list of all delivery routes
func (c *Controller) GetAllRoutes() ([]string, error) {
	r, err := c.model.Db.GetRoutes()
	if err != nil {
		log.Panicln("Error GetAllRoutes:", err)
	}

	return r, err
}

// GetCreditorsOnRoute returns all creditors on a given delivery route
func (c *Controller) GetCreditorsOnRoute(route string) ([]*model.Customer, error) {
	creditors, err := c.model.Db.GetCustomersOnRoute(route)
	if err != nil {
		log.Panicln("Error getCustomersOnRoute:", err)
	}

	return creditors, err
}

// GetAllCreditors returns a list of all creditors
func (c *Controller) GetAllCreditors() ([]*model.Customer, error) {
	creditors, err := c.model.Db.GetAllCustomers()
	if err != nil {
		log.Panicln("Error SelectCustomers:", err)
	}

	for _, c := range creditors {
		c.CalculateDueAmount()
	}

	return creditors, err
}

// GetCreditorByID returns the representation of a creditor with the given ID
func (c *Controller) GetCreditorByID(id int64) (*model.Customer, error) {
	creditor, err := c.model.Db.GetCustomerByID(id)
	if err != nil {
		log.Panicln("Error SelectCustomers:", err)
	}

	return creditor, err
}

// GetCreditorByNameRoute returns the representation of a creditor with the name and delivery route provided
func (c *Controller) GetCreditorByNameRoute(route string, name string) (*model.Customer, error) {
	creditor, err := c.model.Db.GetCustomerByNameRoute(route, name)
	if err != nil {
		log.Panicln("Error SelectCustomers:", err)
	}

	return creditor, err
}

// CreateCreditor stores a new creditor and returns their id
func (c *Controller) CreateCreditor(creditor model.Customer) (int64, error) {
	id, err := c.model.Db.CreateCustomer(creditor)
	if err != nil {
		log.Panicln("Error CreateCustomer:", err)
	}

	return id, err
}

// CreateCredit stores a new credit transaction
func (c *Controller) CreateCredit(credit model.Transaction) error {
	err := c.model.Db.CreateCredit(credit)
	if err != nil {
		log.Panicln("Error CreateCredit:", err)
	}

	return err
}

// CreatePayment stores a new payment transaction
func (c *Controller) CreatePayment(payment model.Transaction) error {
	err := c.model.Db.CreatePayment(payment)
	if err != nil {
		log.Panicln("Error CreatePayment:", err)
	}

	return err
}

// GetPaymentsByCreditor returns all payments made by given creditor
func (c *Controller) GetPaymentsByCreditor(id int64) ([]*model.Transaction, error) {
	p, err := c.model.Db.GetPaymentsByCustomer(id)
	if err != nil {
		log.Panicln("Error GetPaymentsByCreditor:", err)
	}

	return p, err
}

// GetCreditsByCreditor returns all payments made by given creditor
func (c *Controller) GetCreditsByCreditor(id int64) ([]*model.Transaction, error) {
	credits, err := c.model.Db.GetCreditsByCustomer(id)
	if err != nil {
		log.Panicln("Error GetCreditsByCreditor:", err)
	}

	return credits, err
}
