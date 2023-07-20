package services

import (
	"bytes"
	"chatroom/internal/common"
	"chatroom/internal/mgr"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
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

func NewWsService() *WsService {
	return &WsService{
		MessageService: NewMessageService(),
	}
}

type WsService struct {
	MessageService *MessageService
}

func (o *WsService) ReadAndWrite(userId uint64, conn *websocket.Conn) {
	mgr.G_WsMgr.Set(userId, conn)
	go o.Read(userId, conn)
	go o.Write(userId, conn)
}

func (o *WsService) Read(userId uint64, conn *websocket.Conn) {
	var (
		err         error
		messageByte []byte
		unix        = time.Now().Unix()
	)
	fmt.Println("-----------------------------")
	fmt.Println("unix", unix)
	fmt.Println("userId", userId)
	fmt.Printf("conn = %p\n", conn)
	defer func() {
		fmt.Println("unix", unix)
		fmt.Println("userId", userId)
		fmt.Printf("conn = %p\n", conn)
	}()
	for {
		var (
			wsRequest common.WsRequest
			readBody  common.ReadBody
			bodyByte  []byte
		)
		if _, messageByte, err = conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err)
			}
			fmt.Println("Push err1", err)
			break
		}
		messageByte = bytes.TrimSpace(bytes.Replace(messageByte, newline, space, -1))

		if err = json.Unmarshal(messageByte, &wsRequest); err == nil {
			fmt.Println("wsRequest", wsRequest)
			switch wsRequest.Action {
			case "read":
				bodyByte, _ = json.Marshal(wsRequest.Body)
				if err = json.Unmarshal(bodyByte, &readBody); err == nil {
					readBody.UserId = userId
					o.MessageService.PushRead(&readBody)
					//o.MessageService.Read(readBody, userId)
					//fmt.Println("readBody", readBody)
				}
				fmt.Println("err", err)
			}
		}
		fmt.Println("Push err2", err)
	}
	fmt.Println("READ 协程结束")
}

func (o *WsService) Write(userId uint64, conn *websocket.Conn) {
	var (
		err         error
		messageByte []byte
		messages    chan []byte = make(chan []byte, 1000)
		ok          bool
		ticket      *time.Ticker
	)
	ticket = time.NewTicker(pingPeriod)
	defer func() {
		ticket.Stop()
		mgr.G_WsMgr.Del(userId, conn)
	}()
	for {
		select {
		case messageByte, ok = <-messages:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err = conn.WriteMessage(websocket.TextMessage, messageByte); err != nil {
				return
			}
		case <-ticket.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err = conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
	fmt.Println("WRITE 协程结束")
}
