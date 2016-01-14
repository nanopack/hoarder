package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// get
func get(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get(":blob")

	//
	r, err := driver.Read(key)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
		return
	}

	//
	b, err := ioutil.ReadAll(r)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
		return
	}

	//
	rw.Write(b)
}

// getHead
func getHead(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get(":blob")

	//
	fi, err := driver.Stat(key)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
		return
	}

	// set useful headers
	rw.Header().Set("Content-Length", fmt.Sprintf("%d", fi.Size))
	rw.Header().Set("Last-Modified", fi.ModTime.Format(time.RFC1123))
	rw.Header().Set("Date", time.Now().UTC().Format(time.RFC1123))

	//
	writeBody(nil, rw)
}

// create
func create(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get(":blob")

	//
	if err := driver.Write(key, req.Body); err != nil {
		rw.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
		return
	}

	//
	rw.Write([]byte(fmt.Sprintf("'%s' created!\n", key)))
}

// delete
func delete(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get(":blob")

	//
	if err := driver.Remove(key); err != nil {
		rw.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
		return
	}

	//
	rw.Write([]byte(fmt.Sprintf("'%s' destroyed!\n", key)))
}

// list
func list(rw http.ResponseWriter, req *http.Request) {

	//
	fis, err := driver.List()
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	//
	writeBody(fis, rw)
}
