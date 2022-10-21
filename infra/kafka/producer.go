package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// create new producer
func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return p
}

// publish a new message
func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, nil)

	if err != nil {
		return err
	}

	return nil
}
