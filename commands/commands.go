package commands

import (
	"os"
	"path/filepath"
	"strings"
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

	// HoarderCmd ...
	HoarderCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,

		// parse the config if one is provided, or use the defaults. Set the backend
		// driver to be used
		PersistentPreRun: func(ccmd *cobra.Command, args []string) {

			// convert the log level
			logLvl := lumber.LvlInt(viper.GetString("log-level"))

			// configure the logger
			// lumber.Prefix("[hoader]")
			switch viper.GetString("log-type") {
			case "stdout":
				lumber.Level(logLvl)
			case "file":
				// logger := lumber.NewFileLogger(viper.GetString("log-file"), logLvl, lumber.ROTATE, 5000, 1, 100)
				// lumber.SetLogger(logger)
			}

			// if --config is passed, attempt to parse the config file
			if config != "" {

				// get the filepath
				abs, err := filepath.Abs(config)
				if err != nil {
					lumber.Error("Error reading filepath: ", err.Error())
				}

				// get the config name
				base := filepath.Base(abs)

				// get the path
				path := filepath.Dir(abs)

				//
				viper.SetConfigName(strings.Split(base, ".")[0])
				viper.AddConfigPath(path)

				// Find and read the config file; Handle errors reading the config file
				if err := viper.ReadInConfig(); err != nil {
					lumber.Fatal("Failed to read config file: ", err.Error())
					os.Exit(1)
				}
			}
		},

		// either run hoarder as a server, or run it as a CLI depending on what flags
		// are provided
		Run: func(ccmd *cobra.Command, args []string) {

			// if --server is passed start the hoarder server
			if daemon {

				// enable/start garbage collection if age config was changed
				if ccmd.Flag("clean-after").Changed {
					lumber.Debug("Starting garbage collector (data older than %vs)...\n", ccmd.Flag("clean-after").Value)

					viper.Set("garbage-collect", true)

					// start garbage collector
					go api.StartCollection()
				}

				// start the API
				if err := api.Start(); err != nil {
					lumber.Fatal("Failed to start API: ", err.Error())
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
	viper.SetDefault("backend", "file://") // this is set to file here, and then in the filesystem backend defaults to the actual path
	viper.SetDefault("clean-after", uint64(time.Now().Unix()))
	viper.SetDefault("garbage-collect", false)
	viper.SetDefault("host", "127.0.0.1")
	viper.SetDefault("insecure", true)
	viper.SetDefault("log-type", "stdout")
	viper.SetDefault("log-file", "/var/log/hoarder.log")
	viper.SetDefault("log-level", "INFO")
	viper.SetDefault("port", "7410")
	viper.SetDefault("token", "")

	// persistent flags
	HoarderCmd.PersistentFlags().StringP("backend", "b", viper.GetString("backend"), "Hoarder backend driver")
	HoarderCmd.PersistentFlags().IntP("clean-after", "g", viper.GetInt("clean-after"), "Age, in seconds, after which data is deemed garbage")
	HoarderCmd.PersistentFlags().StringP("host", "H", viper.GetString("host"), "Hoarder hostname/IP")
	HoarderCmd.PersistentFlags().BoolP("insecure", "i", viper.GetBool("insecure"), "Whether or not to start the Hoarder server with TLS")
	HoarderCmd.PersistentFlags().String("log-type", viper.GetString("log-type"), "The type of logging (stdout, file)")
	HoarderCmd.PersistentFlags().String("log-file", viper.GetString("log-file"), "If log-type=file, the /path/to/logfile; ignored otherwise")
	HoarderCmd.PersistentFlags().String("log-level", viper.GetString("log-level"), "Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")
	HoarderCmd.PersistentFlags().StringP("port", "p", viper.GetString("port"), "Hoarder port")
	HoarderCmd.PersistentFlags().StringP("token", "t", viper.GetString("token"), "Auth token used when connecting to a secure Hoarder")

	//
	viper.BindPFlag("backend", HoarderCmd.PersistentFlags().Lookup("backend"))
	viper.BindPFlag("clean-after", HoarderCmd.PersistentFlags().Lookup("clean-after"))
	viper.BindPFlag("host", HoarderCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("insecure", HoarderCmd.PersistentFlags().Lookup("insecure"))
	viper.BindPFlag("log-type", HoarderCmd.PersistentFlags().Lookup("log-type"))
	viper.BindPFlag("log-file", HoarderCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("log-level", HoarderCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("port", HoarderCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("token", HoarderCmd.PersistentFlags().Lookup("token"))

	// local flags;
	HoarderCmd.Flags().StringVar(&config, "config", "", "/path/to/config.yml")
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
