package usecase

import (
	"context"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/google/uuid"
)

type AddUserPetUseCase struct {
	dbClient *ent.Client
}

func NewAddUserPetUseCase(
	dbClient *ent.Client,
) *AddUserPetUseCase {
	return &AddUserPetUseCase{
		dbClient: dbClient,
	}
}

type AddUserPetUseCaseInput struct {
	UserID uuid.UUID
	PetID  uuid.UUID
}

func (u *AddUserPetUseCase) Execute(
	ctx context.Context,
	in AddUserPetUseCaseInput,
) error {
	if _, err := u.dbClient.User.
		Update().
		Where(user.ID(in.UserID)).
		AddPetIDs(in.PetID).
		Save(ctx); err != nil {
		return errs.New(err)
	}

	return nil
}
