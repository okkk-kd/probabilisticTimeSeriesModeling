package config

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Postgres Postgres
	Logger   Logger
}

type Server struct {
	IP                          string `validate:"required"`
	Port                        string `validate:"required"`
	APIKey                      string `validate:"required"`
	NasdaqURL                   string `validate:"required"`
	ShowUnknownErrorsInResponse bool   `validate:"required"`
}

type Postgres struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SSLMode  string `validate:"required"`
	PGDriver string `validate:"required"`
	Settings struct {
		MaxOpenConns    int           `validate:"required,min=1"`
		ConnMaxLifetime time.Duration `validate:"required,min=1"`
		MaxIdleConns    int           `validate:"required,min=1"`
		ConnMaxIdleTime time.Duration `validate:"required,min=1"`
	}
}

type Logger struct {
	Level zerolog.Level `validate:"required"`
	File  string
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	return
}
