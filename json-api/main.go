package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Record struct {
	id      string `json:"id"`
	domain  string `json:"domain"`
	name    string `json:"name"`
	address string `json:"address"`
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
		Record{id: "1", domain: "web.com", name: "A", address: "192.62.1.1"},
		Record{id: "2", domain: "web.com", name: "CNAME", address: "192.62.1.1"},
	}

	json.NewEncoder(w).Encode(records)
}

func RecordShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "RecordShow!")
}
