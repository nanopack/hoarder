package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Add file to hoarder storage",
	Long:  ``,

	Run: remove,
}

func remove(ccmd *cobra.Command, args []string) {
}
