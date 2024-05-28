package service

import (
	"context"
	pb "tweethere-Notification/pkg/pb/noti"
	interfaces "tweethere-Notification/pkg/usecase/interface"
	"tweethere-Notification/pkg/utils/models"
)

type NotiServer struct {
	notiUsecase interfaces.NotiUseCase
	pb.UnimplementedNotificationServiceServer
}

func NewnotiServer(usecase interfaces.NotiUseCase) pb.NotificationServiceServer {
	return &NotiServer{
		notiUsecase: usecase,
	}
}

func (ad *NotiServer) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	userid := req.UserID

	result, err := ad.notiUsecase.GetNotification(int(userid), models.Pagination{Limit: int(req.Limit), Offset: int(req.Offset)})
	if err != nil {
		return nil, err
	}
	var final []*pb.Message

	for _, v := range result {
		final = append(final, &pb.Message{
			UserId:   int64(v.UserID),
			Username: v.Username,
			Profile:  v.Profile,
			Message:  v.Message,
			Time:     v.CreatedAt,
		})
	}
	return &pb.GetNotificationResponse{
		Notification: final,
	}, nil
}
