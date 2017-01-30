// Package api defines the routes accessible and the logic when they are hit.
package api

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/gorilla/pat"
	"github.com/jcelliott/lumber"
	nanoauth "github.com/nanobox-io/golang-nanoauth"
	"github.com/spf13/viper"
)

// Start starts the api listener
func Start() error {
	uri, err := url.Parse(viper.GetString("listen-addr"))
	if err != nil {
		return fmt.Errorf("Failed to parse 'listen-addr' - %s", err)
	}

	// blocking...
	nanoauth.DefaultAuth.Header = "X-AUTH-TOKEN"

	// listen http (with auth support)
	if uri.Scheme == "http" {
		lumber.Info("Starting hoarder server at 'http://%s'...", uri.Host)
		return nanoauth.ListenAndServe(uri.Host, viper.GetString("token"), routes(), "/ping")
	}

	// listen https
	lumber.Info("Starting secure hoarder server at 'https://%s'...", uri.Host)
	return nanoauth.ListenAndServeTLS(uri.Host, viper.GetString("token"), routes(), "/ping")
}

// routes registers all api routes with the router
func routes() *pat.Router {
	lumber.Debug("Registering routes...\n")

	// create new router
	router := pat.New()

	// define ping
	router.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong\n"))
	})

	// blobs
	router.Add("HEAD", "/blobs/{blob}", handleRequest(getHead))
	router.Get("/blobs/{blob}", handleRequest(get))
	router.Post("/blobs/{blob}", handleRequest(create))
	router.Put("/blobs/{blob}", handleRequest(update))
	router.Delete("/blobs/{blob}", handleRequest(delete))
	router.Add("HEAD", "/blobs", handleRequest(list))
	router.Get("/blobs", handleRequest(list))

	//
	return router
}

// handleRequest is a wrapper for the actual route handler, simply to provide some
// debug output
func handleRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		//
		fn(rw, req)

		// must be after req returns
		getStatus := func(trw http.ResponseWriter) string {
			r, _ := regexp.Compile("status:([0-9]*)")
			return r.FindStringSubmatch(fmt.Sprintf("%+v", trw))[1]
		}

		getWrote := func(trw http.ResponseWriter) string {
			r, _ := regexp.Compile("written:([0-9]*)")
			return r.FindStringSubmatch(fmt.Sprintf("%+v", trw))[1]
		}

		lumber.Debug(`%s - [%s] %s %s %s(%s) - "User-Agent: %s"`,
			req.RemoteAddr, req.Proto, req.Method, req.RequestURI,
			getStatus(rw), getWrote(rw), // %s(%s)
			req.Header.Get("User-Agent"))
	}
}
