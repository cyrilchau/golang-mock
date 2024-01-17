package main

import (
	"context"
	"myapp/config"
	"myapp/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app := app.NewApp(context.Background(), cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
