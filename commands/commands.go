// Package commands is where all cli logic is, including starting hoarder as a server.
package commands

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jcelliott/lumber"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nanopack/hoarder/api"
	"github.com/nanopack/hoarder/backends"
	"github.com/nanopack/hoarder/collector"
)

var (
	body     io.ReadWriter        // what to read/write requests body
	verbose  bool                 // whether to display request info
	insecure bool          = true // whether to ignore cert or not

	key  string // blob key
	data string // blob raw data (or '-' for stdin)
	file string // blob file location

	config   string // config file location
	daemon   bool   // whether to run as server or not
	showVers bool   // whether to print version info or not

	// to be populated by linker
	version string
	commit  string

	// HoarderCmd ...
	HoarderCmd = &cobra.Command{
		Use:           "hoarder",
		Short:         "hoarder - data storage",
		Long:          ``,
		SilenceErrors: true,
		SilenceUsage:  true,

		// parse the config if one is provided, or use the defaults. Set the backend
		// driver to be used
		PersistentPreRunE: readConfig,

		// print version or help, or continue, depending on flag settings
		PreRunE: preFlight,

		// either run hoarder as a server, or run it as a CLI depending on what flags
		// are provided
		RunE: startHoarder,
	}
)

func readConfig(ccmd *cobra.Command, args []string) error {
	// if --config is passed, attempt to parse the config file
	if config != "" {
		filename := filepath.Base(config)
		viper.SetConfigName(filename[:len(filename)-len(filepath.Ext(filename))])
		viper.AddConfigPath(filepath.Dir(config))

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("ERROR: Failed to read config file: %s\n", err.Error())
			return fmt.Errorf("")
		}
	}

	return nil
}

func preFlight(ccmd *cobra.Command, args []string) error {
	// if --version is passed print the version info
	if showVers {
		fmt.Printf("hoarder %s (%s)\n", version, commit)
		return fmt.Errorf("")
	}

	// if --server is not passed, print help
	if !viper.GetBool("server") {
		ccmd.HelpFunc()(ccmd, args)
		return fmt.Errorf("") // no error, just exit
	}

	return nil
}

func startHoarder(ccmd *cobra.Command, args []string) error {
	// convert the log level
	logLvl := lumber.LvlInt(viper.GetString("log-level"))

	// configure the logger
	lumber.Prefix("[hoader]")
	lumber.Level(logLvl)

	// enable/start garbage collection if age config was changed
	if ccmd.Flag("clean-after").Changed {
		lumber.Debug("Starting garbage collector (data older than %vs)...\n", ccmd.Flag("clean-after").Value)

		// start garbage collector
		go collector.Start()
	}

	// set, and initialize, the backend driver
	if err := backends.Initialize(); err != nil {
		lumber.Error("Failed to initialize backend - %v", err)
		return err
	}

	// start the API
	if err := api.Start(); err != nil {
		lumber.Fatal("Failed to start API: ", err.Error())
		return err
	}

	return nil
}

func init() {
	// set config defaults
	backend := "file:///var/db/hoarder"
	cleanAfter := uint64(time.Now().Unix())
	listenAddr := "https://127.0.0.1:7410"
	logLevel := "INFO"
	token := ""
	server := false

	// cli flags
	HoarderCmd.PersistentFlags().StringP("backend", "b", backend, "Hoarder backend")
	HoarderCmd.PersistentFlags().Uint64P("clean-after", "g", cleanAfter, "Age, in seconds, after which data is deemed garbage")
	HoarderCmd.PersistentFlags().StringP("listen-addr", "H", listenAddr, "Hoarder listen uri (scheme defaults to https)")
	HoarderCmd.PersistentFlags().String("log-level", logLevel, "Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")
	HoarderCmd.PersistentFlags().StringP("token", "t", token, "Auth token used when connecting to a secure Hoarder")
	HoarderCmd.Flags().BoolP("server", "s", server, "Run hoarder as a server")

	// bind config to cli flags
	viper.BindPFlag("backend", HoarderCmd.PersistentFlags().Lookup("backend"))
	viper.BindPFlag("clean-after", HoarderCmd.PersistentFlags().Lookup("clean-after"))
	viper.BindPFlag("listen-addr", HoarderCmd.PersistentFlags().Lookup("listen-addr"))
	viper.BindPFlag("log-level", HoarderCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("token", HoarderCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("server", HoarderCmd.Flags().Lookup("server"))

	// cli-only flags
	HoarderCmd.Flags().StringVarP(&config, "config", "c", "", "Path to config file (with extension)")
	HoarderCmd.Flags().BoolVarP(&showVers, "version", "v", false, "Display the current version of this CLI")

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

func rest(_method, _key string, _body io.ReadWriter) io.Reader {
	if verbose {
		fmt.Fprintf(os.Stderr, "%s: %s/blobs%s", _method, viper.GetString("listen-addr"), _key)
	}

	req, err := http.NewRequest(_method, fmt.Sprintf("%s/blobs%s", viper.GetString("listen-addr"), _key), _body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create request")
	}

	// skip cert check or not (only applies to https)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}

	// set auth token
	req.Header.Add("X-AUTH-TOKEN", viper.GetString("token"))

	// make request to hoarder server
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// print error and exit
		fmt.Fprintf(os.Stderr, "Failed to make request - %v\n", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "%5d\n", res.StatusCode)
	}

	return res.Body
}
