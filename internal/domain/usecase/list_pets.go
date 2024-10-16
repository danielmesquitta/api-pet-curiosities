package usecase

import (
	"context"

	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/fmtutil"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/predicate"
)

type ListPetsUseCase struct {
	dbClient *ent.Client
}

func NewListPetsUseCase(
	dbClient *ent.Client,
) *ListPetsUseCase {
	return &ListPetsUseCase{
		dbClient: dbClient,
	}
}

type ListPetsUseCaseInput struct {
	Specie pet.Specie
	Breed  string
}

func (u *ListPetsUseCase) Execute(
	ctx context.Context,
	in ListPetsUseCaseInput,
) (ent.Pets, error) {
	in.Breed = fmtutil.ToSearchable(in.Breed)

	where := []predicate.Pet{}
	if in.Specie != "" {
		where = append(where, pet.SpecieEQ(in.Specie))
	}

	if in.Breed != "" {
		where = append(where, pet.BreedContainsFold(in.Breed))
	}

	pets, err := u.dbClient.Pet.Query().Where(where...).Order(
		pet.ByBreed(),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	return pets, nil
}
