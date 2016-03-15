package main

import (
	"fmt"
	"os"

	"github.com/nanopack/hoarder/commands"
)

//
func main() {

	//
	if err := commands.HoarderCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
