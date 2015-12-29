package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a file from hoarder storage",
	Long:  ``,

	Run: remove,
}

// remove
func remove(ccmd *cobra.Command, args []string) {
}
