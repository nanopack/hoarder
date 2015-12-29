package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	nanoauth "github.com/nanobox-io/golang-nanoauth"

	"github.com/nanopack/hoarder/config"
)

type object struct {
	Name     string
	CheckSum string
	ModTime  time.Time
	Size     int64
}

// Start
func Start() error {

	// blocking...
	return nanoauth.ListenAndServeTLS(config.Addr, config.Token, routes())
}

// routes registers all api routes with the router
func routes() *pat.Router {
	config.Log.Debug("[hoarder/api] Registering routes...\n")

	//
	router := pat.New()

	//
	router.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong"))
	})

	// blobs
	router.Add("HEAD", "/blobs/{blob}", handleRequest(getBlobHead))
	router.Get("/blobs/{blob}", handleRequest(getBlob))
	router.Get("/blobs", handleRequest(listBlobs))
	router.Post("/blobs/{blob}", handleRequest(createBlob))
	router.Put("/blobs/{blob}", handleRequest(createBlob))
	router.Delete("/blobs/{blob}", handleRequest(deleteBlob))

	return router
}

// handleRequest is a wrapper for the actual route handler, simply to provide some
// debug output
func handleRequest(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		config.Log.Debug(`
Request:
--------------------------------------------------------------------------------
%+v
`, req)

		//
		fn(rw, req)

		config.Log.Debug(`
Response:
--------------------------------------------------------------------------------
%+v
`, rw)
	}
}

// parseBody
// func parseBody(req *http.Request, v interface{}) error {
//
// 	//
// 	b, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		return err
// 	}
//
// 	defer req.Body.Close()
//
// 	//
// 	if err := json.Unmarshal(b, v); err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// writeBody
func writeBody(v interface{}, rw http.ResponseWriter, status int) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(b)

	return nil
}
