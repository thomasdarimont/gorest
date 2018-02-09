package main

import (
	"log"
	"net/http"
	"flag"
	"github.com/gorilla/mux"
	"github.com/thomasdarimont/gotraining/gorest/people"
	"github.com/thomasdarimont/gotraining/gorest/actuator"
)

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/people", people.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", people.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", people.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", people.DeletePersonEndpoint).Methods("DELETE")

	router.HandleFunc("/application/health", actuator.HealthEndpoint).Methods("GET")
	router.HandleFunc("/application/info", actuator.InfoEndpoint).Methods("GET")

	log.Println("API server ready @", *addr)

	log.Fatal(http.ListenAndServe(*addr, router))
}
