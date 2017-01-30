package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	//
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add file to hoarder storage",
		Long:  ``,

		Run: add,
	}

	// alias for add
	createCmd = &cobra.Command{
		Hidden: true,

		Use:   "create",
		Short: "Add file to hoarder storage",
		Long:  ``,

		Run: add,
	}
)

// init
func init() {
	includeAddFlags(addCmd)
	includeAddFlags(createCmd)
}

func includeAddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&key, "key", "k", "", "The key to store the data by")
	cmd.Flags().StringVarP(&data, "data", "d", "", "The raw data to be stored")
	cmd.Flags().StringVarP(&file, "file", "f", "", "The filename of the raw data to be stored")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information about request")
	cmd.Flags().BoolVarP(&insecure, "insecure", "i", insecure, "Whether or not to ignore hoarder certificate.")
}

// add utilizes the api to add data corresponding to a specified key
func add(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Fprintln(os.Stderr, "Missing key - please provide the key for the record you'd like to create")
		return
	case data == "" && file == "":
		fmt.Fprintln(os.Stderr, "Missing data - please provide the data that you would like to create")
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
			fmt.Fprintf(os.Stderr, "Failed to open file to read - %s\n", err)
			return
		}
		defer f.Close()
		body = f
	}

	io.Copy(os.Stdout, rest("POST", "/"+key, body))
}
