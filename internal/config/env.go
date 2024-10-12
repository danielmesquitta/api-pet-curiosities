package config

import (
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/validator"
	"github.com/spf13/viper"
)

type Environment string

const (
	DevelopmentEnv Environment = "development"
	ProductionEnv  Environment = "production"
)

type Env struct {
	val validator.Validator

	Environment  Environment `mapstructure:"ENVIRONMENT"`
	Port         string      `mapstructure:"PORT"`
	DBConnection string      `mapstructure:"DB_CONNECTION"  validate:"required"`
	JWTSecretKey string      `mapstructure:"JWT_SECRET_KEY" validate:"required"`
	OpenAIToken  string      `mapstructure:"OPEN_AI_TOKEN"  validate:"required"`
}

func (e *Env) validate() error {
	if err := e.val.Validate(e); err != nil {
		return err
	}
	if e.Environment == "" {
		e.Environment = DevelopmentEnv
	}
	if e.Port == "" {
		e.Port = "8080"
	}
	return nil
}

func LoadEnv(val validator.Validator) *Env {
	env := &Env{
		val: val,
	}

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	if err := env.validate(); err != nil {
		panic(err)
	}

	return env
}
