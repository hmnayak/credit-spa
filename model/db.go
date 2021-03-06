package model

// Db specifies methods to manage resources
type Db interface {
	DoesUserExist(string) (bool, error)
	CreateUser(string, string) (string, error)

	GetOrganisationID(string) (string, error)

	GetCustomerCount() (int, error)
	UpsertCustomer(Customer) error
	GetAllCustomers(string) ([]*Customer, error)
	GetCustomer(string, string) (Customer, error)
}
