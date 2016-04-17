package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
	"strconv"
)

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
	recordID := vars["ID"]

	record := filter(recordID)
	json.NewEncoder(w).Encode(record)
}

func filter(id string) Record {
	recordID, err := strconv.Atoi(id)

	if err != nil {
		panic("There is actually no record id and you should not be here")
	}

	position := recordID - 1

	return records[position]
}

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
