package main

import (
	"crypto/tls"
	"net/http"

	"github.com/nanopack/hoarder/commands"
	// "github.com/nanopack/hoarder/config"
	// "fmt"
)

//
func main() {

	//
	// fmt.Println("Insecure flag:  ", config.Insecure)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//
	commands.HoarderCmd.Execute()
}
