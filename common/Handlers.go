package common

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var People []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range People {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(People)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Id = params["id"]
	People = append(People, person)
	json.NewEncoder(w).Encode(People)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range People {
		if item.Id == params["id"] {
			People = append(People[:index], People[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(People)
}