package mgr

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	G_WsMgr *WsMgr
)

func InitWsMgr() {
	G_WsMgr = &WsMgr{
		Lock:          sync.Mutex{},
		WsConnMapping: make(map[uint64]*websocket.Conn),
	}
}

type WsMgr struct {
	Lock          sync.Mutex
	WsConnMapping map[uint64]*websocket.Conn
}

func (m *WsMgr) Set(userID uint64, wsConn *websocket.Conn) {
	fmt.Println("set ws conn")
	m.Lock.Lock()
	defer m.Lock.Unlock()
	if oldConn, ok := m.get(userID); ok {
		fmt.Println("oldConn", oldConn)
		fmt.Println("wsConn", wsConn)
		if oldConn != wsConn {
			err := oldConn.Close()
			fmt.Println("Close err", err)
		}
	}
	m.WsConnMapping[userID] = wsConn
}

func (m *WsMgr) Get(userID uint64) (conn *websocket.Conn, ok bool) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	return m.get(userID)
}

func (m *WsMgr) Del(userID uint64, oldConn *websocket.Conn) {
	fmt.Println("del ws conn")
	m.Lock.Lock()
	defer m.Lock.Unlock()
	if conn, ok := m.get(userID); ok {
		if conn == oldConn {
			oldConn.Close()
			delete(m.WsConnMapping, userID)
		}
	}
}

func (m *WsMgr) get(userID uint64) (conn *websocket.Conn, ok bool) {
	conn, ok = m.WsConnMapping[userID]
	return
}
