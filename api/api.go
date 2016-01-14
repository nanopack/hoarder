package api

import (
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

// utilized by the various backends
type (
	Driver interface {
		Init() error
		List() ([]backends.FileInfo, error)
		Read(string) (io.Reader, error)
		Remove(string) error
		Stat(string) (backends.FileInfo, error)
		Write(string, io.Reader) error
	}
)

//
var driver Driver

// Start the api
func Start() error {

	// set, and initialize, the backend driver
	if err := setDriver(); err != nil {
		config.Log.Fatal(err.Error())
		os.Exit(1)
	}

	// start garbage collector
	if config.GarbageCollect {
		config.Log.Debug("Starting garbage collector (data older than %ds)...", config.CleanAfter.Value)
		go startCollection()
	}

	// blocking...
	return nanoauth.ListenAndServeTLS(config.Addr, config.Token, routes())
}

//
func setDriver() error {

	// parse connection string
	u, err := url.Parse(config.Connection)
	if err != nil {
		return err
	}

	// set backend based on connection string's scheme
	switch u.Scheme {
	case "file":
		driver = &backends.Filesystem{Path: u.Path}
	// case "scribble":
	// 	driver = backends.Scribble{Path: u.Path}
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

	// initialize the driver
	return driver.Init()
}

// routes registers all api routes with the router
func routes() *pat.Router {
	config.Log.Debug("Registering routes...\n")

	//
	router := pat.New()

	//
	router.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong\n"))
	})

	// blobs
	router.Add("HEAD", "/blobs/{blob}", handleRequest(getHead))
	router.Get("/blobs/{blob}", handleRequest(get))
	router.Get("/blobs", handleRequest(list))
	router.Add("HEAD", "/blobs", handleRequest(list))
	router.Post("/blobs/{blob}", handleRequest(create))
	router.Put("/blobs/{blob}", handleRequest(create))
	router.Delete("/blobs/{blob}", handleRequest(delete))

	return router
}

// handleRequest is a wrapper for the actual route handler, simply to provide some
// debug output
func handleRequest(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		fn(rw, req)

		// must be after fn if ever going to get rw.status (logging still more meaningful)
		config.Log.Trace(`%v - [%v] %v %v %v(%s) - "User-Agent: %s", "X-Nanobox-Token: %s"`,
			req.RemoteAddr, req.Proto, req.Method, req.RequestURI,
			rw.Header().Get("status"), req.Header.Get("Content-Length"),
			req.Header.Get("User-Agent"), req.Header.Get("X-Nanobox-Token"))
	}
}
