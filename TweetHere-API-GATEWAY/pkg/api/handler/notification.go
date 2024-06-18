package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/logging"
	"TweetHere-API/pkg/utils/models"
	"TweetHere-API/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	GRPC_Client interfaces.NotificationClient
}

func NewNotificationHandler(notiClient interfaces.NotificationClient) *NotificationHandler {
	return &NotificationHandler{
		GRPC_Client: notiClient,
	}
}

// GetNotification godoc
// @Summary Get Notifications
// @Description Retrieves notifications for the logged-in user with pagination.
// @Tags Notification
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id header int true "Logged-in User ID"
// @Param limit query int false "Limit number of notifications to retrieve"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} response.Response{data=[]models.Notification} "Successfully retrieved notifications"
// @Failure 400 {object} response.Response{} "Invalid request format or JWT claims missing"
// @Failure 500 {object} response.Response{} "Server error occurred"
// @Router /noti [get]
func (ad *NotificationHandler) GetNotification(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetNotificationHandler")
	logEntry.Info("Processing GetNotification request")
	var notificationRequest models.NotificationPagination
	if err := c.ShouldBindJSON(&notificationRequest); err != nil {
		logEntry.WithError(err).Error("error in binding")
		errorres := response.ClientResponse(http.StatusBadRequest, "details give in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	id_string, _ := c.Get("id")
	id, _ := id_string.(int)

	result, errs := ad.GRPC_Client.GetNotification(id, notificationRequest)
	if errs != nil {
		logEntry.WithError(errs).Error("error in GetNotification call")
		errss := response.ClientResponse(http.StatusBadRequest, "error in getting notification", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errss)
		return
	}
	logEntry.Info("getNotification successfull")
	succesres := response.ClientResponse(http.StatusOK, "successfully got all notification", result, nil)
	c.JSON(http.StatusOK, succesres)
}
