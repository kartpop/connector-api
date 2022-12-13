package data

import "github.com/sede-x/gopoc-connector/pkg/models"

type DB interface {
	GetAllConnectors() (*[]models.Connector, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(int) (*models.Connector, error)
	UpdateConnectorByID(int, models.Connector) (*models.Connector, error)
	DeleteConnectorByID(id int) error
}
