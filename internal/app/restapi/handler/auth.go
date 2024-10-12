package handler

import (
	"net/http"

	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/dto"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	loginUseCase *usecase.LoginUseCase
}

func NewAuthHandler(
	loginUseCase *usecase.LoginUseCase,
) *AuthHandler {
	return &AuthHandler{
		loginUseCase: loginUseCase,
	}
}

func (h *AuthHandler) Login(
	c echo.Context,
) error {
	req := dto.LoginRequestDTO{}
	if err := c.Bind(&req); err != nil {
		return errs.New(err)
	}

	accessToken, refreshToken, err := h.loginUseCase.Execute(
		c.Request().Context(),
		usecase.LoginUseCaseInput{
			Name:  req.Name,
			Email: req.Email,
		},
	)
	if err != nil {
		return errs.New(err)
	}

	return c.JSON(http.StatusCreated, dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
