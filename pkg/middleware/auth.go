package middleware

import (
	"errors"
	"fmt"
	"myapp/config"
	"myapp/pkg/utils/response"

	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// CustomClaims represents the custom claims you want to include in the JWT payload.
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(data CustomClaims, secretKey string) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	// Extract the expiration time
	claims, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok {
		err := fmt.Errorf("Can't convert token's claims to standard claims")
		return nil, err
	}
	return &claims, nil
}

func AuthorizeJWT(cfg config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractBearerToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, response.NewResponseError(http.StatusUnauthorized, response.MsgFailed, err.Error()))
			}

			token, err := jwt.ParseWithClaims(*auth, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.Authentication.Key), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, response.NewResponseError(http.StatusUnauthorized, response.MsgFailed, err.Error()))
			}

			if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
				c.Set("identity", claims)

				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, response.NewResponseError(http.StatusUnauthorized, response.MsgFailed, err.Error()))
		}
	}
}

func extractBearerToken(c echo.Context) (*string, error) {
	authData := c.Request().Header.Get("Authorization")
	if authData == "" {
		return nil, errors.New("authorization can't be nil")
	}
	parts := strings.Split(authData, " ")
	if len(parts) < 2 {
		return nil, errors.New("invalid authorization value")
	}
	if parts[0] != "Bearer" {
		return nil, errors.New("auth should be bearer")
	}

	return &parts[1], nil
}
