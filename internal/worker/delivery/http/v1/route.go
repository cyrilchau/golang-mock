package delivery

import (
	"myapp/config"
	"myapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserPrivateRoute(version *echo.Group, h Handlers, cfg config.Config) {
	worker := version.Group("worker")
	worker.POST("/create", h.Create, middleware.AuthorizeJWT(cfg))
	worker.POST("/detail", h.Detail, middleware.AuthorizeJWT(cfg))
}
