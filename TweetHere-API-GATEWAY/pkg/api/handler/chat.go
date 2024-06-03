package handler

import (
	interfaces "TweetHere-API/pkg/client/interface"
	"TweetHere-API/pkg/helper"
	"TweetHere-API/pkg/utils/models"
	"TweetHere-API/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

type ChatHandler struct {
	GRPC_Client interfaces.ChatClient
}

func NewChatHandler(chatClient interfaces.ChatClient) *ChatHandler {
	return &ChatHandler{
		GRPC_Client: chatClient,
	}
}

// websocket
func (ch *ChatHandler) FriendMessage(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	UserID, exists := c.Get("id")
	if !exists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, "User ID missing")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	id, ok := UserID.(int)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID format error", nil, "Invalid User ID format")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	fmt.Println("user id isssss", id)

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	defer delete(User, strconv.Itoa(id))
	defer conn.Close()

	user := strconv.Itoa(id)
	User[user] = conn

	for {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		helper.SendMessageToUser(User, msg, user)
	}
}

func (ch *ChatHandler) GetChat(c *gin.Context) {
	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userid, exist := c.Get("id")
	if !exist {
		errs := response.ClientResponse(http.StatusBadRequest, "user id not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID := strconv.Itoa(userid.(int))
	fmt.Println("iddddddd iss", userID)
	fmt.Println("details is ", chatRequest)
	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)

	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	succesres := response.ClientResponse(http.StatusOK, "successfully got all chat details", result, nil)
	c.JSON(http.StatusOK, succesres)
}

// func (ch *ChatHandler) VideoCall(c *gin.Context) {
// 	peerConnectionConfig := webrtc.Configuration{
// 		ICEServers: []webrtc.ICEServer{
// 			{
// 				URLs: []string{"stun:stun.l.google.com:19302"},
// 			},
// 		},
// 	}

// 	peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfig)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}

// 	peerConnection.OnICEConnectionStateChange(func(is webrtc.ICEConnectionState) {
// 		fmt.Printf("connection state has changed %s \n", is.String())
// 	})

// 	offer := webrtc.SessionDescription{}
// 	if err := c.BindJSON(&offer); err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "invalid offer", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}

// 	if err := peerConnection.SetRemoteDescription(offer); err != nil {
// 		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to set remote description", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}

// 	answer, err := peerConnection.CreateAnswer(nil)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}

// 	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)
// 	peerConnection.SetLocalDescription(answer)
// 	<-gatherComplete

// 	c.JSON(http.StatusOK, peerConnection.LocalDescription())
// }
