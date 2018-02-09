package people

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/url"
	"net/http"
)

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(personService.FindAll())
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	person, e := personService.FindByID(params["id"])
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)

	newPerson := personService.CreatePerson(person)

	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	location := url.URL{Scheme: req.URL.Scheme, Host: req.Host, Path: "people/" + newPerson.ID}

	w.Header().Set("Location", location.String())
	w.WriteHeader(http.StatusCreated)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	deleted := personService.DeletePerson(params["id"])

	if deleted {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
