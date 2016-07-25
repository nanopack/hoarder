package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (

	// alias for show
	fetchCmd = &cobra.Command{
		Hidden: true,

		Use:   "fetch",
		Short: "Display a file from the hoarder storage",
		Long:  ``,

		Run: show,
	}

	// alias for show
	getCmd = &cobra.Command{
		Hidden: true,

		Use:   "get",
		Short: "Display a file from the hoarder storage",
		Long:  ``,

		Run: show,
	}

	//
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Display a file from the hoarder storage",
		Long:  ``,

		Run: show,
	}
)

// init
func init() {
	includeShowFlags(fetchCmd)
	includeShowFlags(getCmd)
	includeShowFlags(showCmd)
}

func includeShowFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&key, "key", "k", "", "The key to get the data by")
	cmd.Flags().StringVarP(&file, "file", "f", "", "The filename to save the raw data to")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information about request")
	cmd.Flags().BoolVarP(&insecure, "insecure", "i", insecure, "Whether or not to verify hoarder certificate.")
}

// show utilizes the api to show data associated to key
func show(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Fprintln(os.Stderr, "Missing key - please provide the key for the record you'd like to get")
		return
	}

	var out io.Writer

	out = os.Stdout
	if file != "" {
		f, err := os.Create(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file to write - %v\n", err)
			return
		}
		defer f.Close()
		out = f
	}

	io.Copy(out, rest("GET", "/"+key, nil))

	if file != "" {
		fmt.Fprintln(os.Stderr, "Finished writing file")
	}
}
