package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display a file from the hoarder storage",
	Long:  ``,

	Run: show,
}

// show
func show(ccmd *cobra.Command, args []string) {
}
