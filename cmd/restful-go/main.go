package main

import (
	"encoding/json"
	"fmt"
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

func getQnrResponse(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop through each response, and if the ID matches, return the response
    for _, response := range QnrResponses {
        if response.Id == key {
            json.NewEncoder(w).Encode(response)
        }
    }
}

func getAllQnrResponses(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint hit: all qnr responses")
    json.NewEncoder(w).Encode(QnrResponses)
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/responses", getAllQnrResponses)
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
