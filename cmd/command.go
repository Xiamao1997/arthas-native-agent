package main

import (
	"github.com/Xiamao1997/arthas-native-agent/pkg/agent/boot"
	"github.com/urfave/cli/v2"
)

// @author Xiamao1997 2022-06-20

var (
	cmdStart = cli.Command{
		Name:  "start",
		Usage: "Start Native Agent",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "ip",
				Aliases: []string{"i"},
				Usage:   "The tunnel server ip",
				EnvVars: []string{"ARTHAS_TUNNEL_SERVER_IP"},
				Value:   "127.0.0.1",
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "The tunnel server port",
				EnvVars: []string{"ARTHAS_TUNNEL_SERVER_PORT"},
				Value:   "7777",
			},
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "The native agent name",
			},
		},
		Action: func(c *cli.Context) error {
			ip := c.String("ip")
			port := c.String("port")
			name := c.String("name")
			return boot.Start(ip, port, name)
		},
	}
)
