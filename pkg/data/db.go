package data

import "github.com/sede-x/gopoc-connector/pkg/models"

type DB interface {
	GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(string) (*models.Connector, error)
	UpdateConnector(string, models.Connector) (*models.Connector, error)
	DeleteConnector(id string) error
}
