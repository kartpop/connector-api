package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/sede-x/gopoc-connector/pkg/mocks"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

// Mock struct implementing ConnectorLogic interface - to be injected into ConnectorRestApi instance
type mockConnectorLogic struct{}

func (mcl *mockConnectorLogic) GetAllConnectors() (*[]models.Connector, error) {
	return &mocks.Connectors, nil
}

func (mcl *mockConnectorLogic) AddConnector(c *models.Connector) error {
	c.Id = 6
	return nil
}

func (mcl *mockConnectorLogic) GetConnectorByID(id int) (*models.Connector, error) {
	if id >= 1 && id <= 3 {
		return &mocks.Connectors[id-1], nil // mocks.Connectors have ID's 1,2,3
	}
	return nil, fmt.Errorf("record not found")
}

func (mcl *mockConnectorLogic) UpdateConnector(id int, upcon models.Connector) (*models.Connector, error) {
	if id < 1 || id > 3 {
		return nil, fmt.Errorf("record not found") // mocks.Connectors have ID's 1,2,3
	}
	connector := models.Connector{}
	connector.Id = id
	connector.StationId = upcon.StationId
	connector.Type = upcon.Type
	connector.ChargeSpeed = upcon.ChargeSpeed
	connector.Active = upcon.Active
	return &connector, nil
}

func (mcl *mockConnectorLogic) DeleteConnector(id int) error {
	if id < 1 || id > 3 {
		return fmt.Errorf("record not found") // mocks.Connectors have ID's 1,2,3
	}
	return nil
}

var s Server

func init() {
	s = Server{ConnectorLogic: &mockConnectorLogic{}}
	s.Initialize()
}

func TestGetAllConnectors(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodGet, "/connectors", nil)
	res := httptest.NewRecorder()

	// invoke func GetAllConnectors
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusOK, res.Code)

	var actual []models.Connector
	err := json.NewDecoder(res.Body).Decode(&actual)
	if err != nil {
		log.Fatalln(err)
	}
	expected := mocks.Connectors
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual response from GET /connectors does not match expected expected response.")
	}
}

func TestAddConnectorBadRequest(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodPost, "/connectors", strings.NewReader("Not a valid models.Connector json - illformed request!"))
	res := httptest.NewRecorder()

	// invoke func AddConnector
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusBadRequest, res.Code)
}

func TestAddConnectorValidRequest(t *testing.T) {
	// setup
	newcon := models.Connector{
		StationId:   1,
		Type:        "L1, AC",
		ChargeSpeed: "3 kW",
		Active:      true,
	}
	b, err := json.Marshal(&newcon)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/connectors", bytes.NewReader(b))
	res := httptest.NewRecorder()

	// invoke func AddConnector
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusCreated, res.Code)

	var actual models.Connector
	err = json.NewDecoder(res.Body).Decode(&actual)
	if err != nil {
		log.Fatalln(err)
	}
	if actual.Id != 6 {
		t.Errorf("Expected ID of created connector %d, got %d\n", 6, actual.Id)
	}
}

func TestGetConnectorByIDInvalidID(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodGet, "/connectors/6", nil)
	res := httptest.NewRecorder()

	// invoke func GetConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestGetConnectorByIDValidRequest(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodGet, "/connectors/3", nil)
	res := httptest.NewRecorder()

	// invoke func GetConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusOK, res.Code)

	var actual models.Connector
	err := json.NewDecoder(res.Body).Decode(&actual)
	if err != nil {
		log.Fatalln(err)
	}
	expected := mocks.Connectors[2] // Connector ID 3 is at index 2 in mocks
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual response from GET /connectors does not match expected expected response.")
	}
}

func TestUpdateConnectorBadRequest(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodPut, "/connectors/3", strings.NewReader("Not a valid models.Connector json - illformed request!"))
	res := httptest.NewRecorder()

	// invoke func UpdateConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusBadRequest, res.Code)
}

func TestUpdateConnectorInvalidID(t *testing.T) {
	// setup
	upcon := models.Connector{
		StationId:   15,
		Type:        "L2, AC",
		ChargeSpeed: "9 kW",
		Active:      true,
	}
	b, err := json.Marshal(&upcon)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest(http.MethodPut, "/connectors/6", bytes.NewReader(b))
	res := httptest.NewRecorder()

	// invoke func UpdateConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestUpdateConnectorValidRequest(t *testing.T) {
	// setup
	upcon := models.Connector{
		StationId:   15,
		Type:        "L2, AC",
		ChargeSpeed: "9 kW",
		Active:      true,
	}
	b, err := json.Marshal(&upcon)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest(http.MethodPut, "/connectors/3", bytes.NewReader(b))
	res := httptest.NewRecorder()

	// invoke func UpdateConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusOK, res.Code)

	var actual models.Connector
	err = json.NewDecoder(res.Body).Decode(&actual)
	if err != nil {
		log.Fatalln(err)
	}
	if actual.Id != 3 {
		t.Errorf("Expected ID of updated connector %d, got %d\n", 3, actual.Id)
	}
}

func TestDeleteConnectorInvalidID(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodDelete, "/connectors/6", nil)
	res := httptest.NewRecorder()

	// invoke func DeleteConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestDeleteConnectorValidRequest(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodDelete, "/connectors/3", nil)
	res := httptest.NewRecorder()

	// invoke func DeleteConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusOK, res.Code)

	var actual string
	err := json.NewDecoder(res.Body).Decode(&actual)
	if err != nil {
		log.Fatalln(err)
	}
	if actual != fmt.Sprintf("Connector with Id: %v was deleted successfully.", 3) {
		t.Errorf("Expected delete message did not match actual.")
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d\n", expected, actual)
	}
}
