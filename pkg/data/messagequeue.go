package data

type MessageQueue interface {
	ProduceMessage(topic string, message []byte) error
}
