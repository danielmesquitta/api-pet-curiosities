package usecase

import (
	"context"

	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/google/uuid"
)

type GetUserProfileUseCase struct {
	dbClient *ent.Client
}

func NewGetUserProfileUseCase(
	dbClient *ent.Client,
) *GetUserProfileUseCase {
	return &GetUserProfileUseCase{dbClient: dbClient}
}

type GetUserProfileUseCaseInput struct {
	UserID uuid.UUID
}

func (u *GetUserProfileUseCase) Execute(
	ctx context.Context,
	in GetUserProfileUseCaseInput,
) (*ent.User, error) {
	user, err := u.dbClient.User.Query().Where(
		user.ID(in.UserID),
	).First(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
