package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Database struct {
		Name     string `envconfig:"POSTGRES_DB"`
		User     string `envconfig:"POSTGRES_USER"`
		Password string `envconfig:"POSTGRES_PASSWORD"`
		Host     string `envconfig:"POSTGRES_HOST"`
		Port     int    `envconfig:"POSTGRES_PORT"`
	}
	Openfga struct {
		DatabaseDsn          string `envconfig:"OPENFGA_DATABASE"`
		ApiUrl               string `envconfig:"OPENFGA_APIURL"`
		StoreId              string `envconfig:"OPENFGA_STOREID"`
		FilePath             string `envconfig:OPENFGA_FILEPATH`
		AuthorizationModelId string // writes in runtime, bc why not???
	}
	Debug  bool   `envconfig:"DEBUG"`
	Port   string `envconfig:"PORT"`
	Logger struct {
		Level string `envconfig:"LOGGING_LEVEL"`
		Path  string `envconfig:"LOGGING_PATH"`
	}
	Jwt struct {
		TokenKey       string `envconfig:"TOKEN_SIGNING_KEY"`
		ExpireDuration int64  `envconfig:"TOKEN_EXPIRE_DURATION"`
	}
}

func NewAppConfig() (*AppConfig, error) {
	cfg := new(AppConfig)
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *AppConfig) GenerateDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
}
