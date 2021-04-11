package config

const DefaultCustomersPageSize = 5

const DefaultItemsPageSize = 5

var ApiConfig Config = Config{CustomersPageSize: DefaultCustomersPageSize, ItemsPageSize: DefaultItemsPageSize}

// Config is a container of api configuration data
type Config struct {
	Port              string `yaml:"port"`
	PGConn            string `yaml:"pg_conn"`
	StaticDir         string `yaml:"static_dir"`
	AuthSecret        string `yaml:"authsecret"`
	FBServiceFile     string `yaml:"service_file_location"`
	CustomersPageSize int    `yaml:"customers_page_size"`
	ItemsPageSize     int    `yaml:"items_page_size"`
}
