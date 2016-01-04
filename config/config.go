//
package config

import (
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/jcelliott/lumber"
)

//
const (
	DEFAULT_CONNECTION = "file:///var/db/hoarder"
	DEFAULT_HOST       = "0.0.0.0"
	DEFAULT_LOGLEVEL   = "info"
	DEFAULT_PORT       = ":7410"
	DEFAULT_TOKEN      = "TOKEN"
	VERSION            = "0.0.1"
)

//
var (

	//  configurable options
	Connection = DEFAULT_CONNECTION // the pluggable backend the api will use for storage
	GCInterval = 0                  // the interval between clearning out old storage
	GCAmount   = 0                  // the amount of storage to clear at interval
	Host       = DEFAULT_HOST       // the connection host
	Insecure   = false              // connect insecurly
	LogLevel   = DEFAULT_LOGLEVEL   // the output log level
	Port       = DEFAULT_PORT       // the connection port
	Token      = DEFAULT_TOKEN      // the secury token used to connect with

	// internal options
	Addr = Host + Port       // the host:port connection
	URI  = "https://" + Addr // the connection URI
	Log  lumber.Logger       // the logger to use
)

//
func init() {

	// create a new logger
	Log = lumber.NewConsoleLogger(lumber.LvlInt(LogLevel))
	Log.Prefix("[hoarder]")
}

//
func Parse(path string) error {

	// if a config is provided (and found), parse the config file overwriting any
	// defaults
	if fp, err := filepath.Abs(path); err == nil {

		//
		f, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}

		// parse config file
		options := map[string]string{}
		if err := yaml.Unmarshal(f, options); err != nil {
			return err
		}

		// override defaults
		for k, v := range options {
			switch k {
			case "connection":
				Connection = v
			case "gc_interval":
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				GCInterval = i
			case "gc_amount":
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				GCAmount = i
			case "host":
				Host = v
			case "insecure":
				b, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				Insecure = b
			case "log_level":
				LogLevel = v
			case "port":
				Port = v
			case "token":
				Token = v
			}
		}
	}

	return nil
}

// Update updates any dependencies that may need to change due to config options
// or flags
func Update() {

	//
	Log.Level(lumber.LvlInt(LogLevel))

	//
	Addr = Host + Port

	//
	URI = "https://" + Addr
}
