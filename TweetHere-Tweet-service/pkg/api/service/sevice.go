package service

import (
	"context"
	"fmt"
	pb "tweet-service/pkg/pb/tweet"
	interfaces "tweet-service/pkg/usecase/interface"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TweetServer struct {
	tweetUseCase interfaces.TweetUseCase
	pb.UnimplementedTweetServiceServer
}

func NewTweetServer(usecase interfaces.TweetUseCase) pb.TweetServiceServer {
	return &TweetServer{
		tweetUseCase: usecase,
	}
}

func (ad *TweetServer) AddTweet(ctx context.Context, req *pb.AddTweetRequest) (*pb.AddTweetResponse, error) {
	id := req.Id
	file := req.File
	description := req.Descritption
	fmt.Println("heloooooooooooooooooooooooooo")
	err := ad.tweetUseCase.AddTweet(id, file, description)
	if err != nil {
		return &pb.AddTweetResponse{}, err
	}
	return &pb.AddTweetResponse{}, nil

}

func (ad *TweetServer) AddTweet2(ctx context.Context, req *pb.AddTweet2Request) (*pb.AddTweet2Response, error) {
	id := req.Id
	fmt.Println("hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	description := req.Descritption

	err := ad.tweetUseCase.AddTweet2(id, description)
	if err != nil {
		return &pb.AddTweet2Response{}, err
	}
	return &pb.AddTweet2Response{}, nil

}

func (ad *TweetServer) GetOurTweet(ctx context.Context, req *pb.GetOurTweetRequest) (*pb.GetOurTweetResponse, error) {
	id := req.Id

	result, err := ad.tweetUseCase.GetOurTweet(int(id))
	if err != nil {
		return &pb.GetOurTweetResponse{}, err
	}
	var postdetails []*pb.TweetReponse

	for _, details := range result {
		postdetails = append(postdetails, &pb.TweetReponse{
			Id:          int64(details.UserID),
			Description: details.Description,
			Url:         details.Url,
			Time:        timestamppb.New(details.CreatedAt),
		})
	}
	return &pb.GetOurTweetResponse{
		Postdetails: postdetails,
	}, nil
}

func (ad *TweetServer) GetOthersTweet(ctx context.Context, req *pb.GetOthersTweetRequest) (*pb.GetOthersTweetResponse, error) {
	id := req.Id

	result, err := ad.tweetUseCase.GetOthersTweet(int(id))
	if err != nil {
		return &pb.GetOthersTweetResponse{}, err
	}
	var postdetails []*pb.TweetReponse

	for _, details := range result {
		postdetails = append(postdetails, &pb.TweetReponse{
			Id:          int64(details.UserID),
			Description: details.Description,
			Url:         details.Url,
			Time:        timestamppb.New(details.CreatedAt),
		})
	}
	return &pb.GetOthersTweetResponse{
		Postdetailss: postdetails,
	}, nil
}

func (ad *TweetServer) EditTweet(ctx context.Context, req *pb.EditTweetRequest) (*pb.EditTweetResponse, error) {
	id := req.Id
	postid := req.Postid
	description := req.Description

	err := ad.tweetUseCase.EditTweet(int(id), int(postid), description)
	if err != nil {
		return &pb.EditTweetResponse{}, err
	}
	return &pb.EditTweetResponse{}, nil
}

func (ad *TweetServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.DeletePost(int(id), int(postid))
	if err != nil {
		return &pb.DeletePostResponse{}, err
	}
	return &pb.DeletePostResponse{}, nil
}

func (ad *TweetServer) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.LikePost(int(id), int(postid))
	if err != nil {
		return &pb.LikePostResponse{}, err
	}
	return &pb.LikePostResponse{}, nil
}

func (ad *TweetServer) UnLikePost(ctx context.Context, req *pb.UnLikePostRequest) (*pb.UnLikePostResponse, error) {
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.UnLikePost(int(id), int(postid))
	if err != nil {
		return &pb.UnLikePostResponse{}, err
	}
	return &pb.UnLikePostResponse{}, nil
}

func (ad *TweetServer) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.SavePost(int(id), int(postid))
	if err != nil {
		return &pb.SavePostResponse{}, err
	}

	return &pb.SavePostResponse{}, nil
}

func (ad *TweetServer) UnSavePost(ctx context.Context, req *pb.UnSavePostRequest) (*pb.UnSavePostRespone, error) {
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.UnSavePost(int(id), int(postid))
	if err != nil {
		return &pb.UnSavePostRespone{}, err
	}
	return &pb.UnSavePostRespone{}, nil

}

func (ad *TweetServer) CommentPost(ctx context.Context, req *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	id := req.Id
	postid := req.Postid
	comment := req.Comment

	err := ad.tweetUseCase.CommentPost(int(id), int(postid), comment)
	if err != nil {
		return &pb.CommentPostResponse{}, err
	}
	return &pb.CommentPostResponse{}, nil
}

func (ad *TweetServer) RplyCommentPost(ctx context.Context, req *pb.RplyCommentPostRequest) (*pb.RplyCommentPostResponse, error) {
	id := req.Id
	postid := req.Postid
	comment := req.Comment
	parentid := req.Parentid
	fmt.Println("helloooo")

	err := ad.tweetUseCase.RplyCommentPost(int(id), int(postid), comment, int(parentid))
	if err != nil {
		return &pb.RplyCommentPostResponse{}, err
	}
	return &pb.RplyCommentPostResponse{}, nil

}

// func (ad *TweetServer) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
//     postid := req.Postid

// 	result ,err := ad.tweetUseCase.
// }
