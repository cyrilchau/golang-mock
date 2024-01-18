package delivery

import (
	"myapp/config"
	"myapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserPrivateRoute(version *echo.Group, s Services, cfg config.Config) {
	task := version.Group("task")
	task.POST("/create", s.Create, middleware.AuthorizeJWT(cfg))
	task.POST("/detail", s.Detail, middleware.AuthorizeJWT(cfg))
	task.POST("/all", s.List, middleware.AuthorizeJWT(cfg))
}
