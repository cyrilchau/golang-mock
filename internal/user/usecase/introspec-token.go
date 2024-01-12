package usecase

import (
	"context"
	"myapp/pkg/middleware"

	"github.com/golang-jwt/jwt/v5"
)

func (uc *usecase) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	claims, err := middleware.ParseJWT(accessToken)

	if err != nil {
		return nil, err
	}

	return claims, nil
}
