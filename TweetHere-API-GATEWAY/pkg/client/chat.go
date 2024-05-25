package client

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/config"
	pb "TweetHere-API/pkg/pb/chat"
	"TweetHere-API/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type chatClient struct {
	Client pb.ChatServiceClient
}

func NewChatClient(cfg config.Config) interfaces.ChatClient {
	grpcConnection, err := grpc.Dial(cfg.ChatSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewChatServiceClient(grpcConnection)

	return &chatClient{
		Client: grpcClient,
	}
}

func (ad *chatClient) GetChat(userid string, req models.ChatRequest) ([]models.TempMessage, error) {
	fmt.Println("zzzzzzzzzzz", userid, req)
	data, err := ad.Client.GetFriendChat(context.Background(), &pb.GetFriendChatRequest{
		UserID:   userid,
		FriendID: req.FriendID,
		OffSet:   req.Offset,
		Limit:    req.Limit,
	})
	if err != nil {
		return []models.TempMessage{}, err
	}
	var response []models.TempMessage

	for _, v := range data.FriendChat {
		chatResponse := models.TempMessage{
			SenderID:    v.SenderId,
			RecipientID: v.RecipientId,
			Content:     v.Content,
			Timestamp:   v.Timestamp,
		}
		response = append(response, chatResponse)
	}
	return response, nil
}
