package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

// @author Xiamao1997 2022-06-20

func main() {
	app := &cli.App{
		Name:        "arthas-native-agent",
		Version:     "0.1.0",
		Usage:       "Native agent is for cluster management.",
		Description: "Native agent is for cluster management.",
		Compiled:    time.Now(),
		Action:      cli.ShowAppHelp,
		Commands: []*cli.Command{
			&cmdStart,
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalln("Failed to run native agent app:", err)
	}
}
