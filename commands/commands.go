package commands

import (
	// "github.com/nanopack/hoarder/config"
	"github.com/spf13/cobra"
)

var (
	HoarderCommand = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,
	}
)

func init() {
	// HoarderCommand.PersistentFlags().StringVarP(&config.AuthToken, "auth", "A", "", "Hoarder auth token")
	// HoarderCommand.PersistentFlags().StringVarP(&config.Host, "host", "H", "127.0.0.1", "Hoarder hostname/IP")
	// HoarderCommand.PersistentFlags().IntVarP(&config.Port, "port", "p", 8443, "Hoarder admin port")
	// HoarderCommand.PersistentFlags().BoolVarP(&config.Insecure, "insecure", "i", false, "Disable tls key checking")

	HoarderCommand.AddCommand(addCmd)
	HoarderCommand.AddCommand(removeCmd)
	HoarderCommand.AddCommand(showCmd)
	HoarderCommand.AddCommand(updateCmd)
	HoarderCommand.AddCommand(listCmd)
	HoarderCommand.AddCommand(serverCmd)
}
