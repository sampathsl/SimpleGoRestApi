package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	common "github.com/sampathsl/SimpleGoRestApi/common"
)

func main() {

	router := mux.NewRouter()

	common.People = append(common.People, common.Person{Id: "1", FirstName: "Sampath", LastName: "Test",
		Address: &common.Address{Address1: "Main Road", Address2: "No f boss",
			Address3: "Freedom Work", City: "Dublin", State: "CA", Postcode:"20000"}})
	common.People = append(common.People, common.Person{Id: "2", FirstName: "Ravi", LastName: "Test"})
	common.People = append(common.People, common.Person{Id: "3", FirstName: "Todo", LastName: "Test",
		Address: &common.Address{Address1: "Flower Road", Address2: "Upper",
			Address3: "block c", City: "Dublin", State: "CA", Postcode:"55VDVS"}})

	router.HandleFunc("/people", common.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", common.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", common.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", common.DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))

}
