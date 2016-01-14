//
package config

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ghodss/yaml"
	"github.com/jcelliott/lumber"
)

//
const (
	VERSION = "0.0.1"
)

//
var (
	// configurable options
	// time.Now().Unix() ensures safety if a config file is used to turn on gc
	CleanAfter = uint64(time.Now().Unix()) // the age that data is deemed garbage (seconds)
	Connection = "file://"                 // the pluggable backend the api will use for storage
	Host       = "127.0.0.1"               // the connection host
	Insecure   = true                      // connect insecurly
	LogLevel   = "info"                    // the output log level
	Port       = "7410"                    // the connection port
	Token      = "TOKEN"                   // the secury token used to connect with

	// internal options
	GarbageCollect = false             // to clean or not to clean
	Addr           = Host + ":" + Port // the host:port connection
	URI            = "https://" + Addr // the connection URI
	Log            lumber.Logger       // the logger to use
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
		options := make(map[string]string)
		if err := yaml.Unmarshal(f, &options); err != nil {
			return err
		}

		// override defaults
		for k, v := range options {
			switch k {
			case "clean_after":
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				CleanAfter = uint64(i)
			case "collect":
				b, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				GarbageCollect = b
			case "connection":
				Connection = v
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
