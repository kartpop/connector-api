package logic

import (
	"encoding/json"
	"os"

	"github.com/kartpop/connector-api/pkg/data"
	"github.com/kartpop/connector-api/pkg/models"
)

type ConnectorLogic interface {
	GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error)
	AddConnector(*models.Connector) error
	GetConnectorByID(string) (*models.Connector, error)
	UpdateConnector(id string, upcon models.Connector) (*models.Connector, error)
	DeleteConnector(id string) error
}

type Connector struct {
	DB data.DB
	MQ data.MessageQueue
}

func (c *Connector) GetConnectors(qp models.ConnectorQueryParams) (*models.ConnectorPagination, error) {
	pagedConnectors, err := c.DB.GetConnectors(qp)
	if err != nil {
		return nil, err
	}

	return pagedConnectors, nil
}

func (c *Connector) AddConnector(con *models.Connector) error {
	err := c.DB.AddConnector(con)
	if err != nil {
		return err
	}

	// Broadcast NewConnector message over message queue
	newConnectorMessage, err := json.Marshal(&models.NewConnectorMessage{
		Id:           con.Id,
		Name:         con.Name,
		LocationId:   con.LocationId,
		LocationName: con.LocationName,
	})
	if err != nil {
		return err
	}
	topic := os.Getenv("NEW_CONNECTOR_TOPIC")
	err = c.MQ.ProduceMessage(topic, newConnectorMessage)
	if err != nil {
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
