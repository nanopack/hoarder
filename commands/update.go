package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/config"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a file in hoarder",
	Long:  ``,

	Run: update,
}

// init
func init() {
	updateCmd.Flags().StringVarP(&key, "key", "k", "", "The key to store the data by")
	updateCmd.Flags().StringVarP(&data, "data", "d", "", "The raw data to be stored")
}

// update
func update(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Println("Missing key - please provide the key for the record you'd like to update")
	case data == "":
		fmt.Println("Missing data - please provide the data that you would like to update")
	}

	//
	body := bytes.NewBuffer([]byte(args[1]))

	//
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/blobs/%s", config.URI, args[0]), body)
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

	fmt.Println("UPDATE??", string(b))
}
