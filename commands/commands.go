package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/api"
	"github.com/nanopack/hoarder/config"
)

var (

	//
	conf    string //
	server  bool   //
	version bool   //

	//
	key  string //
	data string //

	//
	HoarderCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,

		// parse the config if one is provided, or use the defaults. Set the backend
		// driver to be used
		PersistentPreRun: func(ccmd *cobra.Command, args []string) {

			// if --config is passed, attempt to parse the config file
			if conf != "" {
				config.Parse(conf)
			}

			// update any dependencies that may need to change due to config values
			// or flags
			config.Update()
		},

		// either run hoarder as a server, or run it as a CLI depending of what flags
		// are provided
		Run: func(ccmd *cobra.Command, args []string) {

			// if --server is passed start the hoarder server
			if server != false {
				fmt.Printf("Starting hoarder server at '%s', listening on port '%s'...\n", config.Host, config.Port)

				// start the API
				if err := api.Start(); err != nil {
					config.Log.Error("Failed to start - ", err.Error())
					os.Exit(1)
				}
			}

			// fall back on default help if no args/flags are passed
			ccmd.HelpFunc()(ccmd, args)
		},
	}
)

func init() {

	// persistent flags
	HoarderCmd.PersistentFlags().StringVarP(&config.Connection, "connection", "c", config.DEFAULT_CONNECTION, "Hoarder backend driver")
	HoarderCmd.PersistentFlags().StringVarP(&config.Host, "host", "H", config.DEFAULT_HOST, "Hoarder hostname/IP")
	HoarderCmd.PersistentFlags().BoolVarP(&config.Insecure, "insecure", "i", false, "Disable tls key checking")
	HoarderCmd.PersistentFlags().StringVarP(&config.LogLevel, "log-level", "", config.DEFAULT_LOGLEVEL, "Hoarder output log level")
	HoarderCmd.PersistentFlags().StringVarP(&config.Port, "port", "p", config.DEFAULT_PORT, "Hoarder port")
	HoarderCmd.PersistentFlags().StringVarP(&config.Token, "token", "t", config.DEFAULT_TOKEN, "Hoarder auth token")

	// local flags
	HoarderCmd.Flags().StringVarP(&conf, "config", "", "", "Path to config options")
	HoarderCmd.Flags().BoolVarP(&server, "server", "", false, "Run hoader as a server")
	HoarderCmd.Flags().BoolVarP(&version, "version", "v", false, "Display the current version of this CLI")

	// commands
	HoarderCmd.AddCommand(addCmd)
	HoarderCmd.AddCommand(listCmd)
	HoarderCmd.AddCommand(removeCmd)
	HoarderCmd.AddCommand(showCmd)
	HoarderCmd.AddCommand(updateCmd)

	// hidden/aliased commands
	HoarderCmd.AddCommand(createCmd)
	HoarderCmd.AddCommand(deleteCmd)
	HoarderCmd.AddCommand(destroyCmd)
	HoarderCmd.AddCommand(fetchCmd)
	HoarderCmd.AddCommand(getCmd)
}
