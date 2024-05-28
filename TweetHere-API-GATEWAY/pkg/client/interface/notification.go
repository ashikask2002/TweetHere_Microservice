package interfaces

import "TweetHere-API/pkg/utils/models"

type NotificationClient interface {
	GetNotification(userid int,req models.NotificationPagination) ([]models.NotificationResponse,error)
}
