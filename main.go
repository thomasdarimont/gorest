package main

import (
	"log"
	"net/http"

	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/url"
)

// Person type
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address type
type Address struct {
	City  string `json:"City,omitempty"`
	State string `json:"State,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = uuid.NewV4().String()

	people = append(people, person)

	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	location := url.URL{Scheme: req.URL.Scheme, Host: req.Host, Path: "people/" + person.ID}

	w.Header().Set("Location", location.String())
	w.WriteHeader(http.StatusCreated)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	people = append(people, Person{ID: "1", Firstname: "Tom", Lastname: "Darimont", Address: &Address{City: "Saarbrücken", State: "Saarland"}})
	people = append(people, Person{ID: "2", Firstname: "Anne", Lastname: "Laurentius", Address: &Address{City: "Saarbrücken", State: "Saarland"}})

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Println("API server ready @", *addr)

	log.Fatal(http.ListenAndServe(*addr, router))
}
