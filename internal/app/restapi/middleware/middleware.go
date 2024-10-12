package middleware

import "github.com/danielmesquitta/api-pet-curiosities/internal/pkg/jwtutil"

type Middleware struct {
	jwtManager jwtutil.JWTManager
}

func NewMiddleware(
	jwtManager jwtutil.JWTManager,
) *Middleware {
	return &Middleware{
		jwtManager: jwtManager,
	}
}
