package model

// Db specifies methods to manage resources
type Db interface {
	Login(string, string) (string, error)
	CreateUserSession(string, string, string) error
	GetUserSession(string) (AuthToken, error)
	DeleteUserSession(string) error
	ValidateUser(string) (AuthToken, error)
	GetRoutes() ([]string, error)

	UpdateCredit(Credit) error
	DeleteCredit(int) error
	CreatePayment(Payment) error
	UpdatePayment(Payment) error
	DeletePayment(int) error
	GetCreditsByCustomer(int64) ([]*Credit, error)
	GetPaymentsByCustomer(int64) ([]*Payment, error)

	DoesUserExist(string) (bool, error)
	CreateUser(string, string) (string, error)

	GetOrganisationID(string) (string, error)

	GetCustomerCount() (count int, err error)
	UpsertCustomer(Customer) (err error)
}
