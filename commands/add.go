package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

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

		Use:   "add",
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
		fmt.Println("Missing key - please provide the key for the record you'd like to create")
		return
	case data == "":
		fmt.Println("Missing data - please provide the data that you would like to create")
		return
	}

	//
	body := bytes.NewBuffer([]byte(args[1]))

	//
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/blobs", config.URI), body)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	//
	req.Header.Add("X-NANOBOX-KEY", args[0])
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

	fmt.Println("ADD??", string(b))
}
