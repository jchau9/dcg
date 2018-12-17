package healthcheck

import (
	"go_development/src/response"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck ...
func HealthCheck(router *mux.Router) {
	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		response.Respond(w, r, "Hello, World!")
	}).Methods("GET")

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		response.Respond(w, r, string(body[:]))
	}).Methods("POST")
}
