package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Add file to hoarder storage",
	Long:  ``,

	Run: server,
}

func server(ccmd *cobra.Command, args []string) {
}
