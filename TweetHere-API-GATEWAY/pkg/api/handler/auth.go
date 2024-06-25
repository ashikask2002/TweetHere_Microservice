package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/logging"
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

// @Summary		Admin Login
// @Description	Login handler for Zsoxial admins
// @Tags			Admin
// @Accept 			json
// @Produce 		json
// @Param			adminDetails	body		models.AdminLogin	true	"Admin login details"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/admin/login  [POST]
func (ad *AuthHandler) LoginHandler(c *gin.Context) {

	logEntry := logging.GetLogger().WithField("context", "LogginHandler")
	logEntry.Info("Processing Loggin request")

	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		fmt.Println("eroorrrr ivide und", err)
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during AdminLogin RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenthicate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logEntry.Info("Login successful for user")
	success := response.ClientResponse(http.StatusOK, "successfully logged in", admin, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		User Sign Up
// @Description	Sign up handler for new users
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			userDetails	body		models.UserSignup	true	"User sign up details"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Failure		500			{object}	response.Response{}
// @Router			/user/signup  [POST]
func (ad *AuthHandler) UserSignUp(c *gin.Context) {
	var userdetails models.UserSignup
	logEntry := logging.GetLogger().WithField("context", "UserSignupHandler")

	fmt.Println("User details before binding JSON:", userdetails)
	logEntry.Info("Processing user Signup request")

	if err := c.ShouldBindJSON(&userdetails); err != nil {
		fmt.Println("errrrrrrrrr", err)
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := ad.GRPC_Client.UserSignup(userdetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during User Signup rpc call")
		errs := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("user signup succesfull for user", user.User.Firstname)
	successres := response.ClientResponse(http.StatusOK, "successfully authenticated user", user, nil)
	c.JSON(http.StatusOK, successres)
}

// @Summary		User Login
// @Description	Login handler for users
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			userDetails	body		models.UserLogin	true	"User login details"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Failure		500			{object}	response.Response{}
// @Router			/user/login  [POST]
func (ad *AuthHandler) UserLogin(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UserLoginHandler")
	logEntry.Info("Processing user Login  request")
	var userdetails models.UserLogin

	if err := c.ShouldBindJSON(&userdetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := ad.GRPC_Client.UserLogin(userdetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during UserLogin rpc call")
		errs := response.ClientResponse(http.StatusBadRequest, "user login failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("user login succesful for user", user.User.Firstname)
	successres := response.ClientResponse(http.StatusOK, "successfully logined the user", user, nil)
	c.JSON(http.StatusOK, successres)

}

// @Summary		User Update Profile
// @Description	Update profile handler for users
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security 		Bearer
// @Param			userDetails	body		models.UserProfile	true	"User profile details"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Failure		502			{object}	response.Response{}
// @Router			/users/profile  [PUT]
func (ad *AuthHandler) UserUpdateProfile(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UserUpdateProfileHandler")
	logEntry.Info("Processing user UpdateProfile  request")
	var userdetails models.UserProfile

	idstring, _ := c.Get("id")
	id, _ := idstring.(int)

	fmt.Println("iddddddddddddddddddd is ", id)

	if err := c.ShouldBindJSON(&userdetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := ad.GRPC_Client.UserUpdateProfile(userdetails, id)
	if err != nil {
		logEntry.WithError(err).Error("Error during UserUpdateProfile rpc call")
		errs := response.ClientResponse(http.StatusBadRequest, "update profile failed", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	logEntry.Info("user UpdateProfile succesful for user")
	successres := response.ClientResponse(http.StatusOK, "successfully updated userprofile", user, nil)
	c.JSON(http.StatusOK, successres)
}

// @Summary		Get User Details
// @Description	Retrieves user details for admin management
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security 		bearer
// @Param			page	query	int	true	"Page number for pagination"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admins/userdetails [GET]
func (ad *AuthHandler) GetUser(c *gin.Context) {

	logEntry := logging.GetLogger().WithField("context", "GetUserHandler")
	logEntry.Info("Processing GetUser  request")

	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)

	if err != nil {
		logEntry.WithError(err).Error("Page number not in correct format")
		errres := response.ClientResponse(http.StatusBadRequest, "page number not in correct foramt", nil, err.Error())
		c.JSON(http.StatusBadRequest, errres)
		return

	}
	fmt.Println("ssssssssssssssssss")
	user, errs := ad.GRPC_Client.GetUser(page)
	if errs != nil {
		logEntry.WithError(err).Error("Error during GetUser rpcCall")
		errres := response.ClientResponse(http.StatusBadRequest, "error in getting the details", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}
	logEntry.Info("Get user Successfull")
	successres := response.ClientResponse(http.StatusOK, "succesfully got all user details", user, nil)
	c.JSON(http.StatusOK, successres)

}

// @Summary		Block User
// @Description	Blocks a user by ID
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"User ID to block"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admins/block [PATCH]
func (ad *AuthHandler) BlockUser(c *gin.Context) {

	logEntry := logging.GetLogger().WithField("context", "BlockUserHandler")
	logEntry.Info("Processing BlockUser  request")
	id := c.Query("id")

	err := ad.GRPC_Client.BlockUser(id)
	if err != nil {
		logEntry.WithError(err).Error("Error during Blockuser rpcCall")
		errs := response.ClientResponse(http.StatusBadRequest, "error in block user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("Block user Successfull")
	succesres := response.ClientResponse(http.StatusOK, "successfully blocked", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary		Unblock User
// @Description	Unblocks a user by ID
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"User ID to unblock"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admins/unblock [PATCH]
func (ad *AuthHandler) UnBlockUser(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "BlockUserHandler")
	logEntry.Info("Processing UnblockUser  request")

	id := c.Query("id")

	err := ad.GRPC_Client.UnBlockUser(id)
	if err != nil {
		logEntry.WithError(err).Error("Error during UnBlockuser rpcCall")
		errs := response.ClientResponse(http.StatusBadRequest, "error in unblock user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("UnBlock user Successfull")
	succesres := response.ClientResponse(http.StatusOK, "successfully unblocked", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary		Change Password
// @Description	Changes the password of the logged-in user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	header	int	true	"User ID"
// @Param			password	body		models.ChangePassword	true	"New password details"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/changepassword [PATCH]
func (ad *AuthHandler) ChangePassword(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "ChangePasswordUserHandler")
	logEntry.Info("Processing ChangePassword  request")

	idstring, _ := c.Get("id")

	id, _ := idstring.(int)

	var ChangePassword models.ChangePassword

	if err := c.BindJSON(&ChangePassword); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "details are not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	fmt.Println("changed passwords details are", ChangePassword)
	fmt.Println("id for changepasword is ", id)

	if err := ad.GRPC_Client.ChangePassword(id, ChangePassword); err != nil {
		logEntry.WithError(err).Error("Error during ChangePassword rpcCall")
		errs := response.ClientResponse(http.StatusBadRequest, "error in changepassword", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("Changepassword user Successfull")
	succesres := response.ClientResponse(http.StatusOK, "successfully changed the password of you", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

// @Summary		Get your Details
// @Description	Retrieves details of the logged-in user
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/getyoudetails [GET]
func (ad *AuthHandler) GetUserDetails(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetUserDetailsHandler")
	logEntry.Info("Processing GetUserDetails  request")

	id_string, _ := c.Get("id")
	id, _ := id_string.(int)

	userdetails, err := ad.GRPC_Client.GetUserDetails(id)
	if err != nil {

		logEntry.WithError(err).Error("Error during ChangePassword rpcCall")
		errorres := response.ClientResponse(http.StatusBadRequest, "failed to get the userdetails", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}
	logEntry.Info("Changepassword user Successfull")
	succesres := response.ClientResponse(http.StatusOK, "Successfully got the details of you", userdetails, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary		User OTP Login
// @Description	Generates OTP for user login
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			userDetails	body		models.UserOTPLogin	true	"User OTP login details"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/user/otplogin [POST]
func (uh *AuthHandler) UserOTPLogin(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UserOTPLoginHandler")
	logEntry.Info("Processing UserOTPLogin  request")

	var userDetails models.UserOTPLogin
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	otp, err := uh.GRPC_Client.UserOTPLogin(userDetails.Email)
	if err != nil {
		logEntry.WithError(err).Error("Error during UserOTPLogin rpcCall")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to generate OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logEntry.Info("UserOTPLogin user Successfull")
	success := response.ClientResponse(http.StatusOK, "OTP generated successfully", map[string]string{"OTP": otp}, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Send Follow Request
// @Description	Sends a follow request to another user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	header	int	true	"Logged-in User ID"
// @Param			id	query	int	true	"User ID to send follow request to"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/followreq [POST]
func (ad *AuthHandler) FollowReq(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "FollowReqHandler")
	logEntry.Info("Processing Follow  request")

	id_string, _ := c.Get("id")
	id := id_string.(int)

	userId := c.Query("id")
	userID, _ := strconv.Atoi(userId)

	if id == userID {
		logEntry.Info("User attempted to follow themself")
		errs := response.ClientResponse(http.StatusBadRequest, "not allowed to send request yourself", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return

	}

	err := ad.GRPC_Client.FollowReq(id, userID)
	if err != nil {
		logEntry.WithError(err).Error("Error sending follow request")
		errores := response.ClientResponse(http.StatusBadRequest, "could not send the request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	logEntry.Info("Successfully sent follow request")
	succesres := response.ClientResponse(http.StatusOK, "successfully send the requset ", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary		Accept Follow Request
// @Description	Accepts a follow request from another user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	header	int	true	"Logged-in User ID"
// @Param			id	query	int	true	"User ID to accept follow request from"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/acceptfollowreq [POST]
func (ad *AuthHandler) AcceptFollowreq(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "AcceptFollowReqHandler")
	logEntry.Info("Processing AcceptFollow  request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	userId := c.Query("id")
	userID, _ := strconv.Atoi(userId)
	err := ad.GRPC_Client.AcceptFollowreq(id, userID)
	if err != nil {
		logEntry.WithError(err).Error("Error Accepting follow request")
		errores := response.ClientResponse(http.StatusBadRequest, "accepting the request is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	logEntry.Info("Successfully Accepted follow request")
	succesres := response.ClientResponse(http.StatusOK, "succesfully accepted the request", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary		Unfollow User
// @Description	Unfollows a user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	header	int	true	"Logged-in User ID"
// @Param			id	query	int	true	"User ID to unfollow"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/unfollow [POST]
func (ad *AuthHandler) Unfollow(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UnfollowRequestHandler	")
	logEntry.Info("Processing AcceptFollow  request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	useId := c.Query("id")
	userID, _ := strconv.Atoi(useId)

	err := ad.GRPC_Client.Unfollow(id, userID)

	if err != nil {
		logEntry.WithError(err).Error("Error  unfollow request")
		errores := response.ClientResponse(http.StatusBadRequest, "unfollowing is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return

	}
	logEntry.Info("Successfully Unfollowed")
	succesres := response.ClientResponse(http.StatusOK, "successfully unfollowed the follower", nil, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary		Get Followers
// @Description	Retrieves the followers of the logged-in user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	header	int	true	"Logged-in User ID"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/followers [GET]
func (ad *AuthHandler) Followers(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "FollowersHandler	")
	logEntry.Info("Processing show Followers  request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	details, err := ad.GRPC_Client.Followers(id)
	if err != nil {
		logEntry.WithError(err).Error("Error  getting Followers request")
		errores := response.ClientResponse(http.StatusBadRequest, "get the followers failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	logEntry.Info("Successfully get Followers")
	successres := response.ClientResponse(http.StatusOK, "successfully got followers", details, nil)
	c.JSON(http.StatusOK, successres)
}

// @Summary		Get Followings
// @Description	Retrieves the users that the logged-in user is following
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	header	int	true	"Logged-in User ID"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/users/followings [GET]
func (ad *AuthHandler) Followings(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "FollowersHandler	")
	logEntry.Info("Processing show Followers  request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	details, err := ad.GRPC_Client.Followings(id)
	if err != nil {
		logEntry.WithError(err).Error("Error  getting Followings request")
		errores := response.ClientResponse(http.StatusBadRequest, "get the followings failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}
	logEntry.Info("Successfully get Followers")
	successres := response.ClientResponse(http.StatusOK, "successfully got followings", details, nil)
	c.JSON(http.StatusOK, successres)
}

// @Summary		Send OTP
// @Description	Sends OTP to the provided phone number
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			phoneNumber	body		models.OTPData	true	"Phone number to send OTP to"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Failure		500			{object}	response.Response{}
// @Router			/users/sendotp [POST]
func (ad *AuthHandler) SendOTP(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "SendOTPHandler")
	logEntry.Info("Processing SendOTP request")
	var phone models.OTPData

	if err := c.BindJSON(&phone); err != nil {
		logEntry.WithError(err).Error("Error in binding")
		errorres := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
	}

	err := ad.GRPC_Client.SendOTP(phone.PhoneNumber)
	if err != nil {
		logEntry.WithError(err).Error("Error sending OTP")
		errores := response.ClientResponse(http.StatusBadRequest, "seding otp is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
	}
	logEntry.Info("Successfully Send OTP")
	succesres := response.ClientResponse(http.StatusOK, "successfully send OTP", nil, nil)
	c.JSON(http.StatusOK, succesres)

}

// @Summary		Verify OTP
// @Description	Verifies the OTP provided by the user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			code	body		models.VerifyData	true	"Verification code received via OTP"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Failure		500			{object}	response.Response{}
// @Router			/users/verifyotp [POST]
func (ad *AuthHandler) VerifyOTP(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "VerifyOTPHandler")
	logEntry.Info("Processing VerifyOTP request")
	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		logEntry.WithError(err).Error("Error in binding")
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.GRPC_Client.VerifyOTP(code)
	if err != nil {
		logEntry.WithError(err).Error("Error Verify OTP")
		errs := response.ClientResponse(http.StatusBadRequest, "verifying the otp is failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry.Info("Successfully Verify OTP")
	successres := response.ClientResponse(http.StatusOK, "successfully verified ", users, nil)
	c.JSON(http.StatusOK, successres)

}

// @Summary		Upload Profile Picture
// @Description	Uploads profile picture for the logged-in user
// @Tags			User
// @Accept			multipart/form-data
// @Produce		json
// @Param			id	header	int	true	"Logged-in User ID"
// @Param			files	formData	file	true	"Image file to upload"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Failure		500			{object}	response.Response{}
// @Router			/users/profilepic [POST]
func (ad *AuthHandler) UploadProfilepic(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UpdateProfilePicHandler")
	logEntry.Info("Processing UpdateProfilepic request")
	id_string, _ := c.Get("id")
	id := id_string.(int)

	form, err := c.MultipartForm()
	if err != nil {
		logEntry.WithError(err).Error("Error retrieving from form")
		errorRes := response.ClientResponse(http.StatusBadRequest, "retreiving images from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		logEntry.Info("files are empty")
		errorRes := response.ClientResponse(http.StatusBadRequest, "no files are provided", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	for _, file := range files {
		err := ad.GRPC_Client.UploadProfilepic(id, file)
		if err != nil {
			logEntry.WithError(err).Error("Error UploadProfilepic")
			errorRes := response.ClientResponse(http.StatusBadRequest, "could not change one or more images", nil, err.Error())
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}
	}
	logEntry.Info("successfully uploaded the picture")
	successres := response.ClientResponse(http.StatusOK, "uploaded successfully ", nil, nil)
	c.JSON(http.StatusOK, successres)

}
