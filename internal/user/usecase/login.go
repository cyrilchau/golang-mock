package usecase

import (
	"context"
	"myapp/internal/user/dtos"
	errorCode "myapp/pkg/errorCode"
	"myapp/pkg/middleware"
	"myapp/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (uc *usecase) Login(ctx context.Context, request dtos.UserLoginRequest) (response dtos.UserLoginResponse, httpCode int, err error) {
	dataLogin, err := uc.repo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	if !strings.EqualFold(utils.Decrypt(dataLogin.Password, uc.cfg), request.Password) {
		return response, http.StatusUnauthorized, errorCode.ErrInvalidPassword
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	identityData := middleware.CustomClaims{
		UserID: dataLogin.UserID,
		Email:  dataLogin.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token, err := middleware.GenerateJWT(identityData, uc.cfg.Authentication.Key)
	if err != nil {
		return response, http.StatusInternalServerError, errorCode.ErrFailedGenerateJWT
	}

	response = dtos.UserLoginResponse{
		AccessToken: token,
		ExpiredAt:   utils.UnixToDuration(identityData.ExpiresAt.Unix()),
	}

	return response, http.StatusOK, nil
}
