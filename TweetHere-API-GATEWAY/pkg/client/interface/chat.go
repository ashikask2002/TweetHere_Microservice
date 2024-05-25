package interfaces

import "TweetHere-API/pkg/utils/models"

type ChatClient interface {
	GetChat(userid string, req models.ChatRequest) ([]models.TempMessage, error)
}
