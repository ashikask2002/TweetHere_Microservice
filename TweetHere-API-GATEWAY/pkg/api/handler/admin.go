package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/utils/models"
	"TweetHere-API/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(adminClient interfaces.AdminClient)*AdminHandler{
	return &AdminHandler{
		GRPC_Client: adminClient,
	}
}

func (ad *AdminHandler) AdminSignup(c *gin.Context){
	var adminDetails models.AdminSignup
	
	if err := c.ShouldBindJSON(&adminDetails);err != nil{
		errs := response.ClinetResponse(http.StatusBadRequest,"Details are not in correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errs)
		return
	}
	
	admin,err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil{
		errs := response.ClinetResponse(http.StatusInternalServerError,"cannot authenticate admin",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errs)
		return
	}
	
	success := response.ClinetResponse(http.StatusOK,"admin created successfully",admin,nil)
	c.JSON(http.StatusOK,success)
}

func (ad *AdminHandler) LoginHandler(c *gin.Context){
	var adminDetails models.AdminLogin	
	if err := c.ShouldBindJSON(adminDetails);err != nil{
		errs := response.ClinetResponse(http.StatusBadRequest,"details are not in correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errs)
		return
	}

	admin,err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil{
		errs := response.ClinetResponse(http.StatusInternalServerError,"Cannot authenthicate user",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errs)
		return
	}
	success := response.ClinetResponse(http.StatusOK,"successfully logged in",admin,nil)
	c.JSON(http.StatusOK,success)
}