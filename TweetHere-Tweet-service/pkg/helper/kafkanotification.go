package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"tweet-service/pkg/config"
	"tweet-service/pkg/utils/models"

	"github.com/IBM/sarama"
)

func SendNotification(data models.Notification, msg []byte) {
	data.Message = string(msg)
	err := KafkanotificationProducer(data)
	if err != nil{
		fmt.Println("error sending notification to kafka",err)
		return
	}
	fmt.Println("==sent like successsfully to kafka==")
}

func KafkanotificationProducer(message models.Notification) error {
	cfg, _ := config.LoadConfig()
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.RequiredAcks = sarama.WaitForAll
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		return err
	}
	result, errs := json.Marshal(message)
	if errs != nil {
		return errs
	}

	msg := &sarama.ProducerMessage{Topic: cfg.KafkaTpic, Key: sarama.StringEncoder("Notifications"), Value: sarama.StringEncoder(result)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("error send message in kafka", err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}
