package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sede-x/gopoc-connector/pkg/logic"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

type ConnectorRestAPI struct {
	logic.ConnectorLogic
}

func (cra *ConnectorRestAPI) StartServer(serveraddr string) {
	router := mux.NewRouter()
	router.HandleFunc("/connectors", cra.GetAllConnectors).Methods(http.MethodGet)
	router.HandleFunc("/connectors", cra.AddConnector).Methods(http.MethodPost)
	router.HandleFunc("/connectors/{id}", cra.GetConnectorByID).Methods(http.MethodGet)
	router.HandleFunc("/connectors/{id}", cra.UpdateConnectorByID).Methods(http.MethodPut)
	router.HandleFunc("/connectors/{id}", cra.DeleteConnectorByID).Methods(http.MethodDelete)

	// start server
	log.Println("Connector API is running")
	http.ListenAndServe(serveraddr, router)
}

// GetAllConnectors returns all connectors.
func (cra *ConnectorRestAPI) GetAllConnectors(w http.ResponseWriter, r *http.Request) {
	connectors, err := cra.ConnectorLogic.GetAllConnectors()

	// TODO: handle http status codes
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(connectors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConnector creates and appends a connector to the collection and sends back a json response of the created connector.
func (cra *ConnectorRestAPI) AddConnector(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // or http.StatusInternalServerError ??
		return
	}

	var connector models.Connector
	err = json.Unmarshal(body, &connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = cra.ConnectorLogic.AddConnector(&connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetConnectorByID returns the connector for the given ID.
func (cra *ConnectorRestAPI) GetConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connector, err := cra.ConnectorLogic.GetConnectorByID(id)
	if err != nil {
		// TODO: must be a better way of doing this!
		if err.Error() == "record not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateConnectorByID updates the connector for given Id and json body in the request.
func (cra *ConnectorRestAPI) UpdateConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // or http.StatusInternalServerError ??
		return
	}

	var updatedConnector models.Connector
	err = json.Unmarshal(body, &updatedConnector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connector, err := cra.ConnectorLogic.UpdateConnectorByID(id, updatedConnector)
	if err != nil {
		// TODO: must be a better way of doing this!
		if err.Error() == "record not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteConnectorByID deletes the connector for given Id.
func (cra *ConnectorRestAPI) DeleteConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = cra.ConnectorLogic.DeleteConnectorByID(id)
	if err != nil {
		// TODO: must be a better way of doing this!
		if err.Error() == "record not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(fmt.Sprintf("Connector with Id: %v was deleted successfully.", id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
