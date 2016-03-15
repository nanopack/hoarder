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

// update utilizes the api to update a key with specified data
func update(ccmd *cobra.Command, args []string) {

	// handle any missing args
	switch {
	case key == "":
		fmt.Println("Missing key - please provide the key for the record you'd like to update")
		return
	case data == "":
		fmt.Println("Missing data - please provide the data that you would like to update")
		return
	}

	fmt.Printf("Updating: %s/blobs/%s\n", uri, key)

	//
	body := bytes.NewBuffer([]byte(data))

	//
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/blobs/%s", uri, key), body)
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
