package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type QnrResponse struct {
    Name string `json:"Name"`
    Body string `json:"Body"`
    Email string `json:"Email"`
}

var QnrResponses []QnrResponse

func getAllQnrResponses(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint hit: all qnr responses")
    json.NewEncoder(w).Encode(QnrResponses)
}

func handleRequests() {
    http.HandleFunc("/responses", getAllQnrResponses)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    QnrResponses = []QnrResponse{
        {Name: "Mikael Rozee", Body: "I am awesome.", Email: "foo@example.com"},
        {Name: "Other Person", Body: "I am less awesome.", Email: "other@example.com"},
        {Name: "Extra Dude", Body: "I am the least awesome.", Email: "extra@example.com"},
    }
    handleRequests()
}
