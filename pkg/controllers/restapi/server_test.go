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

	"github.com/sede-x/gopoc-connector/pkg/helper"
	"github.com/sede-x/gopoc-connector/pkg/mocks"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

var connectorSet = map[string]*models.Connector{
	mocks.Connectors[0].Id: mocks.Connectors[0],
	mocks.Connectors[1].Id: mocks.Connectors[1],
	mocks.Connectors[2].Id: mocks.Connectors[2],
}

// Mock struct implementing ConnectorLogic interface - to be injected into ConnectorRestApi instance
type mockConnectorLogic struct{}

func (mcl *mockConnectorLogic) GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error) {
	// TODO: not implemented
	return &models.ConnectorPagination{
		Connectors: &mocks.Connectors,
	}, nil
}

func (mcl *mockConnectorLogic) AddConnector(c *models.Connector) error {
	c.Id = mocks.Connectors[0].Id
	return nil
}

func (mcl *mockConnectorLogic) GetConnectorByID(id string) (*models.Connector, error) {
	if connector, ok := connectorSet[id]; ok {
		return connector, nil
	}
	return nil, fmt.Errorf("record not found")
}

func (mcl *mockConnectorLogic) UpdateConnector(id string, upcon models.Connector) (*models.Connector, error) {
	connector, ok := connectorSet[id]
	if !ok {
		return nil, fmt.Errorf("record not found")
	}

	connector.LocationId = upcon.LocationId
	connector.Type = upcon.Type
	connector.ChargeSpeed = upcon.ChargeSpeed
	connector.Active = upcon.Active
	return connector, nil
}

func (mcl *mockConnectorLogic) DeleteConnector(id string) error {
	_, ok := connectorSet[id]
	if !ok {
		return fmt.Errorf("record not found")
	}
	return nil
}

var s Server

func init() {
	s = Server{ConnectorLogic: &mockConnectorLogic{}}
	s.Initialize()
}

func TestGetConnectors(t *testing.T) {
	// setup
	req, _ := http.NewRequest(http.MethodGet, "/connectors", nil)
	res := httptest.NewRecorder()

	// invoke func GetAllConnectors
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusOK, res.Code)

	var actual models.ConnectorPagination
	err := json.NewDecoder(res.Body).Decode(&actual)
	if err != nil {
		log.Fatalln(err)
	}
	expected := mocks.Connectors
	if !reflect.DeepEqual(*actual.Connectors, expected) {
		fmt.Println(actual.Connectors)
		fmt.Println(expected)
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
	locationIdMock := helper.GetMD5Hash("some-stationId")
	newcon := models.Connector{
		LocationId:  locationIdMock,
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
	if actual.Id == "" {
		t.Errorf("Expected ID of created connector %s, got empty string", actual.Id)
	}
	if actual.LocationId != locationIdMock {
		t.Errorf("Expected location ID of created connector: got -> %s, want -> %s", actual.LocationId, locationIdMock)
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
	validId := mocks.Connectors[0].Id
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/connectors/%s", validId), nil)
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
	expected := *mocks.Connectors[0]
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
	locationIdMock := helper.GetMD5Hash("some-stationId")
	upcon := models.Connector{
		LocationId:  locationIdMock,
		Type:        "L2, AC",
		ChargeSpeed: "9 kW",
		Active:      true,
	}
	b, err := json.Marshal(&upcon)
	if err != nil {
		log.Fatalln(err)
	}

	invalidId := helper.GetMD5Hash("random-id")
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/connectors/%s", invalidId), bytes.NewReader(b))
	res := httptest.NewRecorder()

	// invoke func UpdateConnectorByID
	s.router.ServeHTTP(res, req)

	// test
	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestUpdateConnectorValidRequest(t *testing.T) {
	// setup
	locationIdMock := helper.GetMD5Hash("some-stationId")
	upcon := models.Connector{
		LocationId:  locationIdMock,
		Type:        "L2, AC",
		ChargeSpeed: "9 kW",
		Active:      true,
	}
	b, err := json.Marshal(&upcon)
	if err != nil {
		log.Fatalln(err)
	}

	validId := mocks.Connectors[1].Id
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/connectors/%s", validId), bytes.NewReader(b))
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
	if actual.Id != validId {
		t.Errorf("Expected ID of updated connector %s, got %s\n", validId, actual.Id)
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
	validId := mocks.Connectors[0].Id
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/connectors/%s", validId), nil)
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
	if actual != fmt.Sprintf("Connector with Id: %v was deleted successfully.", validId) {
		t.Errorf("Expected delete message did not match actual.")
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d\n", expected, actual)
	}
}
