package handler

import (
	"net/http"

	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/dto"
	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/middleware"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/usecase"
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/jwtutil"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserPetHandler struct {
	listUserPetsUseCase  *usecase.ListUserPetsUseCase
	addUserPetUseCase    *usecase.AddUserPetUseCase
	removeUserPetUseCase *usecase.RemoveUserPetUseCase
}

func NewUserPetHandler(
	listUserPetsUseCase *usecase.ListUserPetsUseCase,
	addUserPetUseCase *usecase.AddUserPetUseCase,
	removeUserPetUseCase *usecase.RemoveUserPetUseCase,
) *UserPetHandler {
	return &UserPetHandler{
		listUserPetsUseCase:  listUserPetsUseCase,
		addUserPetUseCase:    addUserPetUseCase,
		removeUserPetUseCase: removeUserPetUseCase,
	}
}

func (h *UserPetHandler) ListUserPets(
	c echo.Context,
) error {
	claims, ok := c.Get(middleware.ClaimsKey).(*jwtutil.UserClaims)
	if !ok {
		return errs.New("invalid claims")
	}

	userID, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return errs.New("invalid user id")
	}

	pets, err := h.listUserPetsUseCase.Execute(
		c.Request().Context(),
		usecase.ListUserPetsUseCaseInput{
			UserID: userID,
		},
	)
	if err != nil {
		return errs.New(err)
	}

	var data []dto.Pet
	if err := copier.Copy(&data, pets); err != nil {
		return errs.New(err)
	}

	return c.JSON(http.StatusOK, dto.ListUserPetsResponseDTO{
		Data: data,
	})
}

func (h *UserPetHandler) AddUserPet(
	c echo.Context,
) error {
	petIDString := c.Param("pet_id")

	petID, err := uuid.Parse(petIDString)
	if err != nil {
		return errs.New("invalid pet id")
	}

	claims, ok := c.Get(middleware.ClaimsKey).(*jwtutil.UserClaims)
	if !ok {
		return errs.New("invalid claims")
	}

	userID, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return errs.New("invalid user id")
	}

	if err = h.addUserPetUseCase.Execute(
		c.Request().Context(),
		usecase.AddUserPetUseCaseInput{
			UserID: userID,
			PetID:  petID,
		},
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *UserPetHandler) RemoveUserPet(
	c echo.Context,
) error {
	petIDString := c.Param("pet_id")

	petID, err := uuid.Parse(petIDString)
	if err != nil {
		return errs.New("invalid pet id")
	}

	claims, ok := c.Get(middleware.ClaimsKey).(*jwtutil.UserClaims)
	if !ok {
		return errs.New("invalid claims")
	}

	userID, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return errs.New("invalid user id")
	}

	if err = h.removeUserPetUseCase.Execute(
		c.Request().Context(),
		usecase.RemoveUserPetUseCaseInput{
			UserID: userID,
			PetID:  petID,
		},
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
