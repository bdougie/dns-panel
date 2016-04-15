package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}

func Index(w http.ResponseWriter, r *http.Request) {
	setCors(w)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func RecordIndex(w http.ResponseWriter, r *http.Request) {
	setCors(w)
	json.NewEncoder(w).Encode(records)
}

func RecordShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordId := vars["Id"]

	record := filter(recordId)
	json.NewEncoder(w).Encode(record)
}

func filter(id string) Record {
	recordId, err := strconv.Atoi(id)

	if err != nil {
		panic("There is actually no record id and you should not be here")
	}

	position := recordId - 1

	return records[position]
}

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
