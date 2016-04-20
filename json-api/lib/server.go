package lib

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func StartSever() {
	port := ":3130"
	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	router.HandleFunc("/records", RecordIndex)
	router.HandleFunc("/records/{ID}", RecordShow)
	http.Handle("/", router)

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}
