package main

import (
	"fmt"
	"log"

	kafkaApp "github.com/dan36poars/simulator-gps/app/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/dan36poars/simulator-gps/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)

	consumer := kafka.NewKafkaComsumer(msgChan)

	go consumer.Consume() // asynchronous thread

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApp.Produce(msg)
	}

}
