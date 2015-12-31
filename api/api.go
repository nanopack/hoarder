package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/pat"
	nanoauth "github.com/nanobox-io/golang-nanoauth"

	"github.com/nanopack/hoarder/backends"
	"github.com/nanopack/hoarder/config"
)

type (

	//
	Driver interface {
		List() ([]backends.FileInfo, error)
		Read(string) (io.Reader, error)
		Remove(string) error
		Stat(string) (backends.FileInfo, error)
		Write(string, io.Reader) error
	}
)

//
var driver Driver

// Start
func Start() error {

	//
	if err := setDriver(); err != nil {
		fmt.Println("BONK!", err)
		os.Exit(1)
	}

	// blocking...
	return nanoauth.ListenAndServeTLS(config.Addr, config.Token, routes())
}

//
func setDriver() error {

	//
	u, err := url.Parse(config.Connection)
	if err != nil {
		return err
	}

	fmt.Printf("URL!! %#v\n", u)

	//
	switch u.Scheme {
	case "file":
		driver = backends.Filesystem{Path: u.Path}
	// case "s3":
	// 	driver = backends.S3{Path: u.Path}
	// case "mongo":
	// 	driver = backends.Mongo{Path: u.Path}
	// case "redis":
	// 	driver = backends.Redis{Path: u.Path}
	// case "postgres":
	// 	driver = backends.Postgres{Path: u.Path}
	default:
		return fmt.Errorf(`
Unrecognized scheme '%s'. You can visit https://github.com/nanopack/hoarder and
submit a pull request adding the scheme or you can submit an issue requesting its
addition.
`, u.Scheme)
	}

	return nil
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
	router.Add("HEAD", "/blobs/{blob}", handleRequest(getHead))
	router.Get("/blobs/{blob}", handleRequest(get))
	router.Get("/blobs", handleRequest(list))
	router.Post("/blobs/{blob}", handleRequest(create))
	router.Put("/blobs/{blob}", handleRequest(create))
	router.Delete("/blobs/{blob}", handleRequest(delete))

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
func writeBody(v interface{}, rw http.ResponseWriter) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// rw.Header().Set("Content-Type", "application/json")
	// rw.WriteHeader(status)
	rw.Write(b)

	return nil
}
