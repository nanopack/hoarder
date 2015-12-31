package main

import (
	"crypto/tls"
	"net/http"

	"github.com/nanopack/hoarder/commands"
)

//
func main() {

	//
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//
	commands.HoarderCmd.Execute()
}
