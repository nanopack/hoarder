package api

import (
	"net/http"

	"github.com/nanopack/hoarder/config"
)

//
func getBlob(rw http.ResponseWriter, req *http.Request) {
	config.Driver.Read(req.URL.Query().Get(":blob"))
}

//
func getBlobHead(rw http.ResponseWriter, req *http.Request) {
	config.Driver.Stat(req.URL.Query().Get(":blob"))
}

//
func createBlob(rw http.ResponseWriter, req *http.Request) {
	// config.Driver.Write(req.URL.Query().Get(":blob"))
}

//
func deleteBlob(rw http.ResponseWriter, req *http.Request) {
	config.Driver.Remove(req.URL.Query().Get(":blob"))
}

//
func listBlobs(rw http.ResponseWriter, req *http.Request) {
	config.Driver.List()
}
