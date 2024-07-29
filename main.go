package main

import (
	"log"
	"os"

	"github.com/moecasts/with-env/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "with-env",
		Usage: "Run command with environment variables from specific .env file",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "path of the .env file",
				Value:   cli.NewStringSlice("~/.env", "./.env"),
			},
		},
		Action: func(ctx *cli.Context) error {
			return internal.WithEnvAction(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
