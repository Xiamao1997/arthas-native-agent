package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// @author Xiamao1997 2022-06-20

func main() {
	app := &cli.App{
		Name:        "arthas-native-boot",
		Version:     "0.1.0",
		Usage:       "Native Agent is for cluster management.",
		Description: "Native Agent is for cluster management.",
		Compiled:    time.Now(),
		Action:      cli.ShowAppHelp,
		Commands: []*cli.Command{
			&cmdStart,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("Failed to run Native Agent App:", err)
	}
}
