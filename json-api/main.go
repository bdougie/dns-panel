package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3130"
	http.HandleFunc("/", Index)
	http.HandleFunc("/records", RecordIndex)

	log.Fatal(http.ListenAndServe(port, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func RecordIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "RecordIndex!")
}
