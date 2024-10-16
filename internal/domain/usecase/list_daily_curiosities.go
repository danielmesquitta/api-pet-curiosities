package usecase

import (
	"context"
	"errors"
	"sync"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/usercuriosity"
	"github.com/google/uuid"
)

type ListDailyCuriosities struct {
	dbClient *ent.Client
}

func NewListDailyCuriosities(
	dbClient *ent.Client,
) *ListDailyCuriosities {
	return &ListDailyCuriosities{
		dbClient: dbClient,
	}
}

type ListDailyCuriositiesInput struct {
	UserID uuid.UUID
}

func (u *ListDailyCuriosities) Execute(
	ctx context.Context,
	in ListDailyCuriositiesInput,
) ([]*ent.Curiosity, error) {
	pets, err := u.dbClient.Pet.Query().
		Where(pet.HasOwnersWith(user.ID(in.UserID))).
		All(ctx)
	if err != nil {
		return nil, errs.New(err)
	}

	type Result struct {
		Curiosity *ent.Curiosity
		Error     error
	}

	jobsCount := len(pets)
	resultCh := make(chan Result, jobsCount)
	wg := sync.WaitGroup{}
	wg.Add(jobsCount)

	for _, p := range pets {
		go func() {
			defer wg.Done()
			uc, err := u.dbClient.UserCuriosity.Query().
				Where(usercuriosity.HasUserWith(
					user.ID(in.UserID),
					user.HasPetsWith(pet.ID(p.ID)),
				)).
				WithCuriosity().
				Order(ent.Desc(usercuriosity.FieldCreatedAt)).
				First(ctx)
			if err != nil {
				resultCh <- Result{Error: errs.New(err)}
				return
			}
			resultCh <- Result{Curiosity: uc.Edges.Curiosity}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	curiosities := []*ent.Curiosity{}
	for res := range resultCh {
		if res.Error != nil {
			err = errors.Join(err, res.Error)
		}
		if res.Curiosity != nil {
			curiosities = append(curiosities, res.Curiosity)
		}
	}

	if err != nil {
		return nil, errs.New(err)
	}

	return curiosities, nil
}
