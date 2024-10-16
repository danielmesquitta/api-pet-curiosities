package main

import (
	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/validator"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db"
)

func main() {
	val := validator.NewValidate()
	env := config.LoadEnv(val)
	_ = db.NewClient(env)
}
