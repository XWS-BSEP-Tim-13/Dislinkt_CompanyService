package config

import "os"

type Config struct {
	Port          string
	CompanyDBHost string
	CompanyDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("COMPANY_SERVICE_PORT"),
		CompanyDBHost: os.Getenv("COMPANY_DB_HOST"),
		CompanyDBPort: os.Getenv("COMPANY_DB_PORT"),
	}
}
