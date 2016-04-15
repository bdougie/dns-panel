package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
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

var records = Records{
	Record{Id: "1", Domain: "web.com", Name: "A", Address: "192.62.1.1"},
	Record{Id: "2", Domain: "web.com", Name: "CNAME", Address: "192.62.1.1"},
}

func main() {
	port := ":3130"
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/records", RecordIndex)
	// router.GET("/records/:name", RecordShow)

	log.Fatal(http.ListenAndServe(port, nil))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func RecordIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(records)
}

func RecordShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
