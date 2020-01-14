package model

// Db specifies methods to manage resources
type Db interface {
	GetRoutes() ([]string, error)
	GetCustomersOnRoute(r string) ([]*Customer, error)
	GetAllCustomers() ([]*Customer, error)
	GetCustomerByID(int64) (*Customer, error)
	GetCustomerByNameRoute(string, string) (*Customer, error)
	CreateCustomer(Customer) (int64, error)
	CreateCredit(Transaction) error
	CreatePayment(Transaction) error
	GetCreditsByCustomer(int64) ([]*Transaction, error)
	GetPaymentsByCustomer(int64) ([]*Transaction, error)
}
