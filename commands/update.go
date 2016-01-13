package commands

import (
	"bytes"
	"fmt"
	"os"
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
		config.Log.Error("Missing key - please provide the key for the record you'd like to update")
		return
	case data == "":
		config.Log.Error("Missing data - please provide the data that you would like to update")
		return
	}

	config.Log.Debug("Updating: %s", fmt.Sprintf("%s/blobs/%s", config.URI, key))

	//
	body := bytes.NewBuffer([]byte(data))

	//
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/blobs/%s", config.URI, key), body)
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
