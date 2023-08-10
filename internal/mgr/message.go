package mgr

import (
	"chatroom/app/models"
	"chatroom/internal/common"
	"chatroom/internal/db"
	"chatroom/internal/mq"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gorilla/websocket"
	"time"
)

var (
	G_MessageMgr *MessageMgr
)

func InitMessageMgr() {
	G_MessageMgr = &MessageMgr{
		readChan:       make(chan *common.ReadBody, 1000),
		AutoCommitChan: make(chan *common.ReadBodyBatch, 1000),
	}
	go G_MessageMgr.Pull()
}

type MessageMgr struct {
	readChan       chan *common.ReadBody
	AutoCommitChan chan *common.ReadBodyBatch
}

func (m *MessageMgr) Push(topic string, msg []byte) (err error) {
	var (
		ctx = context.Background()
	)
	_, err = mq.G_Producer.SendSync(ctx, primitive.NewMessage(topic, msg))
	return
}

func (m *MessageMgr) Pull() {
	var (
		err     error
		message models.Message
	)
	fmt.Println("Pull")
	mq.G_PushConsumer.Subscribe("message", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			if json.Unmarshal(msg.Body, &message) != nil {
				continue
			}
			err = m.WriteMessage(&message)
			if err != nil {
				fmt.Println("err2", err)
			}
		}
		return consumer.ConsumeSuccess, nil
	})

	//已读信号积攒满一定数量或者一定时间后，发送已读信号
	mq.G_PushConsumer.Subscribe("read", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			var (
				readBody common.ReadBody
			)
			if err = json.Unmarshal(msg.Body, &readBody); err != nil {
				fmt.Println("err", err)
				continue
			}
			fmt.Printf("readBody：%+v", readBody)
			m.Push2ReadChan(&readBody)
		}
		return consumer.ConsumeSuccess, nil
	})

	go m.ReadDelayLoop()

	if err = mq.G_PushConsumer.Start(); err != nil {
		fmt.Println(err)
	}

	defer mq.G_PushConsumer.Shutdown()
	for {
		time.Sleep(1 * time.Second)
	}
}

func (m *MessageMgr) SendReadInfo(msg models.Message) (err error) {
	var (
		ok      bool
		conns   []*websocket.Conn
		conn    *websocket.Conn
		user    *models.UserRoom
		users   []*models.UserRoom
		message []byte
	)
	message, err = json.Marshal(msg.Transform(0))
	if err != nil {
		return err
	}
	if msg.ChatType == models.GroupChat {
		if db.G_DB.Where(map[string]any{"room_id": msg.ReceiverID}).Find(&users).Error != nil {
			return err
		}
		fmt.Println("users", users)
		for _, user = range users {
			if conns, ok = G_WsMgr.Get(user.UserID); ok {
				for _, conn = range conns {
					conn.WriteMessage(websocket.TextMessage, message)
				}
			}
		}
	}
	return
}

func (m *MessageMgr) Push2ReadChan(readBody *common.ReadBody) {
	m.readChan <- readBody
}

func (m *MessageMgr) ReadDelayLoop() {
	var (
		readBody      *common.ReadBody
		readBodyBatch *common.ReadBodyBatch
		timeOutBatch  *common.ReadBodyBatch
		commitTimer   *time.Timer
	)

	for {
		select {
		case readBody = <-m.readChan:
			fmt.Println("readBody", readBody)
			if readBodyBatch == nil {
				readBodyBatch = &common.ReadBodyBatch{}
				commitTimer = time.AfterFunc(2*time.Second, func(readBodyBatch *common.ReadBodyBatch) func() {
					return func() {
						m.AutoCommitChan <- readBodyBatch
					}
				}(readBodyBatch))
			}
			readBodyBatch.ReadBodyList = append(readBodyBatch.ReadBodyList, readBody)
			if len(readBodyBatch.ReadBodyList) >= 1000 {
				m.HandleRead(readBodyBatch)
				readBodyBatch = nil
				commitTimer.Stop()
			}
		case timeOutBatch = <-m.AutoCommitChan:
			if timeOutBatch != readBodyBatch {
				continue
			}
			m.HandleRead(readBodyBatch)
			readBodyBatch = nil
		}
	}
}

func (m *MessageMgr) HandleRead(readBodyBatch *common.ReadBodyBatch) {
	var (
		err         error
		msgUsersMap = make(map[uint64][]uint64)
		conns       []*websocket.Conn
		conn        *websocket.Conn
		ok          bool
		msgByte     []byte
	)
	fmt.Println("readBodyBatch", readBodyBatch)
	for _, readBody := range readBodyBatch.ReadBodyList {
		fmt.Println("readBody", readBody)
		msgUsersMap[readBody.MessageID] = append(msgUsersMap[readBody.MessageID], readBody.UserId)
	}
	fmt.Println("msgUsersMap", msgUsersMap)
	for messageID, userIDs := range msgUsersMap {
		var (
			message models.Message
		)
		fmt.Println("messageID", messageID)
		if err = db.G_DB.Where("id = ?", messageID).Limit(1).Find(&message).Error; err == nil && message.ID != 0 {
			if err = db.G_DB.Model(&models.MessageRead{}).Where("message_id = ?", messageID).Where("user_id in ?", userIDs).Update("Read", true).Error; err == nil {
				if conns, ok = G_WsMgr.Get(message.SenderID); ok {
					if err = db.G_DB.Model(&message).Association("ReaderInfo").Find(&message.ReaderInfo); err == nil {
						fmt.Println("message_id", message.ID)
						msgByte, _ = json.Marshal(message.Transform(0))
						fmt.Println("msgByte", string(msgByte))
						for _, conn = range conns {
							conn.WriteMessage(websocket.TextMessage, msgByte)
						}
					}
				}
			}
		}
		fmt.Println("err", err)
	}
}

func (m *MessageMgr) WriteMessage(message *models.Message) (err error) {
	var (
		ok          bool
		conns       []*websocket.Conn
		conn        *websocket.Conn
		user        *models.UserRoom
		users       []*models.UserRoom
		messageByte []byte
	)
	if err = db.G_DB.Model(&message).Association("Sender").Find(&message.Sender); err != nil {
		return
	}
	messageByte, err = json.Marshal(message.Transform(0))
	if err != nil {
		return err
	}
	if message.ChatType == models.GroupChat {
		if db.G_DB.Where(map[string]any{"room_id": message.ReceiverID}).Find(&users).Error != nil {
			return err
		}
		fmt.Println("users", users)
		for _, user = range users {
			if conns, ok = G_WsMgr.Get(user.UserID); ok {
				//conn.SetWriteDeadline(time.Now().Add(writeWait))
				//if !ok {
				//	conn.WriteMessage(websocket.CloseMessage, []byte{})
				//	return
				//}
				fmt.Println("conns", conns)
				for _, conn = range conns {
					conn.WriteMessage(websocket.TextMessage, messageByte)
				}
			}
		}
	}
	return
}
