package commands

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/config"
)

var (

	// alias for show
	fetchCmd = &cobra.Command{
		Hidden: true,

		Use:   "fetch",
		Short: "Display a file from the hoarder storage",
		Long:  ``,

		Run: remove,
	}

	// alias for show
	getCmd = &cobra.Command{
		Hidden: true,

		Use:   "get",
		Short: "Display a file from the hoarder storage",
		Long:  ``,

		Run: remove,
	}

	//
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Display a file from the hoarder storage",
		Long:  ``,

		Run: show,
	}
)

// init
func init() {
	fetchCmd.Flags().StringVarP(&key, "key", "k", "", "The key to get the data by")
	getCmd.Flags().StringVarP(&key, "key", "k", "", "The key to get the data by")
	showCmd.Flags().StringVarP(&key, "key", "k", "", "The key to get the data by")
}

// show
func show(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		config.Log.Error("Missing key - please provide the key for the record you'd like to create")
		return
	}

	config.Log.Debug("Showing: %s", fmt.Sprintf("%s/blobs/%s", config.URI, key))

	//
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blobs/%s", config.URI, key), nil)
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
