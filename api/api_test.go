package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/spf13/viper"
)

var (
	testHost = "127.0.0.1"
	testPort = "7411"
	testAddr = fmt.Sprintf("%s:%s", testHost, testPort)
	testKey  = "testKey"
	testData = "testData"
)

func TestMain(m *testing.M) {

	// manually configure
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	viper.Set("backend", "file:///tmp/hoarder_test")
	viper.Set("host", testHost)
	viper.Set("port", testPort)

	// empty test dir
	os.RemoveAll("/tmp/hoarder_test")

	// start api
	go Start()
	<-time.After(time.Second)
	rtn := m.Run()

	// empty test dir
	os.RemoveAll("/tmp/hoarder_test")

	os.Exit(rtn)
}

// TestAuth
func TestAuth(t *testing.T) {

	// set a token
	viper.Set("token", "TOKEN")

	// try and get some data
	res, err := do("GET", testKey, nil)
	if err != nil {
		t.Fatalf("Failed to get data - %v", err.Error())
	}
	defer res.Body.Close()

	// make sure we weren't authorized to do the action
	if res.StatusCode != 401 {
		t.Fatalf("Unauthorized action!")
	}

	// remove the token for the rest of the tests
	viper.Set("token", "")
}

// TestAddData
func TestAddData(t *testing.T) {

	body := bytes.NewBuffer([]byte(testData))

	//
	res, err := do("POST", testKey, body)
	if err != nil {
		t.Fatalf("Failed to add data - %v", err.Error())
	}
	defer res.Body.Close()

	//
	testResponse(res.Body, fmt.Sprintf("'%s' created!\n", testKey), t)
}

// TestUpdateData
func TestUpdateData(t *testing.T) {

	body := bytes.NewBuffer([]byte(testData))

	//
	res, err := do("PUT", testKey, body)
	if err != nil {
		t.Fatalf("Failed to update data - %v", err.Error())
	}
	defer res.Body.Close()

	//
	testResponse(res.Body, fmt.Sprintf("'%s' created!\n", testKey), t)
}

// TestShowData
func TestShowData(t *testing.T) {

	//
	res, err := do("GET", testKey, nil)
	if err != nil {
		t.Fatalf("Failed to show data - %v", err.Error())
	}
	defer res.Body.Close()

	//
	testResponse(res.Body, "testData", t)
}

// TestHeadData
func TestHeadData(t *testing.T) {

	//
	res, err := do("HEAD", testKey, nil)
	if err != nil {
		t.Fatalf("Failed to show data - %v", err.Error())
	}
	defer res.Body.Close()

	//
	length, _ := strconv.Atoi(res.Header.Get("Content-Length"))
	if length != len(testData) {
		t.Errorf("Unexpected length. Expecting %v got '%v'", len(testData), length)
	}
}

// TestListData
func TestListData(t *testing.T) {

	//
	res, err := do("GET", "", nil)
	if err != nil {
		t.Fatalf("Failed to list data - %v", err.Error())
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)

	var list []map[string]interface{}
	json.Unmarshal(b, &list)

	if list[0]["Name"] != testKey {
		t.Errorf("%q doesn't match expected out", b)
	}
}

// TestRemoveData
func TestRemoveData(t *testing.T) {

	//
	res, err := do("DELETE", testKey, nil)
	if err != nil {
		t.Fatalf("Failed to remove data - %v", err.Error())
	}
	defer res.Body.Close()

	//
	testResponse(res.Body, fmt.Sprintf("'%s' destroyed!\n", testKey), t)
}

// do
func do(method, path string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, fmt.Sprintf("https://%s/blobs/%s", testAddr, path), body)
	return http.DefaultClient.Do(req)
}

// testResponse
func testResponse(body io.Reader, expecting string, t *testing.T) {
	b, _ := ioutil.ReadAll(body)
	if string(b) != expecting {
		t.Fatalf("Unexpected response. Expecting '%s' got '%q'", expecting, b)
	}
}
