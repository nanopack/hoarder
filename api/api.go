package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/pat"
	"github.com/jcelliott/lumber"
	nanoauth "github.com/nanobox-io/golang-nanoauth"
	"github.com/spf13/viper"

	"github.com/nanopack/hoarder/backends"
	"github.com/nanopack/hoarder/util"
)

// utilized by the various backends
type (

	// Driver ...
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

// Start ...
func Start() error {

	// set, and initialize, the backend driver
	if err := setDriver(); err != nil {
		lumber.Error(err.Error())
		os.Exit(1)
	}

	// blocking...

	switch viper.GetBool("insecure") {
	case false:
		lumber.Info("Starting secure hoarder server at '%s'...\n", util.GetURI())
		nanoauth.DefaultAuth.Header = "X-AUTH-TOKEN"
		return nanoauth.ListenAndServeTLS(util.GetURI(), viper.GetString("token"), routes())
	default:
		lumber.Info("Starting hoarder server at '%s'...\n", util.GetURI())
		return http.ListenAndServe(util.GetURI(), routes())
	}
}

// setDriver
func setDriver() error {

	// parse connection string
	u, err := url.Parse(viper.GetString("backend"))
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
	lumber.Debug("Registering routes...\n")

	//
	router := pat.New()

	//
	router.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong\n"))
	})

	// blobs
	router.Add("HEAD", "/blobs/{blob}", authenticate(handleRequest(getHead)))
	router.Get("/blobs/{blob}", authenticate(handleRequest(get)))
	router.Post("/blobs/{blob}", authenticate(handleRequest(create)))
	router.Put("/blobs/{blob}", authenticate(handleRequest(update)))
	router.Delete("/blobs/{blob}", authenticate(handleRequest(delete)))

	router.Add("HEAD", "/blobs", authenticate(handleRequest(list))) // needs to be after get
	router.Get("/blobs", authenticate(handleRequest(list)))         // needs to be after get

	//
	return router
}

// authenticate veerifies that the token is allowed throught the authenticator
func authenticate(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		// if a token was provided at startup then its presumed authentication is
		// desired
		if viper.GetString("token") != "" {

			// check to see if the token was passed either in the header or as a query
			// param
			var xtoken string
			switch {
			case req.Header.Get("X-AUTH-TOKEN") != "":
				xtoken = req.Header.Get("X-AUTH-TOKEN")
			case req.FormValue("X-AUTH-TOKEN") != "":
				xtoken = req.FormValue("X-AUTH-TOKEN")
			}

			// if the tokens don't match then the connection is unauthorized
			if xtoken != viper.GetString("token") {
				rw.WriteHeader(401)
				return
			}
		}

		//
		fn(rw, req)
	}
}

// handleRequest is a wrapper for the actual route handler, simply to provide some
// debug output
func handleRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		//
		fn(rw, req)

		// must be after fn if ever going to get rw.status (logging still more meaningful)
		lumber.Debug(`%v - [%v] %v %v %v(%s) - "User-Agent: %s", "X-AUTH-TOKEN: %s\n"`,
			req.RemoteAddr, req.Proto, req.Method, req.RequestURI,
			rw.Header().Get("status"), req.Header.Get("Content-Length"),
			req.Header.Get("User-Agent"), req.Header.Get("X-AUTH-TOKEN"))
	}
}
