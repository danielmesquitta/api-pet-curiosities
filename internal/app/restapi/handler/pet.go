package handler

import (
	"net/http"

	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/dto"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/usecase"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type PetHandler struct {
	listPetsUseCase *usecase.ListPetsUseCase
}

func NewPetHandler(
	listPetsUseCase *usecase.ListPetsUseCase,
) *PetHandler {
	return &PetHandler{
		listPetsUseCase: listPetsUseCase,
	}
}

func (h *PetHandler) ListPets(
	c echo.Context,
) error {
	specieString := c.QueryParam("specie")
	var specie pet.Specie
	if specieString != "" {
		specie = pet.Specie(specieString)
	}

	breed := c.QueryParam("breed")

	pets, err := h.listPetsUseCase.Execute(
		c.Request().Context(),
		usecase.ListPetsUseCaseInput{
			Specie: specie,
			Breed:  breed,
		},
	)
	if err != nil {
		return err
	}

	var data []dto.Pet
	if err := copier.Copy(&data, pets); err != nil {
		return errs.New(err)
	}

	return c.JSON(http.StatusOK, dto.ListPetsResponseDTO{
		Data: data,
	})
}
