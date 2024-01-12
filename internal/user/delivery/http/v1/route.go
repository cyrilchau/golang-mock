package delivery

import (
	"myapp/config"
	"myapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserPrivateRoute(version *echo.Group, h Handlers, cfg config.Config) {
	user := version.Group("user")
	user.POST("/login", h.Login)
	user.POST("/create", h.Create)
	user.POST("/detail", h.Detail, middleware.AuthorizeJWT(cfg))
}
