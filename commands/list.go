package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/config"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in hoarder storage",
	Long:  ``,

	Run: list,
}

// list utilizes the api to retrieve a list of all keys with associated info
func list(ccmd *cobra.Command, args []string) {

	config.Log.Debug("Listing: %s", fmt.Sprintf("%s/blobs", config.URI))

	//
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blobs", config.URI), nil)
	if err != nil {
		config.Log.Error(err.Error())
	}

	//
	req.Header.Add("X-NANOBOX-TOKEN", config.Token)

	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// most often occurs due to server not listening, Exit to keep output clean
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
