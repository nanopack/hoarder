package main

import (
	"github.com/pagodabox/nanobox-config"
	"strings"
	"os"
	"github.com/jcelliott/lumber"
)

var log lumber.Logger
var dataDir string

func main() {
	configFile := ""
	if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		configFile = os.Args[1]
	}

	conf := map[string]string{
		"listenAddr": ":1234",
		"logLevel":    "info",
		"token":       "token",
		"dataDir":      "/tmp/warehouse/",
	}

	config.Load(conf, configFile)
	conf = config.Config
	// do the stuff
	level := lumber.LvlInt(conf["logLevel"])
	log = lumber.NewConsoleLogger(level)
	log.Prefix("[warehouse]")	
	dataDir = conf["dataDir"]

	err := Start(conf["listenAddr"], conf["token"])
	if err != nil {
		log.Error("could not start: %s", err.Error())
		os.Exit(1)
	}
}