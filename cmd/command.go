package main

import (
	"github.com/Xiamao1997/arthas-native-agent/pkg/agent/boot"
	"github.com/urfave/cli/v2"
)

// @author Xiamao1997 2022-06-20

var (
	cmdStart = cli.Command{
		Name:  "start",
		Usage: "Start native agent",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "tunnel-server-ip",
				Aliases: []string{"i"},
				Usage:   "The tunnel server ip",
				EnvVars: []string{"ARTHAS_TUNNEL_SERVER_IP"},
				Value:   "127.0.0.1",
			},
			&cli.StringFlag{
				Name:    "tunnel-server-port",
				Aliases: []string{"p"},
				Usage:   "The tunnel server port",
				EnvVars: []string{"ARTHAS_TUNNEL_SERVER_PORT"},
				Value:   "7777",
			},
			&cli.StringFlag{
				Name:    "native-agent-name",
				Aliases: []string{"n"},
				Usage:   "The native agent name",
				EnvVars: []string{"ARTHAS_NATIVE_AGENT_NAME"},
			},
			&cli.StringFlag{
				Name:    "arthas-home-dir",
				Aliases: []string{"d"},
				Usage:   "The arthas home directory",
				EnvVars: []string{"ARTHAS_HOME"},
			},
			&cli.StringFlag{
				Name:    "arthas-version",
				Aliases: []string{"v"},
				Usage:   "The arthas version",
				EnvVars: []string{"ARTHAS_VERSION"},
			},
		},
		Action: func(c *cli.Context) error {
			ip := c.String("tunnel-server-ip")
			port := c.String("tunnel-server-port")
			name := c.String("native-agent-name")
			home := c.String("arthas-home")
			version := c.String("arthas-version")
			return boot.Start(ip, port, name, home, version)
		},
	}
)
