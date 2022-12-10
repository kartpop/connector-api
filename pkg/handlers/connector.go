package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sede-x/gopoc-connector/pkg/mocks"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

// GetAllConnectors returns all connectors.
func GetAllConnectors(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(mocks.Connectors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConnector creates and appends a connector to the collection and sends back a json response of the created connector.
func AddConnector(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var connector models.Connector
	err = json.Unmarshal(body, &connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Append to Connectors mock list for now
	connector.Id = uuid.New()
	mocks.Connectors = append(mocks.Connectors, connector)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetConnectorById returns the connector for the given ID.
func GetConnectorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Loop through Connectors mock list for now
	for _, connector := range mocks.Connectors {
		if connector.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(connector)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// Connector not found!
	http.Error(w, fmt.Sprintf("There is no connector with Id: %v. Please provide a valid connector Id.", id), http.StatusNotFound)
}

// UpdateConnectorById updates the connector for given Id and json body in the request.
func UpdateConnectorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedConnector models.Connector
	err = json.Unmarshal(body, &updatedConnector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Loop through Connectors mock list for now
	for index, connector := range mocks.Connectors {
		if connector.Id == id {
			connector.StationId = updatedConnector.StationId
			connector.Type = updatedConnector.Type
			connector.ChargeSpeed = updatedConnector.ChargeSpeed
			connector.Active = updatedConnector.Active
			mocks.Connectors[index] = connector

			w.Header().Add("Context-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(connector)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// Connector not found!
	http.Error(w, fmt.Sprintf("There is no connector with Id: %v. Please provide a valid connector Id.", id), http.StatusNotFound)
}

// DeleteConnectorById deletes the connector for given Id.
func DeleteConnectorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Loop through Connectors mock list for now
	for index, connector := range mocks.Connectors {
		if connector.Id == id {
			mocks.Connectors = append(mocks.Connectors[:index], mocks.Connectors[index+1:]...)

			w.Header().Add("Context-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(fmt.Sprintf("Connector with Id: %v was deleted successfully.", id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// Connector not found!
	http.Error(w, fmt.Sprintf("There is no connector with Id: %v. Please provide a valid connector Id.", id), http.StatusNotFound)
}
