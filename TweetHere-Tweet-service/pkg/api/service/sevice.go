package service

import (
	"context"
	"fmt"
	"tweet-service/pkg/logging"
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
	logEntry := logging.GetLogger().WithField("method", "AddTweet")
	logEntry.Info("Processing AddTweet request for user ID:", req.GetId(), ", with file:", req.GetFile() != nil, ", description:", req.GetDescritption())
	id := req.Id
	file := req.File
	description := req.Descritption
	fmt.Println("heloooooooooooooooooooooooooo")
	err := ad.tweetUseCase.AddTweet(id, file, description)
	if err != nil {
		logEntry.WithError(err).Error("Error adding tweet")
		return &pb.AddTweetResponse{}, err
	}
	logEntry.Info("Tweet successfully added for user ID:", id)
	return &pb.AddTweetResponse{}, nil

}

func (ad *TweetServer) AddTweet2(ctx context.Context, req *pb.AddTweet2Request) (*pb.AddTweet2Response, error) {
	logEntry := logging.GetLogger().WithField("method", "AddTweet2")
	logEntry.Info("Processing AddTweet2 request for user ID:", req.GetId(), ", description:", req.GetDescritption())
	id := req.Id
	fmt.Println("hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	description := req.Descritption

	err := ad.tweetUseCase.AddTweet2(id, description)
	if err != nil {
		logEntry.WithError(err).Error("Error adding tweet")
		return &pb.AddTweet2Response{}, err
	}
	logEntry.Info("Tweet successfully added for user ID:", id)
	return &pb.AddTweet2Response{}, nil

}

func (ad *TweetServer) GetOurTweet(ctx context.Context, req *pb.GetOurTweetRequest) (*pb.GetOurTweetResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "GetOurTweet")
	logEntry.Info("Processing GetOurTweet request for user ID:", req.GetId())
	id := req.Id

	result, err := ad.tweetUseCase.GetOurTweet(int(id))
	if err != nil {
		logEntry.WithError(err).Error("Error getting tweets")
		return &pb.GetOurTweetResponse{}, err
	}
	var postdetails []*pb.TweetReponse

	for _, details := range result {
		postdetails = append(postdetails, &pb.TweetReponse{
			Id:          int64(details.UserID),
			Description: details.Description,
			Url:         details.Url,
			Like:        int64(details.Likes),
			Comment:     int64(details.Comments),
			Time:        timestamppb.New(details.CreatedAt),
		})
	}
	logEntry.Info("Successfully retrieved tweets for user ID", req.Id)
	return &pb.GetOurTweetResponse{
		Postdetails: postdetails,
	}, nil
}

func (ad *TweetServer) GetOthersTweet(ctx context.Context, req *pb.GetOthersTweetRequest) (*pb.GetOthersTweetResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "GetOthersTweet")
	logEntry.Info("Processing GetOthersTweet request for user ID:", req.GetId())
	id := req.Id

	result, err := ad.tweetUseCase.GetOthersTweet(int(id))
	if err != nil {
		logEntry.WithError(err).Error("Error getting tweets")
		return &pb.GetOthersTweetResponse{}, err
	}
	var postdetails []*pb.TweetReponse

	for _, details := range result {
		postdetails = append(postdetails, &pb.TweetReponse{
			Id:          int64(details.UserID),
			Description: details.Description,
			Url:         details.Url,
			Like:        int64(details.Likes),
			Comment:     int64(details.Comments),
			Time:        timestamppb.New(details.CreatedAt),
		})
	}
	logEntry.Info("Successfully retrieved tweets for user ID", req.Id)
	return &pb.GetOthersTweetResponse{
		Postdetailss: postdetails,
	}, nil
}

func (ad *TweetServer) EditTweet(ctx context.Context, req *pb.EditTweetRequest) (*pb.EditTweetResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "EditTweet")
	logEntry.Info("Processing EditTweet request for user ID:", req.GetId(), ", post ID:", req.GetPostid(), ", description:", req.GetDescription())
	id := req.Id
	postid := req.Postid
	description := req.Description

	err := ad.tweetUseCase.EditTweet(int(id), int(postid), description)
	if err != nil {
		logEntry.WithError(err).Error("Error editing tweet")
		return &pb.EditTweetResponse{}, err
	}
	logEntry.Info("Tweet successfully edited for user ID:", req.Id)
	return &pb.EditTweetResponse{}, nil
}

func (ad *TweetServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "DeletePost")
	logEntry.Info("Processing DeletePost request for user ID:", req.GetId(), ", post ID:", req.GetPostid())

	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.DeletePost(int(id), int(postid))
	if err != nil {
		logEntry.WithError(err).Error("Error deleting post")
		return &pb.DeletePostResponse{}, err
	}
	logEntry.Info("Post successfully deleted for user ID:", req.Id)
	return &pb.DeletePostResponse{}, nil
}

func (ad *TweetServer) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "LikePost")
	logEntry.Info("Processing LikePost request for user ID:", req.GetId(), ", post ID:", req.GetPostid())

	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.LikePost(int(id), int(postid))
	if err != nil {
		logEntry.WithError(err).Error("Error liking post")
		return &pb.LikePostResponse{}, err
	}
	logEntry.Info("Post successfully liked for user ID:", req.Id)
	return &pb.LikePostResponse{}, nil
}

func (ad *TweetServer) UnLikePost(ctx context.Context, req *pb.UnLikePostRequest) (*pb.UnLikePostResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "UnLikePost")
	logEntry.Info("Processing UnLikePost request for user ID:", req.GetId(), ", post ID:", req.GetPostid())
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.UnLikePost(int(id), int(postid))
	if err != nil {
		logEntry.WithError(err).Error("Error unliking post")
		return &pb.UnLikePostResponse{}, err
	}
	logEntry.Info("Post successfully unliked for user ID:", req.Id)
	return &pb.UnLikePostResponse{}, nil
}

func (ad *TweetServer) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "SavePost")
	logEntry.Info("Processing SavePost request for user ID:", req.GetId(), ", post ID:", req.GetPostid())
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.SavePost(int(id), int(postid))
	if err != nil {
		logEntry.WithError(err).Error("Error saving post")
		return &pb.SavePostResponse{}, err
	}
	logEntry.Info("Post successfully saved for user ID:", req.Id)
	return &pb.SavePostResponse{}, nil
}

func (ad *TweetServer) UnSavePost(ctx context.Context, req *pb.UnSavePostRequest) (*pb.UnSavePostRespone, error) {
	logEntry := logging.GetLogger().WithField("method", "UnSavePost")
	logEntry.Info("Processing UnSavePost request for user ID:", req.GetId(), ", post ID:", req.GetPostid())
	id := req.Id
	postid := req.Postid

	err := ad.tweetUseCase.UnSavePost(int(id), int(postid))
	if err != nil {
		logEntry.WithError(err).Error("Error unsaving post")
		return &pb.UnSavePostRespone{}, err
	}
	logEntry.Info("Post successfully unsaved for user ID:", req.Id)
	return &pb.UnSavePostRespone{}, nil

}

func (ad *TweetServer) CommentPost(ctx context.Context, req *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "CommentPost")
	logEntry.Info("Processing CommentPost request for user ID:", req.GetId(), ", post ID:", req.GetPostid(), ", comment:", req.GetComment())
	id := req.Id
	postid := req.Postid
	comment := req.Comment

	err := ad.tweetUseCase.CommentPost(int(id), int(postid), comment)
	if err != nil {
		logEntry.WithError(err).Error("Error commenting on post")
		return &pb.CommentPostResponse{}, err
	}
	logEntry.Info("Comment successfully added to post for user ID:", req.Id)
	return &pb.CommentPostResponse{}, nil
}

func (ad *TweetServer) RplyCommentPost(ctx context.Context, req *pb.RplyCommentPostRequest) (*pb.RplyCommentPostResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "RplyCommentPost")
	logEntry.Info("Processing RplyCommentPost request for user ID:", req.GetId(), ", post ID:", req.GetPostid(), ", comment:", req.GetComment(), ", parent ID:", req.GetParentid())
	id := req.Id
	postid := req.Postid
	comment := req.Comment
	parentid := req.Parentid
	fmt.Println("helloooo")

	err := ad.tweetUseCase.RplyCommentPost(int(id), int(postid), comment, int(parentid))
	if err != nil {
		logEntry.WithError(err).Error("Error replying to comment")
		return &pb.RplyCommentPostResponse{}, err
	}
	logEntry.Info("Reply successfully added to comment for user ID:", req.Id)
	return &pb.RplyCommentPostResponse{}, nil

}

func (ad *TweetServer) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "GetComments")
	logEntry.Info("Processing GetComments request for post ID:", req.GetPostid())

	postid := req.Postid

	result, err := ad.tweetUseCase.GetComments(int(postid))
	if err != nil {
		logEntry.WithError(err).Error("Error getting comments")
		return &pb.GetCommentsResponse{}, err
	}

	var comments []*pb.CommentsResponse
	for _, post := range result {
		details := &pb.CommentsResponse{
			Id:       int64(post.UserId),
			Username: post.Username,
			Profile:  post.Profile,
			Comment:  post.Comment,
			Time:     timestamppb.New(post.CreatedAt),
		}
		comments = append(comments, details)
	}
	logEntry.Info("Successfully retrieved comments for post ID:", req.Postid)
	return &pb.GetCommentsResponse{
		Comentdetails: comments,
	}, nil

}

func (ad *TweetServer) EditComments(ctx context.Context, req *pb.EditCommentsRequet) (*pb.EditCommentsResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "EditComments")
	logEntry.Info("Processing EditComments request for user ID:", req.GetId(), ", comment ID:", req.GetCommentid(), ", new comment:", req.GetComment())
	id := req.Id
	commentid := req.Commentid
	comment := req.Comment

	err := ad.tweetUseCase.EditComments(int(id), int(commentid), comment)
	if err != nil {
		logEntry.WithError(err).Error("Error editing comment")
		return &pb.EditCommentsResponse{}, err
	}
	logEntry.Info("Comment successfully edited for user ID:", req.Id)
	return &pb.EditCommentsResponse{}, nil
}

func (ad *TweetServer) DeleteComments(ctx context.Context, req *pb.DeleteCommentsRequest) (*pb.DeleteCommentsResponse, error) {
	logEntry := logging.GetLogger().WithField("method", "DeleteComments")
	logEntry.Info("Processing DeleteComments request for user ID:", req.GetId(), ", comment ID:", req.GetCommentid())
	id := req.Id
	commentid := req.Commentid

	err := ad.tweetUseCase.DeleteComments(int(id), int(commentid))
	if err != nil {
		logEntry.WithError(err).Error("Error deleting comment")
		return &pb.DeleteCommentsResponse{}, err
	}
	logEntry.Info("Comment successfully deleted for user ID:", req.Id)
	return &pb.DeleteCommentsResponse{}, nil
}

func (ad *TweetServer) Home(ctx context.Context, req *pb.HomeRequest) (*pb.HomeResponse, error) {
	data, err := ad.tweetUseCase.Home(int(req.Userid))
	if err != nil {
		return &pb.HomeResponse{}, err
	}
	var allpostResponses []*pb.CreatePostResponse
	for _, post := range data {
		userdata := &pb.UserData{
			Userid:   int64(post.Author.UserID),
			Username: post.Author.Username,
			Imageurl: post.Author.Profile,
		}

		details := &pb.CreatePostResponse{
			Id:          int64(post.ID),
			User:        userdata,
			Description: post.Description,
			Url:         post.Url,
			Like:        int64(post.Likes),
			Comment:     int64(post.Comments),
			Time:        timestamppb.New(post.CreatedAt),
		}

		allpostResponses = append(allpostResponses, details)
	}
	return &pb.HomeResponse{
		Allpost: allpostResponses,
	}, nil
}
