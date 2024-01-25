package web

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/pkg/datasource"
	"myapp/pkg/middleware"
	"myapp/pkg/otel/zerolog"
	"myapp/pkg/validator"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	DB   *gorm.DB
	Echo *echo.Echo
	Log  *zerolog.Logger
	Cfg  config.Config
}

func NewApp(ctx context.Context, cfg config.Config) *App {
	db, err := datasource.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	return &App{
		DB:   db,
		Echo: echo.New(),
		Log:  zerolog.NewZeroLog(ctx, os.Stdout),
		Cfg:  cfg,
	}
}

func (app *App) Start() error {
	if err := app.StartService(); err != nil {
		app.Log.Z().Err(err).Msg("[app]StartService")
		return err
	}

	app.Echo.Debug = app.Cfg.Server.Debug
	app.Echo.Validator = validator.NewValidator()

	app.Echo.Use(middleware.AppCors())
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.CacheWithRevalidation)

	return app.Echo.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%s", app.Cfg.Server.RESTPort),
		ReadTimeout:  app.Cfg.Server.ReadTimeout,
		WriteTimeout: app.Cfg.Server.WriteTimeout,
	})
}
