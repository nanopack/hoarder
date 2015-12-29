package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a file in hoarder",
	Long:  ``,

	Run: update,
}

// update
func update(ccmd *cobra.Command, args []string) {
}
