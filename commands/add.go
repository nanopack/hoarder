package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	key  string
	data string

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

// add utilizes the api to add data corresponding to a specified key
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

	fmt.Printf("Adding: %s/blobs/%s\n", uri, key)

	//
	body := bytes.NewBuffer([]byte(data))

	//
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/blobs/%s", uri, key), body)
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

	fmt.Print(string(b))
}
