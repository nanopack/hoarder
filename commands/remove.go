package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (

	// alias for remove
	deleteCmd = &cobra.Command{
		Hidden: true,

		Use:   "delete",
		Short: "Remove a file from hoarder storage",
		Long:  ``,

		Run: remove,
	}

	// alias for remove
	destroyCmd = &cobra.Command{
		Hidden: true,

		Use:   "destroy",
		Short: "Remove a file from hoarder storage",
		Long:  ``,

		Run: remove,
	}

	//
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a file from hoarder storage",
		Long:  ``,

		Run: remove,
	}
)

// init
func init() {
	includeRemoveFlags(deleteCmd)
	includeRemoveFlags(destroyCmd)
	includeRemoveFlags(removeCmd)
}

func includeRemoveFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&key, "key", "k", "", "The key to remove the data by")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information about request")
	cmd.Flags().BoolVarP(&insecure, "insecure", "i", insecure, "Whether or not to verify hoarder certificate.")
}

// remove utilizes the api to remove a key and associated data
func remove(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Fprintln(os.Stderr, "Missing key - please provide the key for the record you'd like to remove")
		return
	}

	io.Copy(os.Stdout, rest("DELETE", "/"+key, nil))
}
