package main_test

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jcelliott/lumber"

	"github.com/nanopack/hoarder/config"
	"github.com/nanopack/hoarder/api"
)

func TestMain(m *testing.M) {
	// manually configure
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.Insecure}
	config.Log = lumber.NewConsoleLogger(lumber.LvlInt("info"))
	config.Log.Prefix("[hoarder]")
	config.Connection = "file:///tmp/hoarder_test"
	config.Addr = "127.0.0.1:7411"

	// empty test dir
	os.RemoveAll("/tmp/hoarder_test")

	// start api
	go api.Start()
	<-time.After(time.Second)
	rtn := m.Run()

	// clean test dir
	os.RemoveAll("/tmp/hoarder_test")

	os.Exit(rtn)
}

// test adding data
func TestAddData(t *testing.T) {
	key, data := "test", "data"

	body := bytes.NewBuffer([]byte(data))

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/blobs/%s", config.Addr, key), body)
	req.Header.Add("X-NANOBOX-TOKEN", "TOKEN")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Unable to ADD data - ", err)
		return
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
	if string(b) != "'test' created!\n" {
		t.Errorf("%q doesn't match expected out", b)
	}
}

// test updating data
func TestUpdateData(t *testing.T) {
	key, data := "test2", "data2"

	body := bytes.NewBuffer([]byte(data))

	req, _ := http.NewRequest("PUT", fmt.Sprintf("https://%s/blobs/%s", config.Addr, key), body)
	req.Header.Add("X-NANOBOX-TOKEN", "TOKEN")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Unable to ADD data - ", err)
		return
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
	if string(b) != "'test2' created!\n" {
		t.Errorf("%q doesn't match expected out", b)
	}
}

// test showing data
func TestShowData(t *testing.T) {
	key:= "test"

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s/blobs/%s", config.Addr, key), nil)
	req.Header.Add("X-NANOBOX-TOKEN", "TOKEN")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Unable to SHOW data - ", err)
		return
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
	if string(b) != "data" {
		t.Errorf("%q doesn't match expected out", b)
	}
}

// test removing data
func TestRemoveData(t *testing.T) {
	key:= "test"

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("https://%s/blobs/%s", config.Addr, key), nil)
	req.Header.Add("X-NANOBOX-TOKEN", "TOKEN")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Unable to REMOVE data - ", err)
		return
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
	if string(b) != "'test' destroyed!\n" {
		t.Errorf("%q doesn't match expected out", b)
	}
}

// test listing data
func TestListData(t *testing.T) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s/blobs", config.Addr), nil)
	req.Header.Add("X-NANOBOX-TOKEN", "TOKEN")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Unable to LIST data - ", err)
		return
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
	if string(b) != "[{\"Name\":\"test2\",\"Size\":5}]" {
		t.Errorf("%q doesn't match expected out", b)
	}
}
