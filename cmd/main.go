package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sede-x/gopoc-connector/pkg/db"
	"github.com/sede-x/gopoc-connector/pkg/handlers"
)

func main() {
	// TODO: check if it is OK to use godotenv package to handle environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}
	dbURL := os.Getenv("DBURL")

	gormdb, err := db.Init(dbURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	h := handlers.New(gormdb)
	router := mux.NewRouter()

	router.HandleFunc("/connectors", h.GetAllConnectors).Methods(http.MethodGet)
	router.HandleFunc("/connectors", h.AddConnector).Methods(http.MethodPost)
	router.HandleFunc("/connectors/{id}", h.GetConnectorByID).Methods(http.MethodGet)
	router.HandleFunc("/connectors/{id}", h.UpdateConnectorByID).Methods(http.MethodPut)
	router.HandleFunc("/connectors/{id}", h.DeleteConnectorByID).Methods(http.MethodDelete)

	log.Println("Connector API is running")
	http.ListenAndServe(":4000", router)
}
