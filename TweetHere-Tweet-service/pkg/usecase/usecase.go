package usecase

import (
	"errors"
	"fmt"
	Ainterfaces "tweet-service/pkg/client/interface"
	"tweet-service/pkg/helper"
	interfaces "tweet-service/pkg/repository/interface"
	services "tweet-service/pkg/usecase/interface"
	"tweet-service/pkg/utils/models"

	"github.com/google/uuid"
)

type tweetUseCase struct {
	tweetRepository interfaces.TweetRepository
	authRepository  Ainterfaces.AuthClient
}

func NewTweetUseCase(repository interfaces.TweetRepository, authRepo Ainterfaces.AuthClient) services.TweetUseCase {
	return &tweetUseCase{
		tweetRepository: repository,
		authRepository:  authRepo,
	}
}

func (ad *tweetUseCase) AddTweet(id int64, file []byte, discription string) error {
	userexist, err := ad.authRepository.DoesUserExist(id)
	if err != nil {
		fmt.Println("error finded", err)
		return err
	}

	if !userexist {
		return errors.New("problem in checking user exits")
	}

	tID, errs := ad.tweetRepository.AddTweet(int(id), discription)
	if errs != nil {

		return errs
	}

	username, err := ad.authRepository.FindUserName(id)
	if err != nil {
		fmt.Println("zzzzzzzzzzzzzzzzzzz")
		return err
	}

	fileUID := uuid.New()
	fileName := fileUID.String()
	s3path := username + fileName

	url, errp := helper.AddImageToAwsS3(file, s3path)
	if errp != nil {

		return errp
	}

	errr := ad.tweetRepository.UploadMedia(tID, url)
	if errr != nil {
		return errr
	}
	return nil

}

func (ad *tweetUseCase) AddTweet2(id int64, discription string) error {
	userexist, err := ad.authRepository.DoesUserExist(id)
	if err != nil {
		fmt.Println("error finded", err)
		return err
	}

	if !userexist {
		return errors.New("problem in checking user exits")
	}

	_, errs := ad.tweetRepository.AddTweet(int(id), discription)
	if errs != nil {

		return errs
	}
	return nil
}

func (ad *tweetUseCase) GetOurTweet(id int) ([]models.PostResponse, error) {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return []models.PostResponse{}, errors.New("user doesnt exist")
	}

	details, err := ad.tweetRepository.GetOurTweet(id)
	if err != nil {
		return []models.PostResponse{}, err
	}
	fmt.Println("her are uellllllllllllll", details)
	return details, nil
}

func (ad *tweetUseCase) GetOthersTweet(id int) ([]models.PostResponse, error) {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return []models.PostResponse{}, errors.New("user doesnt exist")
	}

	details, err := ad.tweetRepository.GetOurTweet(id)
	if err != nil {
		return []models.PostResponse{}, err
	}
	fmt.Println("her are uellllllllllllll", details)
	return details, nil
}

func (ad *tweetUseCase) EditTweet(id int, tweetid int, description string) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}
	err := ad.tweetRepository.EditTweet(id, tweetid, description)
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetUseCase) DeletePost(id int, postid int) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}

	err := ad.tweetRepository.DeletePost(id, postid)
	if err != nil {
		return err
	}
	return nil
}
func (ad *tweetUseCase) LikePost(id int, postid int) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}

	err := ad.tweetRepository.LikePost(id, postid)
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetUseCase) UnLikePost(id int, postid int) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}

	err := ad.tweetRepository.UnLikePost(id, postid)
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetUseCase) SavePost(id int, postid int) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}
	err := ad.tweetRepository.SavePost(id, postid)
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetUseCase) UnSavePost(id int, postid int) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}

	err := ad.tweetRepository.UnSavePost(id, postid)
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetUseCase) CommentPost(id int, postid int, comment string) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}

	err := ad.tweetRepository.CommentPost(id, postid, comment)
	if err != nil {
		return err
	}
	return nil
}

func (ad *tweetUseCase) RplyCommentPost(id int, postid int, comment string, parentid int) error {
	userexist, _ := ad.authRepository.DoesUserExist(int64(id))
	if !userexist {
		return errors.New("user is not exist")
	}

	postexist := ad.tweetRepository.PostExist(postid)
	if !postexist {
		return errors.New("this post is doesnt exist")
	}

	err := ad.tweetRepository.RplyCommentPost(id, postid, comment, parentid)
	if err != nil {
		return err
	}
	return nil
}

// func (ad *tweetUseCase) GetComments(postid int)([]models.CommentsResponse,error){
	
// 	details,err := ad.tweetRepository.G
// }