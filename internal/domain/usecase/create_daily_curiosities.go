package usecase

import (
	"context"
	"errors"
	"sync"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/usercuriosity"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/gpt"
	"github.com/google/uuid"
)

type CreateDailyCuriositiesUseCase struct {
	dbClient             *ent.Client
	gptProvider          gpt.Provider
	makeCuriosityUseCase *MakeCuriosityUseCase
}

func NewCreateDailyCuriositiesUseCase(
	dbClient *ent.Client,
	gptProvider gpt.Provider,
	makeCuriosityUseCase *MakeCuriosityUseCase,
) *CreateDailyCuriositiesUseCase {
	return &CreateDailyCuriositiesUseCase{
		dbClient:             dbClient,
		gptProvider:          gptProvider,
		makeCuriosityUseCase: makeCuriosityUseCase,
	}
}

func (u *CreateDailyCuriositiesUseCase) Execute(
	ctx context.Context,
) error {
	petWithOwnersAndUserCuriosities, curiosities, err := u.fetchData(ctx)
	if err != nil {
		return errs.New(err)
	}

	curiositiesByPet := map[uuid.UUID][]*ent.Curiosity{}
	for _, curiosity := range curiosities {
		curiositiesByPet[curiosity.Edges.Pet.ID] = append(
			curiositiesByPet[curiosity.Edges.Pet.ID],
			curiosity,
		)
	}

	newCuriosities := []*ent.Curiosity{}
	usersByCuriosities := map[uuid.UUID][]uuid.UUID{}
	for _, p := range petWithOwnersAndUserCuriosities {
		newCuriosityID := uuid.New()
		petCuriosities := curiositiesByPet[p.ID]

		for _, owner := range p.Edges.Owners {
			haveNotSeenLatestCuriosity := len(
				owner.Edges.UserCuriosities,
			) > 0 &&
				!owner.Edges.UserCuriosities[0].Viewed
			if haveNotSeenLatestCuriosity {
				continue
			}

			mustCreateNewCuriosity := len(
				petCuriosities,
			) == len(
				owner.Edges.UserCuriosities,
			)
			if mustCreateNewCuriosity {
				usersByCuriosities[newCuriosityID] = append(
					usersByCuriosities[newCuriosityID],
					owner.ID,
				)
				continue
			}

			userCuriosityIDs := map[uuid.UUID]struct{}{}
			for _, userCuriosity := range owner.Edges.UserCuriosities {
				userCuriosityIDs[userCuriosity.Edges.Curiosity.ID] = struct{}{}
			}

			for _, curiosity := range petCuriosities {
				if _, ok := userCuriosityIDs[curiosity.ID]; ok {
					continue
				}
				usersByCuriosities[curiosity.ID] = append(
					usersByCuriosities[curiosity.ID],
					owner.ID,
				)
				break
			}
		}

		if _, ok := usersByCuriosities[newCuriosityID]; ok {
			curiosityTitles := []string{}
			for _, curiosity := range petCuriosities {
				curiosityTitles = append(curiosityTitles, curiosity.Title)
			}

			curiosity, err := u.makeCuriosityUseCase.Execute(
				ctx,
				p.Breed,
				curiosityTitles,
			)
			if err != nil {
				return errs.New(err)
			}

			curiosity.Edges.Pet = p
			curiosity.ID = newCuriosityID
			newCuriosities = append(newCuriosities, curiosity)
		}
	}

	if err := u.saveCuriosities(ctx, newCuriosities); err != nil {
		return errs.New(err)
	}

	newUserCuriosities := []*ent.UserCuriosity{}
	for curiosityID, userIDs := range usersByCuriosities {
		for _, userID := range userIDs {
			newUserCuriosities = append(newUserCuriosities, &ent.UserCuriosity{
				Edges: ent.UserCuriosityEdges{
					User: &ent.User{
						ID: userID,
					},
					Curiosity: &ent.Curiosity{
						ID: curiosityID,
					},
				},
			})
		}
	}

	if err := u.saveUserCuriosities(ctx, newUserCuriosities); err != nil {
		return errs.New(err)
	}

	return nil
}

func (u *CreateDailyCuriositiesUseCase) fetchData(
	ctx context.Context,
) ([]*ent.Pet, []*ent.Curiosity, error) {
	type Result struct {
		Pets        []*ent.Pet
		Curiosities []*ent.Curiosity
		Error       error
	}

	jobsCount := 2
	wg := sync.WaitGroup{}
	wg.Add(jobsCount)
	resultCh := make(chan Result, jobsCount)

	go func() {
		defer wg.Done()
		petWithOwnersAndUserCuriosities, err := u.dbClient.Pet.Query().
			Where(pet.HasOwners()).
			WithOwners(func(q *ent.UserQuery) {
				q.WithUserCuriosities(func(q *ent.UserCuriosityQuery) {
					q.Order(ent.Desc(usercuriosity.FieldCreatedAt))
				})
			}).
			All(ctx)
		resultCh <- Result{
			Pets:  petWithOwnersAndUserCuriosities,
			Error: errs.New(err),
		}
	}()

	go func() {
		defer wg.Done()
		curiosities, err := u.dbClient.Curiosity.Query().All(ctx)
		resultCh <- Result{
			Curiosities: curiosities,
			Error:       errs.New(err),
		}
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var pets []*ent.Pet
	var curiosities []*ent.Curiosity
	var err error
	for res := range resultCh {
		if res.Error != nil {
			err = errors.Join(err, res.Error)
		}
		if len(res.Pets) > 0 {
			pets = res.Pets
		}
		if len(res.Curiosities) > 0 {
			curiosities = res.Curiosities
		}
	}

	if err != nil {
		return nil, nil, errs.New(err)
	}

	return pets, curiosities, nil
}

func (u *CreateDailyCuriositiesUseCase) saveCuriosities(
	ctx context.Context,
	curiosities []*ent.Curiosity,
) error {
	var builders []*ent.CuriosityCreate
	for _, curiosity := range curiosities {
		builders = append(
			builders,
			u.dbClient.Curiosity.Create().
				SetID(curiosity.ID).
				SetTitle(curiosity.Title).
				SetContent(curiosity.Content).
				SetPetID(curiosity.Edges.Pet.ID),
		)
	}

	if _, err := u.dbClient.Curiosity.
		CreateBulk(builders...).Save(ctx); err != nil {
		return errs.New(err)
	}

	return nil
}

func (u *CreateDailyCuriositiesUseCase) saveUserCuriosities(
	ctx context.Context,
	userCuriosities []*ent.UserCuriosity,
) error {
	var builders []*ent.UserCuriosityCreate
	for _, userCuriosity := range userCuriosities {
		builders = append(
			builders,
			u.dbClient.UserCuriosity.Create().
				SetCuriosityID(userCuriosity.Edges.Curiosity.ID).
				SetUserID(userCuriosity.Edges.User.ID),
		)
	}
	_, err := u.dbClient.UserCuriosity.
		CreateBulk(builders...).Save(ctx)
	if err != nil {
		return errs.New(err)
	}

	return nil
}
