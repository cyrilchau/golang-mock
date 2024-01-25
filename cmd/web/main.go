package main

import (
	"context"
	"myapp/config"
	webApp "myapp/internal/app/web"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cfg, err := config.LoadConfig("web")
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) {
				webApp := webApp.NewApp(context.Background(), cfg)
				if err := webApp.Start(); err != nil {
					panic(err)
				}
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	app.Run(os.Args)
}
