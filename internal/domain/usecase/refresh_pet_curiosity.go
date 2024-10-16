package usecase

import (
	"context"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/usercuriosity"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/gpt"
	"github.com/google/uuid"
)

type RefreshPetCuriosity struct {
	dbClient             *ent.Client
	gptProvider          gpt.Provider
	makeCuriosityUseCase *MakeCuriosityUseCase
}

func NewRefreshPetCuriosity(
	dbClient *ent.Client,
	gptProvider gpt.Provider,
	makeCuriosityUseCase *MakeCuriosityUseCase,
) *RefreshPetCuriosity {
	return &RefreshPetCuriosity{
		dbClient:             dbClient,
		gptProvider:          gptProvider,
		makeCuriosityUseCase: makeCuriosityUseCase,
	}
}

type RefreshPetCuriosityUseCaseInput struct {
	UserID   uuid.UUID
	UserTier user.Tier
	PetID    uuid.UUID
}

func (u *RefreshPetCuriosity) Execute(
	ctx context.Context,
	in RefreshPetCuriosityUseCaseInput,
) (*ent.Curiosity, error) {
	if in.UserTier != user.TierPRO {
		return nil, errs.ErrUserNotAllowed
	}

	userPet, err := u.dbClient.Pet.Query().
		Where(pet.ID(in.PetID)).
		WithCuriosities(func(q *ent.CuriosityQuery) {
			q.WithUserCuriosities(func(q *ent.UserCuriosityQuery) {
				q.Where(usercuriosity.HasUserWith(user.ID(in.UserID))).
					Order(ent.Desc(usercuriosity.FieldCreatedAt))
			})
		}).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, errs.ErrPetNotFound
	}
	if err != nil {
		return nil, errs.New(err)
	}

	userCuriosities := []*ent.UserCuriosity{}

	for _, curiosity := range userPet.Edges.Curiosities {
		if len(curiosity.Edges.UserCuriosities) == 0 {
			continue
		}
		userCuriosities = append(
			userCuriosities,
			curiosity.Edges.UserCuriosities...,
		)
	}

	mustCreateNewCuriosity := len(
		userPet.Edges.Curiosities,
	) == len(
		userCuriosities,
	)
	if mustCreateNewCuriosity {
		curiosityTitles := []string{}
		for _, curiosity := range userPet.Edges.Curiosities {
			curiosityTitles = append(curiosityTitles, curiosity.Title)
		}

		curiosity, err := u.makeCuriosityUseCase.Execute(
			ctx,
			userPet.Breed,
			curiosityTitles,
		)
		if err != nil {
			return nil, errs.New(err)
		}

		curiosity.ID = uuid.New()
		curiosity.Edges.Pet = userPet

		if err = u.saveCuriosity(ctx, curiosity); err != nil {
			return nil, errs.New(err)
		}

		userCuriosity := &ent.UserCuriosity{
			Edges: ent.UserCuriosityEdges{
				User: &ent.User{
					ID: in.UserID,
				},
				Curiosity: curiosity,
			},
		}
		if err = u.saveUserCuriosity(ctx, userCuriosity); err != nil {
			return nil, errs.New(err)
		}

		return curiosity, nil
	}

	userCuriosityIDs := map[uuid.UUID]struct{}{}
	for _, uc := range userCuriosities {
		userCuriosityIDs[uc.Edges.Curiosity.ID] = struct{}{}
	}

	var userCuriosity *ent.UserCuriosity
	var curiosity *ent.Curiosity
	for _, c := range userPet.Edges.Curiosities {
		if _, ok := userCuriosityIDs[c.ID]; ok {
			continue
		}
		userCuriosity = &ent.UserCuriosity{
			Edges: ent.UserCuriosityEdges{
				User: &ent.User{
					ID: in.UserID,
				},
				Curiosity: c,
			},
		}
		curiosity = c
		break
	}

	if userCuriosity == nil {
		return nil, errs.New("no curiosity found")
	}

	if err = u.saveUserCuriosity(ctx, userCuriosity); err != nil {
		return nil, errs.New(err)
	}

	return curiosity, nil
}

func (u *RefreshPetCuriosity) saveCuriosity(
	ctx context.Context,
	curiosity *ent.Curiosity,
) error {
	_, err := u.dbClient.Curiosity.Create().
		SetContent(curiosity.Content).
		SetTitle(curiosity.Title).
		SetPetID(curiosity.Edges.Pet.ID).
		Save(ctx)
	if err != nil {
		return errs.New(err)
	}

	return nil
}

func (u *RefreshPetCuriosity) saveUserCuriosity(
	ctx context.Context,
	userCuriosity *ent.UserCuriosity,
) error {
	return u.dbClient.UserCuriosity.Create().
		SetCuriosityID(userCuriosity.Edges.Curiosity.ID).
		SetUserID(userCuriosity.Edges.User.ID).Exec(ctx)
}
