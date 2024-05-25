package interfaces

import "tweethere-chat/pkg/utils/models"

type ChatRepository interface {
	StoreFriendsChat(models.MessageReq) error
	UpdateReadAsMessage(string, string) error
	GetFriendChat(string,string,models.Pagination) ([]models.Message,error)
}
