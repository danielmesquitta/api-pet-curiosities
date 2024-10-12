package main

import (
	"context"
	"sync"
	"time"

	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/validator"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/seed"
)

func main() {
	val := validator.NewValidate()
	env := config.LoadEnv(val)
	dbClient := db.NewClient(env)

	defer dbClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	jobsCount := 2

	wg.Add(jobsCount)

	go func() {
		defer wg.Done()
		if err := seed.CreateDogs(ctx, dbClient); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := seed.CreateCats(ctx, dbClient); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
