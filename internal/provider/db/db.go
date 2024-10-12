package db

import (
	_ "github.com/lib/pq"

	"context"
	"log"

	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
)

func NewClient(env *config.Env) *ent.Client {
	client, err := ent.Open(
		"postgres",
		env.DBConnection,
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
