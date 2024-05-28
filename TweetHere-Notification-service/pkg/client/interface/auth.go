package interfaces

import "tweethere-Notification/pkg/utils/models"


type Newauthclient interface{
	UserData(userid int)(models.UserData,error)
}