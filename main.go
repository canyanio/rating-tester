package main

import (
	"os"

	"github.com/mendersoftware/go-lib-micro/log"
	"github.com/urfave/cli"
)

func main() {
	doMain(os.Args)
}

func doMain(args []string) {
	var configDebug bool

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "Enable debug mode and verbose logging",
				Destination: &configDebug,
			},
		},
		Commands: []cli.Command{
			{
				Name:   "load-test",
				Usage:  "Run the load testing suite",
				Action: cmdLoadTest,
				Flags:  []cli.Flag{},
			},
		},
	}
	app.Usage = "rating-tester"
	app.Version = "1.0.0"
	app.Action = cmdLoadTest

	app.Before = func(args *cli.Context) error {
		log.Setup(configDebug)

		return nil
	}

	err := app.Run(args)
	if err != nil {
		cli.NewExitError(err.Error(), 1)
	}
}

func cmdLoadTest(args *cli.Context) error {
	return nil
}
