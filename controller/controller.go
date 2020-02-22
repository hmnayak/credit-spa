package controller

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/hmnayak/credit/db"
	"github.com/hmnayak/credit/model"
)

// Controller processes http requests.
// It contains a reference to the data store to perform CRUD operations.
type Controller struct {
	model      *model.Model
	authSecret string
}

// Init sets up a connection to database with configuration provided
func (c *Controller) Init(dbConfig db.Config, authSecret string) error {
	db, err := db.InitDb(dbConfig)
	if err != nil {
		log.Println("Error InitDb: %v", err)
		return err
	}

	c.model = model.New(db)
	c.authSecret = authSecret

	return nil
}

// Login verifies a login attempt with the provided credentials, returns auth token if successful
func (c *Controller) Login(username string, password string) (model.AuthToken, error) {
	var authToken model.AuthToken
	auth, err := c.model.Db.Login(username, password)
	if err != nil {
		log.Println("Error logging in:", err)
		return authToken, err
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":      username,
		"authorization": auth,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(c.authSecret))
	if err != nil {
		log.Println("Error signing auth token:", err)
		return authToken, err
	}

	a, err := c.model.Db.GetUserSession(username)
	if err == nil {
		return a, nil
	}

	err = c.model.Db.CreateUserSession(tokenString, username, auth)
	if err != nil {
		log.Println("Error creating user session:", err)
		return authToken, err
	}

	authToken.Token = tokenString
	authToken.UserName = username
	authToken.Authorization = auth
	return authToken, nil
}

// Logout invalidates the authentication token provided.
func (c *Controller) Logout(token string) error {
	err := c.model.Db.DeleteUserSession(token)
	if err != nil {
		log.Println("Error deleting user session:", err)
	}
	return err
}

// ValidateUser confirms the validity of authentication tokens
func (c *Controller) ValidateUser(tokenStr string) (string, error) {
	var authorization string
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.authSecret), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return authorization, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		a, err := c.model.Db.ValidateUser(tokenStr)
		if err != nil {
			log.Println("Error validating user:", err)
			return authorization, err
		}
		if a.UserName == claims["username"] && a.Authorization == claims["authorization"] {
			return claims["authorization"].(string), nil
		}
	} else {
		return authorization, jwt.ErrSignatureInvalid
	}

	return authorization, nil
}

// GetAllRoutes returns a list of all delivery routes
func (c *Controller) GetAllRoutes() ([]string, error) {
	r, err := c.model.Db.GetRoutes()
	if err != nil {
		log.Println("Error GetAllRoutes:", err)
	}

	return r, err
}

// GetCreditorsOnRoute returns all creditors on a given delivery route
func (c *Controller) GetCreditorsOnRoute(route string) ([]*model.Customer, error) {
	creditors, err := c.model.Db.GetCustomersOnRoute(route)
	if err != nil {
		log.Println("Error GetCustomersOnRoute:", err)
	}

	return creditors, err
}

// GetAllCreditors returns a list of all creditors
func (c *Controller) GetAllCreditors() ([]*model.Customer, error) {
	creditors, err := c.model.Db.GetAllCustomers()
	if err != nil {
		log.Println("Error GetAllCreditors:", err)
		return creditors, err
	}

	return creditors, nil
}

// GetCreditorByID returns the representation of a creditor with the given ID
func (c *Controller) GetCreditorByID(id int64) (*model.Customer, error) {
	creditor, err := c.model.Db.GetCustomerByID(id)
	if err != nil {
		log.Println("Error GetCreditorByID:", err)
	}

	return creditor, err
}

// GetCreditorByNameRoute returns the representation of a creditor with the name and delivery route provided
func (c *Controller) GetCreditorByNameRoute(route string, name string) (*model.Customer, error) {
	creditor, err := c.model.Db.GetCustomerByNameRoute(route, name)
	if err != nil {
		log.Println("Error GetCreditorByNameRoute:", err)
	}

	return creditor, err
}

// CreateCreditor stores a new creditor and returns their id
func (c *Controller) CreateCreditor(creditor model.Customer) (int64, error) {
	id, err := c.model.Db.CreateCustomer(creditor)
	if err != nil {
		log.Println("Error CreateCustomer:", err)
	}

	return id, err
}

// CreateCredit stores a new credit transaction
func (c *Controller) CreateCredit(credit model.Credit) error {
	err := c.model.Db.CreateCredit(credit)
	if err != nil {
		log.Println("Error CreateCredit:", err)
	}

	return err
}

// GetCreditsByCreditor returns all payments made by given creditor
func (c *Controller) GetCreditsByCreditor(id int64) ([]*model.Credit, error) {
	credits, err := c.model.Db.GetCreditsByCustomer(id)
	if err != nil {
		log.Println("Error GetCreditsByCreditor:", err)
	}

	return credits, err
}

// UpdateCredit updates an existing new credit transaction
func (c *Controller) UpdateCredit(credit model.Credit) error {
	err := c.model.Db.UpdateCredit(credit)
	if err != nil {
		log.Println("Error UpdateCredit:", err)
	}

	return err
}

// DeleteCredit deletes an existing credit transaction
func (c *Controller) DeleteCredit(id int) error {
	err := c.model.Db.DeleteCredit(id)
	if err != nil {
		log.Println("Error UpdateCredit:", err)
	}

	return err
}

// CreatePayment stores a new payment transaction
func (c *Controller) CreatePayment(payment model.Payment) error {
	err := c.model.Db.CreatePayment(payment)
	if err != nil {
		log.Println("Error CreatePayment:", err)
	}

	return err
}

// GetPaymentsByCreditor returns all payments made by given creditor
func (c *Controller) GetPaymentsByCreditor(id int64) ([]*model.Payment, error) {
	p, err := c.model.Db.GetPaymentsByCustomer(id)
	if err != err {
		log.Println("Error GetPaymentsByCreditor:", err)
	}

	return p, err
}

// UpdatePayment updates an existing payment transaction
func (c *Controller) UpdatePayment(payment model.Payment) error {
	err := c.model.Db.UpdatePayment(payment)
	if err != nil {
		log.Println("Error UpdatePayment:", err)
	}

	return err
}

// DeletePayment deletes an existing payment transaction
func (c *Controller) DeletePayment(id int) error {
	err := c.model.Db.DeletePayment(id)
	if err != nil {
		log.Println("Error UpdatePayment:", err)
	}

	return err
}

// GetAllDefaulters returns all creditors whose due amount exceeds their credit limit
func (c *Controller) GetAllDefaulters() ([]*model.Customer, error) {
	d, err := c.model.Db.GetAllDefaulters()
	if err != nil {
		log.Println("Error GetAllDefaulters:", err)
	}

	return d, err
}

// GetAllDefaultersNew returns all creditors whose due amount exceeds their credit limit
func (c *Controller) GetAllDefaultersNew() ([]*model.Defaulter, error) {
	d, err := c.model.Db.GetAllDefaultersNew()
	if err != nil {
		log.Println("Error GetAllDefaultersNew:", err)
	}

	return d, err
}
