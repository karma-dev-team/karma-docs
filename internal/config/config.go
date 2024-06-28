package config

import "github.com/kelseyhightower/envconfig"

type AppConfig struct {
	Database struct {
		Name     string `envconfig:"POSTGRES_DB"`
		User     string `envconfig:"POSTGRES_USER"`
		Password string `envconfig:"POSTGRES_PASSWORD"`
		Host     string `envconfig:"POSTGRES_HOST"`
		Port     int    `envconfig:"POSTGRES_PORT"`
	}
	Openfga struct {
		DatabaseDsn string `envconfig:"OPENFGA_DATABASE"`
		ApiUrl      string `envconfig:"OPENFGA_APIURL"`
		StoreId     string `envconfig:"OPENFGA_STOREID"`
	}
	Debug  bool `envconfig:"DEBUG"`
	Logger struct {
		Level string `envconfig:"LOGGING_LEVEL"`
	}
	Jwt struct {
		TokenKey       string `envconfig:"TOKEN_SIGNING_KEY"`
		ExpireDuration int64  `envconfig:"TOKEN_EXPIRE_DURATION"`
	}
}

func NewAppConfig(path string) (*AppConfig, error) {
	cfg := new(AppConfig)
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
