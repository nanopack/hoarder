package commands

import (
	"fmt"
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
		fmt.Println("Missing key - please provide the key for the record you'd like to create")
		return
	}

	//
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blobs/%s", config.URI, key), nil)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	//
	req.Header.Add("X-NANOBOX-TOKEN", config.Token)

	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERR!!", err)
	}
	defer res.Body.Close()

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	fmt.Println("SHOW??", string(b))
}
