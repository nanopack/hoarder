//
package config

import (
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/jcelliott/lumber"

	"github.com/nanopack/hoarder/backends"
)

//
const (
	DEFAULT_BACKEND  = "filesystem"
	DEFAULT_HOST     = "0.0.0.0"
	DEFAULT_LOGLEVEL = "info"
	DEFAULT_PORT     = ":7410"
	DEFAULT_TOKEN    = "TOKEN"
	VERSION          = "0.0.1"
)

//
var (

	//  configurable options
	Backend    = DEFAULT_BACKEND  //
	Driver     backends.Driver    //
	GCInterval = 0                //
	GCAmount   = 0                //
	Host       = DEFAULT_HOST     //
	Insecure   = false            //
	LogLevel   = DEFAULT_LOGLEVEL //
	Port       = DEFAULT_PORT     //
	Token      = DEFAULT_TOKEN    //

	// internal options
	Addr = Host + Port //
	Log  lumber.Logger //
)

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
			case "backend":
				Backend = v
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

	// create a new logger
	Log = lumber.NewConsoleLogger(lumber.LvlInt(LogLevel))
	Log.Prefix("[hoarder]")

	return nil
}
