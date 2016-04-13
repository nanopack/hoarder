package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nanopack/hoarder/util"
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

// show utilizes the api to show data associated to key
func show(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Println("Missing key - please provide the key for the record you'd like to create")
		return
	}

	fmt.Printf("Showing: %s/blobs/%s\n", util.GetURI(), key)

	//
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blobs/%s", util.GetURI(), key), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	//
	req.Header.Add("X-AUTH-TOKEN", viper.GetString("token"))

	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// most often occurs due to server not listening, Exit to keep output clean
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	//
	fmt.Print(string(b))
}
