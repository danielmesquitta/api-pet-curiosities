package usecase

import (
	"context"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/google/uuid"
)

type RemoveUserPetUseCase struct {
	dbClient *ent.Client
}

func NewRemoveUserPetUseCase(
	dbClient *ent.Client,
) *RemoveUserPetUseCase {
	return &RemoveUserPetUseCase{
		dbClient: dbClient,
	}
}

type RemoveUserPetUseCaseInput struct {
	UserID uuid.UUID
	PetID  uuid.UUID
}

func (u *RemoveUserPetUseCase) Execute(
	ctx context.Context,
	in RemoveUserPetUseCaseInput,
) error {
	if _, err := u.dbClient.User.
		Update().
		Where(user.ID(in.UserID)).
		RemovePetIDs(in.PetID).
		Save(ctx); err != nil {
		return errs.New(err)
	}

	return nil
}
