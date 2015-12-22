package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Add file to hoarder storage",
	Long:  ``,

	Run: show,
}

func show(ccmd *cobra.Command, args []string) {
}
