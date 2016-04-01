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

// remove utilizes the api to remove a key and associated data
func remove(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Println("Missing key - please provide the key for the record you'd like to create")
		return
	}

	fmt.Printf("Removing: %s/blobls/%s\n", util.GetURI(), key)

	//
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/blobs/%s", util.GetURI(), key), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	//
	req.Header.Add("x-auth-token", viper.GetString("token"))

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
