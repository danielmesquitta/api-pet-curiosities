package restapi

import (
	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/handler"
	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/middleware"
	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/router"
	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/usecase"
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/jwtutil"
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/validator"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func Start() {
	depsProvider := fx.Provide(
		// Config
		config.LoadEnv,

		// PKGs
		fx.Annotate(
			validator.NewValidate,
			fx.As(new(validator.Validator)),
		),
		fx.Annotate(
			jwtutil.NewJWT,
			fx.As(new(jwtutil.JWTManager)),
		),

		// Providers
		db.NewClient,

		// Use cases
		usecase.NewLoginUseCase,
		usecase.NewListPetsUseCase,
		usecase.NewAddUserPetUseCase,
		usecase.NewListUserPetsUseCase,
		usecase.NewRemoveUserPetUseCase,

		// Handlers
		handler.NewAuthHandler,
		handler.NewHealthHandler,
		handler.NewPetHandler,
		handler.NewUserPetHandler,

		// Middleware
		middleware.NewMiddleware,

		// Router
		router.NewRouter,

		// App
		NewServer,
	)

	container := fx.New(
		depsProvider,
		fx.Invoke(func(*echo.Echo) {}),
	)

	container.Run()
}
