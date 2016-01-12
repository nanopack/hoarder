package commands

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/config"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in hoarder storage",
	Long:  ``,

	Run: list,
}

// list
func list(ccmd *cobra.Command, args []string) {

	//
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blobs", config.URI), nil)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	//
	req.Header.Add("X-NANOBOX-TOKEN", config.Token)

	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERR!!", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	fmt.Print("LIST??", string(b))
}
