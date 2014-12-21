package main

import (
	"os"

	"bitbucket.org/pferdefleisch/dbpm/commands"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dbpm"
	app.Usage = "pick machine server and tasks"

	app.Commands = []cli.Command{
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "search for a term",
			Action: func(c *cli.Context) {
				term := c.Args().First()
				commands.Search(term)
			},
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "update pick database from api",
			Action: func(c *cli.Context) {
				commands.Update()
			},
		},
	}

	app.Run(os.Args)
}
