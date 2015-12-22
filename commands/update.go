package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Add file to hoarder storage",
	Long:  ``,

	Run: update,
}

func update(ccmd *cobra.Command, args []string) {
}
