package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/config"
)

var (

	// alias for remove
	deleteCmd = &cobra.Command{
		Hidden: true,

		Use:   "delete",
		Short: "Remove a file from hoarder storage",
		Long:  ``,

		Run: remove,
	}

	// alias for remove
	destroyCmd = &cobra.Command{
		Hidden: true,

		Use:   "destroy",
		Short: "Remove a file from hoarder storage",
		Long:  ``,

		Run: remove,
	}

	//
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a file from hoarder storage",
		Long:  ``,

		Run: remove,
	}
)

// init
func init() {
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "The key to remove the data by")
	destroyCmd.Flags().StringVarP(&key, "key", "k", "", "The key to remove the data by")
	removeCmd.Flags().StringVarP(&key, "key", "k", "", "The key to remove the data by")
}

// remove
func remove(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		config.Log.Error("Missing key - please provide the key for the record you'd like to create")
		return
	}

	config.Log.Debug("Removing: %s", fmt.Sprintf("%s/blobs/%s", config.URI, key))

	//
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/blobs/%s", config.URI, key), nil)
	if err != nil {
		config.Log.Error(err.Error())
	}

	//
	req.Header.Add("X-NANOBOX-TOKEN", config.Token)

	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		config.Log.Fatal(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		config.Log.Error(err.Error())
	}

	fmt.Print(string(b))
}
