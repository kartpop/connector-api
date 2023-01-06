package logic

import (
	"github.com/sede-x/gopoc-connector/pkg/data"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

type ConnectorLogic interface {
	GetAllConnectors() (*[]models.Connector, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(string) (*models.Connector, error)
	UpdateConnector(id string, upcon models.Connector) (*models.Connector, error)
	DeleteConnector(id string) error
}

type Connector struct {
	data.DB
}

func (c *Connector) GetAllConnectors() (*[]models.Connector, error) {
	connectors, err := c.DB.GetAllConnectors()
	if err != nil {
		return nil, err
	}

	return connectors, nil
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
