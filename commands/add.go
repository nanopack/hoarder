package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/nanopack/hoarder/config"
)

var (

	//
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add file to hoarder storage",
		Long:  ``,

		Run: add,
	}

	// alias for add
	createCmd = &cobra.Command{
		Hidden: true,

		Use:   "create",
		Short: "Add file to hoarder storage",
		Long:  ``,

		Run: add,
	}
)

// init
func init() {
	addCmd.Flags().StringVarP(&key, "key", "k", "", "The key to store the data by")
	addCmd.Flags().StringVarP(&data, "data", "d", "", "The raw data to be stored")

	createCmd.Flags().StringVarP(&key, "key", "k", "", "The key to store the data by")
	createCmd.Flags().StringVarP(&data, "data", "d", "", "The raw data to be stored")
}

// add
func add(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		config.Log.Error("Missing key - please provide the key for the record you'd like to create")
		return
	case data == "":
		config.Log.Error("Missing data - please provide the data that you would like to create")
		return
	}

	config.Log.Debug("Adding: %s", fmt.Sprintf("%s/blobs/%s", config.URI, key))

	//
	body := bytes.NewBuffer([]byte(data))

	//
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/blobs/%s", config.URI, key), body)
	if err != nil {
		config.Log.Error(err.Error())
	}

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
