package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Add file to hoarder storage",
	Long:  ``,

	Run: list,
}

func list(ccmd *cobra.Command, args []string) {
}
