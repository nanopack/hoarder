// Hoarder is a simple, api-driven, storage system.
//
// To run as a server, using the defaults, starting hoarder is as simple as
//  hoarder -s
// For more specific usage information, refer to the help doc (hoarder -h):
//
//  hoarder - data storage
//
//  Usage:
//    hoarder [flags]
//    hoarder [command]
//
//  Available Commands:
//    add         Add file to hoarder storage
//    list        List all files in hoarder storage
//    remove      Remove a file from hoarder storage
//    show        Display a file from the hoarder storage
//    update      Update a file in hoarder
//
//  Flags:
//    -b, --backend string       Hoarder backend (default "file:///var/db/hoarder")
//    -g, --clean-after uint     Age, in seconds, after which data is deemed garbage (default 0)
//    -c, --config string        Path to config file
//    -H, --listen-addr string   Hoarder listen uri (scheme defaults to https) (default "https://127.0.0.1:7410")
//        --log-level string     Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL) (default "INFO")
//    -s, --server               Run hoarder as a server
//    -t, --token string         Auth token used when connecting to a secure Hoarder
//    -v, --version              Display the current version of this CLI
//
//  Use "hoarder [command] --help" for more information about a command.
//
//
package main

import (
	"fmt"

	"github.com/nanopack/hoarder/commands"
)

func main() {
	err := commands.HoarderCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
