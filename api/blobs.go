package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// get
func get(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get(":blob")

	//
	r, err := driver.Read(key)
	if err != nil {
		writeBody(err.Error(), rw)
		return
	}

	//
	b, err := ioutil.ReadAll(r)
	if err != nil {
		writeBody(err.Error(), rw)
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
		writeBody(err.Error(), rw)
		return
	}

	//
	writeBody(fi, rw)
}

// create
func create(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get(":blob")

	//
	if err := driver.Write(key, req.Body); err != nil {
		writeBody(err.Error(), rw)
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
		writeBody(err.Error(), rw)
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
		writeBody(err.Error(), rw)
		return
	}

	//
	writeBody(fis, rw)
}
