package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file to hoarder storage",
	Long:  ``,

	Run: add,
}

func add(ccmd *cobra.Command, args []string) {
}
