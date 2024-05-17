package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/utils/models"
	"TweetHere-API/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAuthHandler(adminClient interfaces.AdminClient) *AuthHandler {
	return &AuthHandler{
		GRPC_Client: adminClient,
	}
}

func (ad *AuthHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		fmt.Println("eroorrrr ivide und", err)
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenthicate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "successfully logged in", admin, nil)
	c.JSON(http.StatusOK, success)
}

//user side/////////////////////////////////

func (ad *AuthHandler) UserSignUp(c *gin.Context) {
	var userdetails models.UserSignup

	// Print userdetails before binding JSON
	fmt.Println("User details before binding JSON:", userdetails)

	// Pass a pointer to ShouldBindJSON
	if err := c.ShouldBindJSON(&userdetails); err != nil {
		fmt.Println("errrrrrrrrr", err)
		errs := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := ad.GRPC_Client.UserSignup(userdetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	successres := response.ClientResponse(http.StatusOK, "successfully authenticated user", user, nil)
	c.JSON(http.StatusOK, successres)
}

func (ad *AuthHandler) UserLogin(c *gin.Context) {
	var userdetails models.UserLogin

	if err := c.ShouldBindJSON(&userdetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := ad.GRPC_Client.UserLogin(userdetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "user login failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "successfully logined the user", user, nil)
	c.JSON(http.StatusOK, successres)

}

func (ad *AuthHandler) UserUpdateProfile(c *gin.Context) {
	var userdetails models.UserProfile

	idstring, _ := c.Get("id")
	id, _ := idstring.(int)

	fmt.Println("iddddddddddddddddddd is ", id)

	if err := c.ShouldBindJSON(&userdetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ad.GRPC_Client.UserUpdateProfile(userdetails, id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "update profile failed", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "successfully updated userprofile", user, nil)
	c.JSON(http.StatusOK, successres)
}

func (ad *AuthHandler) GetUser(c *gin.Context) {
	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)

	if err != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "page number not in correct foramt", nil, err.Error())
		c.JSON(http.StatusBadRequest, errres)
		return

	}
	fmt.Println("ssssssssssssssssss")
	user, errs := ad.GRPC_Client.GetUser(page)
	if errs != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "error in getting the details", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}

	successres := response.ClientResponse(http.StatusOK, "succesfully got all user details", user, nil)
	c.JSON(http.StatusOK, successres)

}

func (ad *AuthHandler) BlockUser(c *gin.Context) {
	id := c.Query("id")

	err := ad.GRPC_Client.BlockUser(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in block user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully blocked", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *AuthHandler) UnBlockUser(c *gin.Context) {
	id := c.Query("id")

	err := ad.GRPC_Client.UnBlockUser(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in unblock user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully unblocked", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *AuthHandler) ChangePassword(c *gin.Context) {
	idstring, _ := c.Get("id")

	id, _ := idstring.(int)

	var ChangePassword models.ChangePassword

	if err := c.BindJSON(&ChangePassword); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	fmt.Println("changed passwords details are", ChangePassword)
	fmt.Println("id for changepasword is ", id)

	if err := ad.GRPC_Client.ChangePassword(id, ChangePassword); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in changepassword", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully changed the password of you", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

func (ad *AuthHandler) GetUserDetails(c *gin.Context) {
	id_string, _ := c.Get("id")
	id, _ := id_string.(int)

	userdetails, err := ad.GRPC_Client.GetUserDetails(id)
	if err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "failed to get the userdetails", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "Successfully got the details of you", userdetails, nil)
	c.JSON(http.StatusOK, succesres)
}

func (uh *AuthHandler) UserOTPLogin(c *gin.Context) {
	var userDetails models.UserOTPLogin
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	otp, err := uh.GRPC_Client.UserOTPLogin(userDetails.Email)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to generate OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "OTP generated successfully", map[string]string{"OTP": otp}, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AuthHandler) FollowReq(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	userId := c.Query("id")
	userID, _ := strconv.Atoi(userId)

	err := ad.GRPC_Client.FollowReq(id, userID)
	if err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "could not send the request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully send the requset ", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *AuthHandler) AcceptFollowreq(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	userId := c.Query("id")
	userID, _ := strconv.Atoi(userId)
	err := ad.GRPC_Client.AcceptFollowreq(id, userID)
	if err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "accepting the request is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "succesfully accepted the request", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *AuthHandler) Unfollow(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	useId := c.Query("id")
	userID, _ := strconv.Atoi(useId)

	err := ad.GRPC_Client.Unfollow(id, userID)

	if err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "unfollowing is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return

	}
	succesres := response.ClientResponse(http.StatusOK, "successfully unfollowed the follower", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

func (ad *AuthHandler) Followers(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	details, err := ad.GRPC_Client.Followers(id)
	if err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "get the followers failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "successfully got followers", details, nil)
	c.JSON(http.StatusOK, successres)
}

func (ad *AuthHandler) Followings(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	details, err := ad.GRPC_Client.Followings(id)
	if err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "get the followings failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "successfully got followings", details, nil)
	c.JSON(http.StatusOK, successres)
}

func (ad *AuthHandler) SendOTP(c *gin.Context) {
	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
	}

	err := ad.GRPC_Client.SendOTP(phone.PhoneNumber)
	if err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "seding otp is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
	}

}

func (ad *AuthHandler) VerifyOTP(c *gin.Context) {
	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.GRPC_Client.VerifyOTP(code)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "verifying the otp is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "successfully verified ", users, nil)
	c.JSON(http.StatusOK, successres)

}

func (ad *AuthHandler) UploadProfilepic(c *gin.Context) {
	id_string, _ := c.Get("id")
	id := id_string.(int)

	form, err := c.MultipartForm()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "retreiving images from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		errorRes := response.ClientResponse(http.StatusBadRequest, "no files are provided", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	for _, file := range files {
		err := ad.GRPC_Client.UploadProfilepic(id, file)
		if err != nil {
			errorRes := response.ClientResponse(http.StatusBadRequest, "could not change one or more images", nil, err.Error())
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}
	}
	successres := response.ClientResponse(http.StatusOK, "uploaded successfully ", nil, nil)
	c.JSON(http.StatusOK, successres)

}
