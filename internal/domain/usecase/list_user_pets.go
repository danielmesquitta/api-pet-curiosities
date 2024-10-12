package usecase

import (
	"context"

	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/google/uuid"
)

type ListUserPetsUseCase struct {
	dbClient *ent.Client
}

func NewListUserPetsUseCase(
	dbClient *ent.Client,
) *ListUserPetsUseCase {
	return &ListUserPetsUseCase{
		dbClient: dbClient,
	}
}

type ListUserPetsUseCaseInput struct {
	UserID uuid.UUID
}

func (u *ListUserPetsUseCase) Execute(
	ctx context.Context,
	in ListUserPetsUseCaseInput,
) (ent.Pets, error) {
	pets, err := u.dbClient.Pet.Query().Where(
		pet.HasOwnersWith(
			user.ID(in.UserID),
		),
	).All(ctx)

	if err != nil {
		return nil, err
	}

	return pets, nil
}
