package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaComsumer struct {
	MsgChan chan *ckafka.Message
}

// Fabric Pattern
func NewKafkaComsumer(msgChan chan *ckafka.Message) *KafkaComsumer {
	return &KafkaComsumer{
		MsgChan: msgChan,
	}
}

// kafka message consuming
func (k *KafkaComsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		log.Fatal("error consuming kafka message" + err.Error())
	}

	// what topics consuming
	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started ...")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
