package main

import (
	"github.com/urfave/cli/v2"

	"github.com/gstolarz/drone-mikrotik/plugin"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "address",
			Usage:       "address",
			EnvVars:     []string{"PLUGIN_ADDRESS"},
			Destination: &settings.Address,
		},
		&cli.BoolFlag{
			Name:        "tls",
			Usage:       "tls",
			EnvVars:     []string{"PLUGIN_TLS"},
			Destination: &settings.TLS,
		},
		&cli.BoolFlag{
			Name:        "insecure",
			Usage:       "insecure",
			EnvVars:     []string{"PLUGIN_INSECURE"},
			Destination: &settings.Insecure,
		},
		&cli.StringFlag{
			Name:        "username",
			Usage:       "username",
			EnvVars:     []string{"PLUGIN_USERNAME"},
			Destination: &settings.Username,
		},
		&cli.StringFlag{
			Name:        "password",
			Usage:       "password",
			EnvVars:     []string{"PLUGIN_PASSWORD"},
			Destination: &settings.Password,
		},
		&cli.StringSliceFlag{
			Name:        "script",
			Usage:       "script",
			EnvVars:     []string{"PLUGIN_SCRIPT"},
			Destination: &settings.Script,
		},
	}
}
