package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/pferdefleisch/dbpm/commands"
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
	}

	app.Run(os.Args)
}
