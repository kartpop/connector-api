package logic

import (
	"github.com/sede-x/gopoc-connector/pkg/data"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

type ConnectorLogic interface {
	GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(string) (*models.Connector, error)
	UpdateConnector(id string, upcon models.Connector) (*models.Connector, error)
	DeleteConnector(id string) error
}

type Connector struct {
	data.DB
}

func (c *Connector) GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error) {
	pagedConnectors, err := c.DB.GetConnectors(qp)
	if err != nil {
		return nil, err
	}

	return pagedConnectors, nil
}

func (c *Connector) AddConnector(con *models.Connector) error {
	if err := c.DB.AddConnector(con); err != nil {
		return err
	}

	return nil
}

func (c *Connector) GetConnectorByID(id string) (*models.Connector, error) {
	connector, err := c.DB.GetConnectorByID(id)
	if err != nil {
		return nil, err
	}

	return connector, nil
}

func (c *Connector) UpdateConnector(id string, upcon models.Connector) (*models.Connector, error) {
	connector, err := c.DB.UpdateConnector(id, upcon)
	if err != nil {
		return nil, err
	}

	return connector, nil
}

func (c *Connector) DeleteConnector(id string) error {
	if err := c.DB.DeleteConnector(id); err != nil {
		return err
	}

	return nil
}
