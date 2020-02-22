package model

// Db specifies methods to manage resources
type Db interface {
	Login(string, string) (string, error)
	CreateUserSession(string, string, string) error
	GetUserSession(string) (AuthToken, error)
	DeleteUserSession(string) error
	ValidateUser(string) (AuthToken, error)
	GetRoutes() ([]string, error)
	GetCustomersOnRoute(string) ([]*Customer, error)
	GetAllCustomers() ([]*Customer, error)
	GetCustomerByID(int64) (*Customer, error)
	GetCustomerByNameRoute(string, string) (*Customer, error)
	CreateCustomer(Customer) (int64, error)
	CreateCredit(Credit) error
	UpdateCredit(Credit) error
	DeleteCredit(int) error
	CreatePayment(Payment) error
	UpdatePayment(Payment) error
	DeletePayment(int) error
	GetCreditsByCustomer(int64) ([]*Credit, error)
	GetPaymentsByCustomer(int64) ([]*Payment, error)
	GetAllDefaulters() ([]*Customer, error)
	GetAllDefaultersNew() ([]*Defaulter, error)
}
