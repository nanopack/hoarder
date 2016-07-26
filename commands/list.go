package commands

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in hoarder storage",
	Long:  ``,

	Run: list,
}

func init() {
	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information about request")
	listCmd.Flags().BoolVarP(&insecure, "insecure", "i", insecure, "Whether or not to ignore hoarder certificate.")
}

// list utilizes the api to retrieve a list of all keys with associated info
func list(ccmd *cobra.Command, args []string) {
	io.Copy(os.Stdout, rest("GET", "/", nil))
}
