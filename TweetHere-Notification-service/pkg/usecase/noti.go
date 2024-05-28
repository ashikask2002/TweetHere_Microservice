package usecase

import (
	"encoding/json"
	"fmt"
	"time"
	authface "tweethere-Notification/pkg/client/interface"
	"tweethere-Notification/pkg/config"
	interfaces "tweethere-Notification/pkg/repository/interface"
	services "tweethere-Notification/pkg/usecase/interface"
	"tweethere-Notification/pkg/utils/models"

	"github.com/IBM/sarama"
)

type notiUsecase struct {
	notiRepository interfaces.NotiRepository
	authclient     authface.Newauthclient
}

func NewnotiUsecase(repository interfaces.NotiRepository, authface authface.Newauthclient) services.NotiUseCase {
	return &notiUsecase{
		notiRepository: repository,
		authclient:     authface,
	}
}

func (c *notiUsecase) ConsumeNotification() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error in load config")
	}

	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("error creating kafka consumer", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaTpic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error creating partition consumer", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("kafka consumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelNotification(message.Value)
			if err != nil {
				fmt.Println("error unmarshelling message", err)
				continue
			}
			fmt.Println("received message", msg)
			err = c.notiRepository.StorenotificationReq(*msg)

			if err != nil {
				fmt.Println("error storing notification in repository", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("kafka cosumer error", err)

		}
	}
}

func (c *notiUsecase) UnmarshelNotification(data []byte) (*models.NotificationReq, error) {
	var notification models.NotificationReq
	err := json.Unmarshal(data, &notification)
	if err != nil {
		return nil, err
	}
	notification.CreatedAt = time.Now()

	return &notification, nil
}

func (c *notiUsecase) GetNotification(userid int, mod models.Pagination) ([]models.NotificationResponse, error) {
	data, err := c.notiRepository.GetNotification(userid, mod)
	if err != nil {
		return []models.NotificationResponse{}, err
	}
	var response []models.NotificationResponse
	fmt.Println("dddddddd", data)
	for _, v := range data {
		userdata, err := c.authclient.UserData(v.SenderID)
		if err != nil {
			fmt.Println("heloooooooooooo")
			return nil, err
		}
		response = append(response, models.NotificationResponse{
			UserID:    int(userdata.UserId),
			Username:  userdata.Username,
			Profile:   userdata.Profile,
			Message:   v.Message,
			CreatedAt: v.CreatedAt.String(),
		})
	}
	return response, nil
}
