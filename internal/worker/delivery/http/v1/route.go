package delivery

import (
	"myapp/config"
	"myapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserPrivateRoute(version *echo.Group, s Services, cfg config.Config) {
	worker := version.Group("worker")
	worker.POST("/create", s.Create, middleware.AuthorizeJWT(cfg))
	worker.POST("/detail", s.Detail, middleware.AuthorizeJWT(cfg))
}
