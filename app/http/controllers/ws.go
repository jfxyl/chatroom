package controllers

import (
	"chatroom/app/http/services"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	newline []byte = []byte{'\n'}

	space []byte = []byte{' '}

	writeWait = 10 * time.Second

	pongWait = 5 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize int64 = 512
)

func NewWsController() *WsController {
	return &WsController{
		MessageService: services.NewMessageService(),
		WsService:      services.NewWsService(),
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

type WsController struct {
	upgrader       *websocket.Upgrader
	WsService      *services.WsService
	MessageService *services.MessageService
}

func (o *WsController) Conn(c *gin.Context) {
	fmt.Println("ws conn")
	var (
		err  error
		conn *websocket.Conn
	)
	if conn, err = o.upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
		return
	}
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	o.WsService.ReadAndWrite(auth.User(c).ID, conn)
}
