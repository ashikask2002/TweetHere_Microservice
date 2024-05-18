package interfaces

import "tweet-service/pkg/utils/models"


type AuthClient interface{
	DoesUserExist(id int64)(bool,error)
	FindUserName(id int64)(string,error)
	UserData(userid int)(models.UserData,error)
}