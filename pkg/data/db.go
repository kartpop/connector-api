package data

import "github.com/kartpop/connector-api/pkg/models"

type DB interface {
	GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(string) (*models.Connector, error)
	UpdateConnector(string, models.Connector) (*models.Connector, error)
	DeleteConnector(id string) error
}
