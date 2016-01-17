package main

import (
	imp "geopardy/jImport"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	var log = logrus.New()
	app := cli.NewApp()
	app.Name = "Geopardy"
	app.Usage = "Import and query Jeopardy questions"
	app.Action = func(c *cli.Context) {
		log.Out = os.Stderr
		log.WithFields(logrus.Fields{
			"args": c.Args(),
		}).Info("Called with args")
	}

	app.Commands = []cli.Command{
		{
			Name:  "import",
			Usage: "Import Jeopardy questions from the web to the database",
			Action: func(c *cli.Context) {
				err := imp.RunImport(c.Args())
				if err != nil {
					log.Error(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
