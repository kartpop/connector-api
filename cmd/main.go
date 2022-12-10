package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sede-x/gopoc-connector/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/connectors", handlers.GetAllConnectors).Methods(http.MethodGet)
	router.HandleFunc("/connectors", handlers.AddConnector).Methods(http.MethodPost)
	router.HandleFunc("/connectors/{id}", handlers.GetConnectorById).Methods(http.MethodGet)
	router.HandleFunc("/connectors/{id}", handlers.UpdateConnectorById).Methods(http.MethodPut)
	router.HandleFunc("/connectors/{id}", handlers.DeleteConnectorById).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
