package service

import (
	"context"
	"fmt"
	"time"
	"tweethere-chat/pkg/logging"
	pb "tweethere-chat/pkg/pb/chat"
	interfaces "tweethere-chat/pkg/usecase/interface"
	"tweethere-chat/pkg/utils/models"
)

type ChatServer struct {
	chatUseCase interfaces.ChatUseCase
	pb.UnimplementedChatServiceServer
}

func NeChatServer(usecase interfaces.ChatUseCase) pb.ChatServiceServer {
	return &ChatServer{
		chatUseCase: usecase,
	}
}

func (ad *ChatServer) GetFriendChat(ctx context.Context, req *pb.GetFriendChatRequest) (*pb.GetFriendChatResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "GetFriendChat")
	logEntry.Info("Processing GetFriendChat request with limit:", req.GetLimit(), "and offset:", req.GetOffSet())
	ind, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println("details is ", req)
	result, err := ad.chatUseCase.GetFriendChat(req.UserID, req.FriendID, models.Pagination{Limit: req.Limit, OffSet: req.OffSet})
	if err != nil {
		logEntry.WithError(err).Errorf("Error getting friend chat ")

		return nil, err
	}
	var finalResult []*pb.Message
	for _, val := range result {
		finalResult = append(finalResult, &pb.Message{
			MessageID:   val.ID,
			SenderId:    val.SenderID,
			RecipientId: val.RecipientID,
			Content:     val.Content,
			Timestamp:   val.Timestamp.In(ind).String(),
		})
	}
	return &pb.GetFriendChatResponse{FriendChat: finalResult}, nil
}
