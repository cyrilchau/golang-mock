package delivery

import (
	"myapp/config"
	"myapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserPrivateRoute(version *echo.Group, h Handlers, cfg config.Config) {
	task := version.Group("task")
	task.POST("/create", h.Create, middleware.AuthorizeJWT(cfg))
	task.POST("/detail", h.Detail, middleware.AuthorizeJWT(cfg))
	task.POST("/all", h.List, middleware.AuthorizeJWT(cfg))
}
