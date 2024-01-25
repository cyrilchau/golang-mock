package auth

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/pkg/datasource"
	"myapp/pkg/middleware"
	"myapp/pkg/otel/zerolog"
	"myapp/pkg/validator"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type App struct {
	DB   *gorm.DB
	Echo *echo.Echo
	GRPC *grpc.Server
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
		GRPC: grpc.NewServer(),
		Log:  zerolog.NewZeroLog(ctx, os.Stdout),
		Cfg:  cfg,
	}
}

func (app *App) Start() error {
	if err := app.StartService(); err != nil {
		app.Log.Z().Err(err).Msg("[app]StartService")
		return err
	}

	go func () {
		listen, err := net.Listen("tcp", fmt.Sprintf(":%s", app.Cfg.Server.RPCPort))
		if err != nil {
			app.Log.Z().Err(err).Msg("[app]StartService")
	
		}
		if err := app.GRPC.Serve(listen); err != nil {
			app.Log.Z().Err(err).Msg("[app]StartService")
		}
	}()
	
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
