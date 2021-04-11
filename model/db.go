package model

// Db specifies methods to manage resources
type Db interface {
	DoesUserExist(string) (bool, error)
	CreateUser(string, string) (string, error)

	GetOrganisationID(string) (string, error)

	UpsertCustomer(Customer) error
	GetLatestCustomerID(string) (string, error)
	GetCustomersPaginated(string, int, int) ([]*Customer, error)
	GetAllCustomers(string) ([]*Customer, error)
	GetCustomer(string, string) (Customer, error)
	GetCustomersCount(string) (int, error)

	UpsertItem(Item) error
	GetLatestItemID(string) (string, error)
	GetItemsPaginated(string, int, int) ([]*Item, error)
	GetItemsCount(string) (int, error)
}
