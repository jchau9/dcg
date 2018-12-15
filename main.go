package main

import (
	"go_development/src/api/healthcheck"
	"go_development/src/api/v1"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":3000"
	router := mux.NewRouter()
	router.StrictSlash(true)
	healthcheck.HealthCheck(router)
	version1.Version1(router)
	/*
		fmt.Println("testing base connector...")
		fmt.Println(tests.TestBaseConnector())
		fmt.Println("testing vdms connector...")
		fmt.Println(tests.TestVDMSConnector())
	*/
	log.Fatal(http.ListenAndServe(port, router))
}
