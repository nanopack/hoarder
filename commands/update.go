package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a file in hoarder",
	Long:  ``,

	Run: update,
}

// init
func init() {
	includeAddFlags(updateCmd)
}

// update utilizes the api to update a key with specified data (`add` but with a `PUT` instead of a `POST`)
func update(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Fprintln(os.Stderr, "Missing key - please provide the key for the record you'd like to update")
		return
	case data == "":
		fmt.Fprintln(os.Stderr, "Missing data - please provide the data that you would like to update")
		return
	}

	// use stdin as data to send if "-" is specified on command line `-d -`
	if data == "-" {
		body = os.Stdin
	} else {
		// todo: no buffer?
		body = bytes.NewBuffer([]byte(data))
	}

	// if file is specified, use that instead of any `-d`
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Failed to open file to read - %v\n", err)
			return
		}
		defer f.Close()
		body = f
	}

	b := rest("PUT", "/"+key, body)

	io.Copy(os.Stdout, b)
}
