package main

import (
	"context"
	"myapp/config"
	authApp "myapp/internal/app/auth"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cfg, err := config.LoadConfig("auth")
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Action: func(c *cli.Context) {
				authApp := authApp.NewApp(context.Background(), cfg)
				if err := authApp.Start(); err != nil {
					panic(err)
				}
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "a")
	}

	app.Run(os.Args)
}
