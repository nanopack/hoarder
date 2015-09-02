package main

import(
	"github.com/gorilla/pat"
	"net/http"
	"encoding/json"
	"github.com/jcelliott/lumber"
	"time"
	"bitbucket.org/nanobox/nanoauth"
)


type object struct {
	Name     string
	CheckSum string
	ModTime  time.Time
	Size     int64
}

// Start
func Start(addr, token string) error {
	return nanoauth.ListenAndServeTLS(addr, token, routes())
}

func routes() *pat.Router {
	router := pat.New()

	// builds
	router.Get("/builds/{file}", handleRequest(getBuild))
	router.Add("HEAD", "/builds/{file}", handleRequest(getBuildHead))
	router.Post("/builds/{file}", handleRequest(createBuild))
	router.Put("/builds/{file}", handleRequest(createBuild))
	router.Delete("/builds/{file}", handleRequest(deleteBuild))
	router.Get("/builds", handleRequest(listBuilds))

	// libs
	router.Get("/libs", handleRequest(getLibs))
	router.Put("/libs", handleRequest(createLibs))
	router.Post("/libs", handleRequest(createLibs))

	router.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Γειά σου Κόσμε"))
	})
	return router
}

func handleRequest(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		lumber.Debug(`
Request:
--------------------------------------------------------------------------------
%#v

`, req)

		//
		fn(rw, req)
		lumber.Debug(`
Response:
--------------------------------------------------------------------------------
%#v

`, rw)
	}
}

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
