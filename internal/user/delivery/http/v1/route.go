package delivery

import (
	"myapp/config"
	"myapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserPrivateRoute(version *echo.Group, s Services, cfg config.Config) {
	user := version.Group("user")
	user.POST("/login", s.Login)
	user.POST("/create", s.Create)
	user.POST("/detail", s.Detail, middleware.AuthorizeJWT(cfg))
}
