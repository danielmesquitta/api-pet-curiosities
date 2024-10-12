package jwtutil

import (
	"time"

	"github.com/danielmesquitta/api-pet-curiosities/internal/config"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/user"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager interface {
	NewAccessToken(claims UserClaims) (accessToken string, err error)
	NewRefreshToken(
		claims jwt.RegisteredClaims,
	) (refreshToken string, err error)
	ValidateAccessToken(accessToken string) (*UserClaims, error)
	ValidateRefreshToken(refreshToken string) (*jwt.RegisteredClaims, error)
}

type UserClaims struct {
	Tier user.Tier `json:"tier,omitempty"`
	jwt.RegisteredClaims
}

type JWT struct {
	secretKey []byte
}

func NewJWT(
	env *config.Env,
) *JWT {
	return &JWT{
		secretKey: []byte(env.JWTSecretKey),
	}
}

func (j *JWT) NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString(j.secretKey)
}

func (j *JWT) NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString(j.secretKey)
}

func (j *JWT) ValidateAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(_ *jwt.Token) (interface{}, error) {
			return j.secretKey, nil
		},
	)
	if err != nil {
		return nil, errs.New(err)
	}

	userClaims, ok := parsedAccessToken.Claims.(*UserClaims)
	if !ok {
		return nil, errs.New("invalid claims")
	}

	if j.isExpired(&userClaims.RegisteredClaims) {
		return nil, errs.New("token is expired")
	}

	return userClaims, nil
}

func (j *JWT) ValidateRefreshToken(
	refreshToken string,
) (*jwt.RegisteredClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(
		refreshToken,
		&jwt.RegisteredClaims{},
		func(_ *jwt.Token) (interface{}, error) {
			return j.secretKey, nil
		},
	)
	if err != nil {
		return nil, errs.New(err)
	}

	claims, ok := parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errs.New("invalid claims")
	}

	if j.isExpired(claims) {
		return nil, errs.New("token is expired")
	}

	return claims, nil
}

func (j *JWT) isExpired(claims *jwt.RegisteredClaims) bool {
	return claims.ExpiresAt.Before(time.Now())
}
