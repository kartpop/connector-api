package logic

import (
	"github.com/sede-x/gopoc-connector/pkg/data"
	"github.com/sede-x/gopoc-connector/pkg/models"
)

type ConnectorLogic interface {
	GetAllConnectors() (*[]models.Connector, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(int) (*models.Connector, error)
	UpdateConnector(id int, upcon models.Connector) (*models.Connector, error)
	DeleteConnector(id int) error
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

func (c *Connector) GetConnectorByID(id int) (*models.Connector, error) {
	connector, err := c.DB.GetConnectorByID(id)
	if err != nil {
		return nil, err
	}

	return connector, nil
}

func (c *Connector) UpdateConnector(id int, upcon models.Connector) (*models.Connector, error) {
	connector, err := c.DB.UpdateConnector(id, upcon)
	if err != nil {
		return nil, err
	}

	return connector, nil
}

func (c *Connector) DeleteConnector(id int) error {
	if err := c.DB.DeleteConnector(id); err != nil {
		return err
	}

	return nil
}
