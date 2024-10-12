package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/view"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/gpt"
)

type CreateDailyCuriosities struct {
	dbClient    *ent.Client
	gptProvider gpt.Provider
}

func NewCreateDailyCuriosities(
	dbClient *ent.Client,
	gptProvider gpt.Provider,
) *CreateDailyCuriosities {
	return &CreateDailyCuriosities{
		dbClient:    dbClient,
		gptProvider: gptProvider,
	}
}

type Curiosity struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	PetID   uuid.UUID `json:"pet_id"`
}

/*
- Execute creates daily curiosities for pets with owners.
- New curiosities should be created only if:
- - The pet has owners
- - The latest curiosity for the pet was viewed OR the pet has no curiosities
*/
func (u *CreateDailyCuriosities) Execute(
	ctx context.Context,
) error {
	petsHavingOwners, curiositiesWithViews, err := u.fetchFromDatabase(
		ctx,
	)
	if err != nil {
		return err
	}

	if len(petsHavingOwners) == 0 {
		return nil
	}

	latestCuriositiesByPets := u.groupLatestCuriositiesByPets(
		curiositiesWithViews,
	)

	curiosities, err := u.generateNewCuriosities(
		petsHavingOwners,
		latestCuriositiesByPets,
	)
	if err != nil {
		return errs.New(err)
	}

	if len(curiosities) == 0 {
		return nil
	}

	if err := u.saveCuriosities(ctx, curiosities); err != nil {
		return errs.New(err)
	}

	return nil
}

func (u *CreateDailyCuriosities) fetchFromDatabase(
	ctx context.Context,
) ([]*ent.Pet, []*ent.Curiosity, error) {
	var petsHavingOwners []*ent.Pet
	var curiositiesWithViews []*ent.Curiosity

	var wg sync.WaitGroup
	errChan := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()

		var err error
		petsHavingOwners, err = u.dbClient.Pet.Query().Where(
			pet.HasOwners(),
		).All(ctx)
		errChan <- err
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		var err error
		curiositiesWithViews, err = u.dbClient.Curiosity.Query().WithViews(
			func(q *ent.ViewQuery) {
				q.Order(ent.Desc(view.FieldCreatedAt)).Limit(1)
			},
		).Order(ent.Desc(view.FieldCreatedAt)).All(ctx)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var joinedErrs error
	for err := range errChan {
		if err != nil {
			joinedErrs = errors.Join(joinedErrs, err)
		}
	}

	if joinedErrs != nil {
		return nil, nil, errs.New(joinedErrs)
	}

	return petsHavingOwners, curiositiesWithViews, nil
}

func (u *CreateDailyCuriosities) generateNewCuriosities(
	petsHavingOwners []*ent.Pet,
	latestCuriositiesByPets map[string][]*ent.Curiosity,
) ([]Curiosity, error) {
	var wg sync.WaitGroup
	curiosities := []Curiosity{}
	curiosityChan := make(chan Curiosity)
	errChan := make(chan error)

	wg.Add(len(petsHavingOwners))
	for _, pet := range petsHavingOwners {
		go func() {
			defer wg.Done()

			message := fmt.Sprintf(
				"Write a short title and interesting fact about the %s breed (max 3 lines), "+
					"and return in JSON format with the keys \"title\" and \"content\"",
				pet.Breed,
			)

			var topics string
			if latestPetCuriosities, ok := latestCuriositiesByPets[pet.ID.String()]; ok {
				latestPetCuriosity := latestPetCuriosities[0]

				if latestPetCuriosityWasNotViewed := latestPetCuriosity != nil &&
					len(latestPetCuriosity.Edges.Views) == 0; latestPetCuriosityWasNotViewed {
					return
				}

				titles := make([]string, len(latestPetCuriosities))
				for _, curiosity := range latestPetCuriosities {
					title := fmt.Sprintf("\"%s\"", curiosity.Title)
					titles = append(titles, title)
				}

				topics = strings.Join(titles, ", ")
			}

			if topics != "" {
				message += fmt.Sprintf(
					". Avoid topics on %s",
					topics,
				)
			}

			completion, err := u.gptProvider.CreateChatCompletion(message)
			if err != nil {
				errChan <- err
				return
			}

			var curiosity Curiosity
			if err := json.Unmarshal([]byte(completion), &curiosity); err != nil {
				errChan <- err
				return
			}

			curiosity.PetID = pet.ID
			curiosityChan <- curiosity
		}()
	}

	go func() {
		wg.Wait()
		close(curiosityChan)
		close(errChan)
	}()

	var joinedErrs error
	for {
		select {
		case curiosity, ok := <-curiosityChan:
			if ok {
				curiosities = append(curiosities, curiosity)
			}
		case err, ok := <-errChan:
			if ok {
				joinedErrs = errors.Join(joinedErrs, err)
			}
		}
		if len(curiosityChan) == 0 && len(errChan) == 0 {
			break
		}
	}

	if joinedErrs != nil {
		return nil, errs.New(joinedErrs)
	}

	return curiosities, nil
}

func (u *CreateDailyCuriosities) groupLatestCuriositiesByPets(
	curiositiesWithViews []*ent.Curiosity,
) map[string][]*ent.Curiosity {
	latestCuriositiesByPets := make(map[string][]*ent.Curiosity)
	for _, curiosity := range curiositiesWithViews {
		petID := curiosity.Edges.Pet.ID.String()
		latestCuriositiesByPets[petID] = append(
			latestCuriositiesByPets[petID],
			curiosity,
		)
	}

	return latestCuriositiesByPets
}

func (u *CreateDailyCuriosities) saveCuriosities(
	ctx context.Context,
	curiosities []Curiosity,
) error {
	var builders []*ent.CuriosityCreate
	for _, curiosity := range curiosities {
		builders = append(
			builders,
			u.dbClient.Curiosity.Create().
				SetTitle(curiosity.Title).
				SetContent(curiosity.Content).
				SetPetID(curiosity.PetID),
		)
	}

	if _, err := u.dbClient.Curiosity.
		CreateBulk(builders...).Save(ctx); err != nil {
		return errs.New(err)
	}

	return nil
}
