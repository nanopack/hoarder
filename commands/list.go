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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in hoarder storage",
	Long:  ``,

	Run: list,
}

// list utilizes the api to retrieve a list of all keys with associated info
func list(ccmd *cobra.Command, args []string) {

	fmt.Printf("Listing: %s/blobs\n", util.GetURI())

	//
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blobs", util.GetURI()), nil)
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
