package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Record struct {
	Id      string `json:"id"`
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Records []Record

func main() {
	port := ":3130"
	http.HandleFunc("/", Index)
	http.HandleFunc("/records", RecordIndex)
	http.HandleFunc("/records/{recordId}", RecordShow)

	log.Fatal(http.ListenAndServe(port, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo!")
}

func RecordIndex(w http.ResponseWriter, r *http.Request) {
	records := Records{
		Record{Id: "1", Domain: "web.com", Name: "A", Address: "192.62.1.1"},
		Record{Id: "2", Domain: "web.com", Name: "CNAME", Address: "192.62.1.1"},
	}

	json.NewEncoder(w).Encode(records)
}

func RecordShow(w http.ResponseWriter, r *http.Request) {
	// not sure how to identify /records/:id yet
	fmt.Fprintln(w, "RecordShow!")
}
