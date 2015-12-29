package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in hoarder storage",
	Long:  ``,

	Run: list,
}

// list
func list(ccmd *cobra.Command, args []string) {
}
