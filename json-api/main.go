package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"strconv"
)

type Record struct {
	Id      string `json:"id"`
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Records []Record

var records = Records{
	Record{Id: "1", Domain: "web.com", Name: "A", Address: "192.62.1.1"},
	Record{Id: "2", Domain: "web.com", Name: "CNAME", Address: "192.62.1.1"},
}

func main() {
	port := ":3130"
	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	router.HandleFunc("/records", RecordIndex)
	router.HandleFunc("/records/{Id}", RecordShow)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(port, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func RecordIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(records)
}

func RecordShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordId := vars["Id"]

	record := filter(recordId)
	json.NewEncoder(w).Encode(record)
}

func filter(id string) Record {
	// string to int conversion, this could be better
	// Better there is probably a
	recordId, err := strconv.Atoi(id)

	if err != nil {
		panic("There is actually no record id and you should not be here")
	}

	position := recordId - 1

	return records[position]
}
