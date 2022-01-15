package kafka_service

import (
	"fmt"
	"github.com/Shopify/sarama"
	utility "mariadbapp/utilities"
)

func InitProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	prd, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	return prd, err
}

func Publish(message string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: utility.Topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("error publish", err.Error())
	}
}
