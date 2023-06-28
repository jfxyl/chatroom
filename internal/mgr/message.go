package mgr

import (
	"chatroom/app/models"
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
	G_MessageMgr = &MessageMgr{}
	go G_MessageMgr.Pull()
}

type MessageMgr struct {
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
		msgByte []byte
	)
	fmt.Println("Pull")
	mq.G_PushConsumer.Subscribe("message", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			if json.Unmarshal(msg.Body, &message) != nil {
				continue
			}
			fmt.Printf("message1：%+v", message)
			if db.G_DB.Model(&message).Association("Sender").Find(&message.Sender) != nil {
				continue
			}
			fmt.Printf("message2：%+v", message)

			err = m.SendMessage(message)
			if err != nil {
				fmt.Println("err2", err)
			}
			fmt.Println("message3：", string(msgByte))
		}
		return consumer.ConsumeSuccess, nil
	})
	if err = mq.G_PushConsumer.Start(); err != nil {
		fmt.Println(err)
	}

	defer mq.G_PushConsumer.Shutdown()
	for {
		time.Sleep(1 * time.Second)
	}
}

func (m *MessageMgr) SendMessage(msg models.Message) (err error) {
	var (
		ok      bool
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
			if conn, ok = G_WsMgr.get(user.UserID); ok {
				//conn.SetWriteDeadline(time.Now().Add(writeWait))
				//if !ok {
				//	conn.WriteMessage(websocket.CloseMessage, []byte{})
				//	return
				//}
				conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
	return
}

func (m *MessageMgr) SendReadInfo(msg models.Message) (err error) {
	var (
		ok      bool
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
			if conn, ok = G_WsMgr.get(user.UserID); ok {
				//conn.SetWriteDeadline(time.Now().Add(writeWait))
				//if !ok {
				//	conn.WriteMessage(websocket.CloseMessage, []byte{})
				//	return
				//}
				conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
	return
}
