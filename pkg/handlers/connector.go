package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sede-x/gopoc-connector/pkg/models"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

// GetAllConnectors returns all connectors.
func (h Handler) GetAllConnectors(w http.ResponseWriter, r *http.Request) {
	var connectors []models.Connector

	if result := h.DB.Find(&connectors); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(connectors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConnector creates and appends a connector to the collection and sends back a json response of the created connector.
func (h Handler) AddConnector(w http.ResponseWriter, r *http.Request) {
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

	if result := h.DB.Create(&connector); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
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
func (h Handler) GetConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var connector models.Connector
	if result := h.DB.First(&connector, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
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
func (h Handler) UpdateConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
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

	var connector models.Connector
	if result := h.DB.First(&connector, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	// TODO: check if updatedConnector has all fields set.
	// If it excludes some fields, default values would get set for the original connector in DB causing data loss.
	connector.StationId = updatedConnector.StationId
	connector.Type = updatedConnector.Type
	connector.ChargeSpeed = updatedConnector.ChargeSpeed
	connector.Active = updatedConnector.Active
	if result := h.DB.Save(&connector); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
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
func (h Handler) DeleteConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var connector models.Connector
	if result := h.DB.First(&connector, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	if result := h.DB.Delete(&connector); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(fmt.Sprintf("Connector with Id: %v was deleted successfully.", id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
