package config

const DefaultCustomersPageSize = 5

// AppConfig is a container of api configuration data
type AppConfig struct {
	Port              string `yaml:"port"`
	PGConn            string `yaml:"pg_conn"`
	StaticDir         string `yaml:"static_dir"`
	AuthSecret        string `yaml:"authsecret"`
	FBServiceFile     string `yaml:"service_file_location"`
	CustomersPageSize int    `yaml:"customers_page_size"`
}
