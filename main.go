package main

import (
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

	app.Run(os.Args)
}
