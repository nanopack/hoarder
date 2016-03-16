package commands

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jcelliott/lumber"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nanopack/hoarder/api"
)

var (
	log lumber.Logger

	//
	config  string //
	daemon  bool   //
	version bool   //

	//
	uri string

	//
	HoarderCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,

		// parse the config if one is provided, or use the defaults. Set the backend
		// driver to be used
		PersistentPreRun: func(ccmd *cobra.Command, args []string) {

			// configure the logger
			lumber.Level(lumber.LvlInt(viper.GetString("log_level")))
			// lumber.Prefix("[hoader]")

			// if --config is passed, attempt to parse the config file
			if config != "" {

				//
				viper.SetConfigName("config")
				viper.AddConfigPath(config)

				// Find and read the config file; Handle errors reading the config file
				if err := viper.ReadInConfig(); err != nil {
					panic(fmt.Errorf("Fatal error config file: %s \n", err))
				}
			}

			// configure InsecureSkipVerify using setting from 'insecure' flag
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: viper.GetBool("insecure")}
		},

		// either run hoarder as a server, or run it as a CLI depending on what flags
		// are provided
		Run: func(ccmd *cobra.Command, args []string) {

			// if --server is passed start the hoarder server
			if daemon {

				// enable/start garbage collection if age config was changed
				if ccmd.Flag("clean-after").Changed {
					fmt.Printf("Starting garbage collector (data older than %vs)...\n", ccmd.Flag("clean-after").Value)

					viper.Set("garbage-collect", true)

					// start garbage collector
					go api.StartCollection()
				}

				// start the API
				if err := api.Start(); err != nil {
					fmt.Println("Failed to start API!", err)
					os.Exit(1)
				}
			}

			// fall back on default help if no args/flags are passed
			ccmd.HelpFunc()(ccmd, args)
		},
	}
)

func init() {

	// set config defaults; these are overriden if a --config file is provided
	// (see above)
	viper.SetDefault("backend", "file://")
	viper.SetDefault("clean-after", uint64(time.Now().Unix()))
	viper.SetDefault("garbage-collect", false)
	viper.SetDefault("host", "127.0.0.1")
	viper.SetDefault("insecure", true)
	viper.SetDefault("log-level", "INFO")
	viper.SetDefault("port", "7410")
	viper.SetDefault("token", "")
	viper.SetDefault("uri", fmt.Sprintf("%v:%v", viper.GetString("host"), viper.GetString("port")))

	//
	uri = fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))

	// persistent flags
	HoarderCmd.PersistentFlags().StringP("backend", "b", viper.GetString("backend"), "Hoarder backend driver")
	HoarderCmd.PersistentFlags().IntP("clean-after", "g", viper.GetInt("clean-after"), "Age data is deemed garbage (seconds)")
	HoarderCmd.PersistentFlags().StringP("host", "H", viper.GetString("host"), "Hoarder hostname/IP")
	HoarderCmd.PersistentFlags().BoolP("insecure", "i", viper.GetBool("insecure"), "Disable tls key checking")
	HoarderCmd.PersistentFlags().String("log-level", viper.GetString("log-level"), "Hoarder output log level")
	HoarderCmd.PersistentFlags().StringP("port", "p", viper.GetString("port"), "Hoarder port")
	HoarderCmd.PersistentFlags().StringP("token", "t", viper.GetString("token"), "Hoarder auth token")

	//
	viper.BindPFlag("backend", HoarderCmd.PersistentFlags().Lookup("backend"))
	viper.BindPFlag("clean-after", HoarderCmd.PersistentFlags().Lookup("clean-after"))
	viper.BindPFlag("log-level", HoarderCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("host", HoarderCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("insecure", HoarderCmd.PersistentFlags().Lookup("insecure"))
	viper.BindPFlag("port", HoarderCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("token", HoarderCmd.PersistentFlags().Lookup("token"))

	// local flags;
	HoarderCmd.Flags().StringVar(&config, "config", "", "Path to config options")
	HoarderCmd.Flags().BoolVar(&daemon, "server", false, "Run hoarder as a server")
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
