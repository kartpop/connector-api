package kafkaq

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	Producer *kafka.Producer
}

func New() (*Kafka, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"client.id":         "1000",
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}

	return &Kafka{Producer: producer}, nil
}

func (k *Kafka) ProduceMessage(topic string, message []byte) error {
	delivery_chan := make(chan kafka.Event, 10000)
	err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message},
		delivery_chan,
	)
	if err != nil {
		return err
	}

	return nil
}
