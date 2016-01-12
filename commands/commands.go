package commands

import (
	"os"
	"net/http"
	"crypto/tls"

	"github.com/spf13/cobra"
	"github.com/jcelliott/lumber"

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
			// create a new logger
			config.Log = lumber.NewConsoleLogger(lumber.LvlInt(config.LogLevel))
			config.Log.Prefix("[hoarder]")


			// if --config is passed, attempt to parse the config file
			if conf != "" {
				if err := config.Parse(conf); err != nil {
					config.Log.Error("Failed to parse config '%s' - %s", conf, err.Error())
				}
			}

			// configure InsecureSkipVerify using setting from 'insecure' flag
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.Insecure}
		},

		// either run hoarder as a server, or run it as a CLI depending on what flags
		// are provided
		Run: func(ccmd *cobra.Command, args []string) {

			// if --server is passed start the hoarder server
			if server != false {
				config.Log.Info("Starting hoarder server at '%s', listening on port '%s'...\n", config.Host, config.Port)

				// start the API
				if err := api.Start(); err != nil {
					config.Log.Fatal("Failed to start - %s", err.Error())
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
	HoarderCmd.PersistentFlags().StringVarP(&config.Connection, "connection", "c", config.Connection, "Hoarder backend driver")
	HoarderCmd.PersistentFlags().StringVarP(&config.Host, "host", "H", config.Host, "Hoarder hostname/IP")
	HoarderCmd.PersistentFlags().BoolVarP(&config.Insecure, "insecure", "i", true, "Disable tls key checking")
	HoarderCmd.PersistentFlags().StringVarP(&config.LogLevel, "log-level", "", config.LogLevel, "Hoarder output log level")
	HoarderCmd.PersistentFlags().StringVarP(&config.Port, "port", "p", config.Port, "Hoarder port")
	HoarderCmd.PersistentFlags().StringVarP(&config.Token, "token", "t", config.Token, "Hoarder auth token")

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
