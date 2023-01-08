package data

import "github.com/sede-x/gopoc-connector/pkg/models"

type DB interface {
	GetAllConnectors() (*[]models.Connector, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(string) (*models.Connector, error)
	UpdateConnector(string, models.Connector) (*models.Connector, error)
	DeleteConnector(id string) error
	GetConnectors(locationIds []string, types []string) (*[]models.Connector, error)
}
