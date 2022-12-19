package data

import "github.com/sede-x/gopoc-connector/pkg/models"

type DB interface {
	GetAllConnectors() (*[]models.Connector, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(int) (*models.Connector, error)
	UpdateConnector(int, models.Connector) (*models.Connector, error)
	DeleteConnector(id int) error
}
