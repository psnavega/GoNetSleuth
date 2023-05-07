package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI application"
	app.Usage = "Search ips and server names"

	return app
}
