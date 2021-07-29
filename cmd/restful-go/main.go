package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type QnrResponse struct {
    Id string `json:"Id"`
    Name string `json:"Name"`
    Body string `json:"Body"`
    Email string `json:"Email"`
}

var QnrResponses []QnrResponse

func addQnrResponse(w http.ResponseWriter, r *http.Request) {
    // get the body of the POST request
    // return the response body as a string
    reqBody, _ := ioutil.ReadAll(r.Body)
    var response QnrResponse
    json.Unmarshal(reqBody, &response)

    // append response to global variable
    QnrResponses = append(QnrResponses, response)

    json.NewEncoder(w).Encode(response)
}

func getQnrResponse(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop through each response, and if the ID matches, return the response
    for _, response := range QnrResponses {
        if response.Id == key {
            json.NewEncoder(w).Encode(response)
        }
    }
}

func getAllQnrResponses(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(QnrResponses)
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/responses", getAllQnrResponses).Methods("GET")

    router.HandleFunc("/responses", addQnrResponse).Methods("POST")
    router.HandleFunc("/responses/{id}", getQnrResponse)

    log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
    QnrResponses = []QnrResponse{
        {Id: "1", Name: "Mikael Rozee", Body: "I am awesome.", Email: "foo@example.com"},
        {Id: "2", Name: "Other Person", Body: "I am less awesome.", Email: "other@example.com"},
        {Id: "3", Name: "Extra Dude", Body: "I am the least awesome.", Email: "extra@example.com"},
    }
    handleRequests()
}
