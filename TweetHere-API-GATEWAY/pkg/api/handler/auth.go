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

func (ad *AuthHandler) AdminSignUp(c *gin.Context) {
	var adminDetails models.AdminSignup

	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate admin0 ", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Admin created successfully", admin, nil)
	c.JSON(http.StatusOK, success)
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

	if err := ad.GRPC_Client.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.NewPassword, ChangePassword.RePassword); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "error in changepassword", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully changed the password of you", nil, nil)
	c.JSON(http.StatusOK, succesres)

}
