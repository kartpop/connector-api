package logic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	KafkaProducer *kafka.Producer
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

	// Broadcast connector created message over Kafka
	jsonPayload, err := json.Marshal(&models.NewConnectorMessage{
		Id:           con.Id,
		Name:         con.Name,
		LocationId:   con.LocationId,
		LocationName: con.LocationName,
	})
	if err != nil {
		return err
	}
	kafkaTopic := os.Getenv("KAFKA_TOPIC_CONNECTOR")
	delivery_chan := make(chan kafka.Event, 10000)
	err = c.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kafkaTopic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(jsonPayload)},
		delivery_chan,
	)
	if err != nil {
		fmt.Println("Error producing Kafka message: " + err.Error())
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
