package helper

import (
	"TweetHere-API/pkg/config"
	"TweetHere-API/pkg/utils/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

func SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	fmt.Println("poooooooooo")
	fmt.Println("1111111")
	var message models.Message
	fmt.Println("222222")
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		fmt.Println("error while unmarshel", err)
	}
	fmt.Println("333333")

	message.SenderID = userID
	fmt.Println("444444")
	recipientConn, ok := User[message.RecipientID]
	fmt.Println("55555")
	if ok {
		recipientConn.WriteMessage(websocket.TextMessage, msg)
	}
	err := KafkaProducer(message)
	fmt.Println("==send successfully==", err)

}

func KafkaProducer(message models.Message) error {
	fmt.Println("shoooooo")
	fmt.Println("from kafka", message)
	cfg, _ := config.LoadConfig()
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		return err
	}

	result, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{Topic: cfg.KafkaTpic, Key: sarama.StringEncoder("Friend message"), Value: sarama.StringEncoder(result)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("err send message in kafka", err)
	}
	log.Printf("[producer] partition id:%d; offset:%d, value:%v\n", partition, offset, message)
	return nil

}
