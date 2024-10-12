package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/pkg/jwtutil"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	entUser "github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/golang-jwt/jwt/v5"
)

type LoginUseCase struct {
	dbClient   *ent.Client
	jwtManager jwtutil.JWTManager
}

func NewLoginUseCase(
	dbClient *ent.Client,
	jwtManager jwtutil.JWTManager,
) *LoginUseCase {
	return &LoginUseCase{
		dbClient:   dbClient,
		jwtManager: jwtManager,
	}
}

type LoginUseCaseInput struct {
	Name  string
	Email string
}

func (u *LoginUseCase) Execute(
	ctx context.Context,
	in LoginUseCaseInput,
) (accessToken, refreshToken string, err error) {
	in.Email = strings.ToLower(in.Email)

	user, err := u.dbClient.User.Query().Where(
		entUser.Email(in.Email),
	).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", "", errs.New(err)
	}

	if user == nil {
		user, err = u.dbClient.User.Create().
			SetTier(entUser.TierFREE).
			SetName(in.Name).
			SetEmail(in.Email).
			Save(ctx)
		if err != nil {
			return "", "", errs.New(err)
		}
	}

	accessToken, err = u.jwtManager.NewAccessToken(jwtutil.UserClaims{
		Tier: user.Tier,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	})
	if err != nil {
		return "", "", errs.New(err)
	}

	refreshToken, err = u.jwtManager.NewRefreshToken(jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
	})
	if err != nil {
		return "", "", errs.New(err)
	}

	return accessToken, refreshToken, nil
}
