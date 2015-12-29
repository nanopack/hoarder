package api

import "net/http"

//
func getBlob(rw http.ResponseWriter, req *http.Request) {
	driver.Read(req.URL.Query().Get(":blob"))
}

//
func getBlobHead(rw http.ResponseWriter, req *http.Request) {
	driver.Stat(req.URL.Query().Get(":blob"))
}

//
func createBlob(rw http.ResponseWriter, req *http.Request) {
	// driver.Write(req.URL.Query().Get(":blob"))
}

//
func deleteBlob(rw http.ResponseWriter, req *http.Request) {
	driver.Remove(req.URL.Query().Get(":blob"))
}

//
func listBlobs(rw http.ResponseWriter, req *http.Request) {
	driver.List()
}
