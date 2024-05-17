package client

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/config"
	pb "TweetHere-API/pkg/pb/tweet"
	"TweetHere-API/pkg/utils/models"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"

	"google.golang.org/grpc"
)

type tweetClient struct {
	Client pb.TweetServiceClient
}

func NewTweetClient(cfg config.Config) interfaces.TweetClient {
	grpcConnection, err := grpc.Dial(cfg.TweeSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewTweetServiceClient(grpcConnection)

	return &tweetClient{
		Client: grpcClient,
	}
}

func (ad *tweetClient) AddTweet(id int, file *multipart.FileHeader, postdetails string) error {
	filecontent, err := file.Open()
	if err != nil {
		return err
	}

	fileBytes, err := ioutil.ReadAll(filecontent)
	if err != nil {
		return err
	}
	_, errs := ad.Client.AddTweet(context.Background(), &pb.AddTweetRequest{
		Id:           int64(id),
		Descritption: postdetails,
		File:         fileBytes,
	})
	if errs != nil {
		return errs
	}
	return nil
}

func (ad *tweetClient) AddTweet2(id int, postdetails string) error {

	_, errs := ad.Client.AddTweet2(context.Background(), &pb.AddTweet2Request{
		Id:           int64(id),
		Descritption: postdetails,
	})
	if errs != nil {
		return errs
	}
	return nil
}

func (ad *tweetClient) GetOurTweet(id int) ([]models.PostResponse, error) {
	res, err := ad.Client.GetOurTweet(context.Background(), &pb.GetOurTweetRequest{
		Id: int64(id),
	})
	if err != nil {
		return []models.PostResponse{}, errors.New("error in getting post details")
	}

	var posdetials []models.PostResponse

	for _, ud := range res.Postdetails {
		posdetials = append(posdetials, models.PostResponse{
			UserID:      int(ud.Id),
			Description: ud.Description,
			Url:         ud.Url,
			CreatedAt:   ud.Time.AsTime(),
		})
	}
	return posdetials, nil
}

func (ad *tweetClient) GetOthersTweet(id int) ([]models.PostResponse, error) {
	res, err := ad.Client.GetOthersTweet(context.Background(), &pb.GetOthersTweetRequest{
		Id: int64(id),
	})
	if err != nil {
		return []models.PostResponse{}, errors.New("error in getting post details")
	}

	var posdetials []models.PostResponse

	for _, ud := range res.Postdetailss {
		posdetials = append(posdetials, models.PostResponse{
			UserID:      int(ud.Id),
			Description: ud.Description,
			Url:         ud.Url,
			CreatedAt:   ud.Time.AsTime(),
		})
	}
	return posdetials, nil
}

func (ad *tweetClient) EditTweet(id int, postid int, discription string) error {
	_, err := ad.Client.EditTweet(context.Background(), &pb.EditTweetRequest{
		Id:          int64(id),
		Postid:      int64(postid),
		Description: discription,
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) DeletePost(id int, postid int) error {
	_, err := ad.Client.DeletePost(context.Background(), &pb.DeletePostRequest{
		Id:     int64(id),
		Postid: int64(postid),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) LikePost(id int, postid int) error {
	_, err := ad.Client.LikePost(context.Background(), &pb.LikePostRequest{
		Id:     int64(id),
		Postid: int64(postid),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) UnLikePost(id int, postid int) error {
	_, err := ad.Client.UnLikePost(context.Background(), &pb.UnLikePostRequest{
		Id:     int64(id),
		Postid: int64(postid),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) SavePost(id int, postid int) error {
	_, err := ad.Client.SavePost(context.Background(), &pb.SavePostRequest{
		Id:     int64(id),
		Postid: int64(postid),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) UnSavePost(id int, postid int) error {
	_, err := ad.Client.UnSavePost(context.Background(), &pb.UnSavePostRequest{
		Id:     int64(id),
		Postid: int64(postid),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) RplyCommentPost(id int, postid int, comment string, parentid int) error {
	_, err := ad.Client.RplyCommentPost(context.Background(), &pb.RplyCommentPostRequest{
		Id:       int64(id),
		Postid:   int64(postid),
		Comment:  comment,
		Parentid: int64(parentid),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) CommentPost(id int, postid int, comment string) error {
	_, err := ad.Client.CommentPost(context.Background(), &pb.CommentPostRequest{
		Id:      int64(id),
		Postid:  int64(postid),
		Comment: comment,
	})
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetClient) GetComments(postid int)([]models.CommentsResponse,error){
	details,err := ad.Client.GetComments(context.Background(),&pb.GetCommentsRequest{
		Postid: int64(postid),
	})
	if err != nil{
		return []models.CommentsResponse{},err
	}

	var commentdetails []models.CommentsResponse

	for _,ud := range details.Comentdetails{
		commentdetails = append(commentdetails, models.CommentsResponse{
			Username: ud.Username,
			Profile: ud.Profile,
			Comment: ud.Comment,
			CreatedAt: ud.Time.AsTime(),
		})
	}
	return commentdetails,nil
}