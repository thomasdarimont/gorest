package people

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/url"
	"net/http"
)

var people []Person


func Init(){
	people = append(people, Person{ID: "1", Firstname: "Tom", Lastname: "Darimont", Address: &Address{City: "Saarbrücken", State: "Saarland"}})
	people = append(people, Person{ID: "2", Firstname: "Anne", Lastname: "Laurentius", Address: &Address{City: "Saarbrücken", State: "Saarland"}})
}

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