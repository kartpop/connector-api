package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sede-x/gopoc-connector/pkg/logic"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

type Server struct {
	logic.ConnectorLogic
	router *mux.Router
}

func (s *Server) Start(serveraddr string) {
	s.Initialize()

	// start server
	log.Printf("Connector REST API is running at http://localhost%s/", serveraddr)
	http.ListenAndServe(serveraddr, s.router)
}

func (s *Server) Initialize() {
	s.router = mux.NewRouter()

	// CRUD
	s.router.HandleFunc("/connectors", s.AddConnector).Methods(http.MethodPost)
	s.router.HandleFunc("/connectors/{id}", s.GetConnectorByID).Methods(http.MethodGet)
	s.router.HandleFunc("/connectors/{id}", s.UpdateConnector).Methods(http.MethodPut)
	s.router.HandleFunc("/connectors/{id}", s.DeleteConnector).Methods(http.MethodDelete)

	// Filter - POST request with query params in body
	s.router.HandleFunc("/connectors/filter", s.FilterConnectors).Methods(http.MethodPost)

	// [Redundant - to be removed] - GET request with query params in URL
	s.router.HandleFunc("/connectors", s.GetConnectors).Methods(http.MethodGet)
}

// FilterConnectors returns a list of connectors based on the query paramerters sent in the body of the 
// POST request. If there are no query parameters, all connectors will be returned.
func (s *Server) FilterConnectors(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // or http.StatusInternalServerError ??
		return
	}

	var queryParams models.ConnectorQueryParams
	err = json.Unmarshal(body, &queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var pagedConnectors *models.ConnectorPagination
	pagedConnectors, err = s.ConnectorLogic.GetConnectors(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(pagedConnectors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


// GetConnectors returns a list of connectors based on the query paramerters. If there are no
// query parameters, all connectors will be returned.
func (s *Server) GetConnectors(w http.ResponseWriter, r *http.Request) {
	queryParams, err := ValidateAndGetQueryParams(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var pagedConnectors *models.ConnectorPagination
	pagedConnectors, err = s.ConnectorLogic.GetConnectors(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(pagedConnectors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConnector creates and appends a connector to the collection and sends back a json response of the created connector.
func (s *Server) AddConnector(w http.ResponseWriter, r *http.Request) {
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

	err = s.ConnectorLogic.AddConnector(&connector)
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
func (s *Server) GetConnectorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	connector, err := s.ConnectorLogic.GetConnectorByID(id)
	if err != nil {
		// TODO: must be a better way of doing this!
		// can do `errors.Is(err, gorm.ErrRecordNotFound)` - but this introduces
		// dependency of controller layer on DB ORM library
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

// UpdateConnector updates the connector for given Id and json body in the request.
func (s *Server) UpdateConnector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

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

	connector, err := s.ConnectorLogic.UpdateConnector(id, updatedConnector)
	if err != nil {
		// TODO: check how `errors.Is(err, gorm.ErrRecordNotFound)` can be incorporated
		// without indroducing dependency on ORM library
		if err.Error() == "record not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(connector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteConnector deletes the connector for given Id.
func (s *Server) DeleteConnector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := s.ConnectorLogic.DeleteConnector(id)
	if err != nil {
		// TODO: check how `errors.Is(err, gorm.ErrRecordNotFound)` can be incorporated
		// without indroducing dependency on ORM library
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
