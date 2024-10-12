package router

import (
	"github.com/labstack/echo/v4"

	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/handler"
	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/middleware"
	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
)

type Router struct {
	env            *config.Env
	mid            *middleware.Middleware
	authHandler    *handler.AuthHandler
	healthHandler  *handler.HealthHandler
	petHandler     *handler.PetHandler
	userPetHandler *handler.UserPetHandler
}

func NewRouter(
	env *config.Env,
	mid *middleware.Middleware,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
	petHandler *handler.PetHandler,
	userPetHandler *handler.UserPetHandler,
) *Router {
	return &Router{
		env:            env,
		mid:            mid,
		authHandler:    authHandler,
		healthHandler:  healthHandler,
		petHandler:     petHandler,
		userPetHandler: userPetHandler,
	}
}

func (r *Router) Register(
	app *echo.Echo,
) {
	basePath := "/api/v1"
	apiV1 := app.Group(basePath)

	// Unauthenticated routes
	apiV1.GET("/health", r.healthHandler.Health)
	apiV1.POST("/auth/login", r.authHandler.Login)

	// Authenticated routes
	apiV1.Use(r.mid.EnsureAuthenticated)

	apiV1.GET("/pets", r.petHandler.ListPets)

	apiV1.GET("/user/pets", r.userPetHandler.ListUserPets)
	apiV1.POST("/user/pets/:pet_id", r.userPetHandler.AddUserPet)
	apiV1.DELETE("/user/pets/:pet_id", r.userPetHandler.RemoveUserPet)
}
