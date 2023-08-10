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
		WsConnMapping: make(map[uint64][]*websocket.Conn),
	}
}

type WsMgr struct {
	Lock          sync.Mutex
	WsConnMapping map[uint64][]*websocket.Conn
}

func (m *WsMgr) Set(userID uint64, wsConn *websocket.Conn) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	if conns, ok := m.get(userID); ok {

		fmt.Println("wsConn", wsConn)
		fmt.Println("conns", conns)
		for _, conn := range conns {
			fmt.Println("conn == wsConn", conn == wsConn)
			fmt.Println("wsConn.RemoteAddr()", wsConn.RemoteAddr())
			fmt.Println("conn.RemoteAddr()", conn.RemoteAddr())
			if conn.RemoteAddr() == wsConn.RemoteAddr() {
				return
			}
		}
	}
	m.WsConnMapping[userID] = append(m.WsConnMapping[userID], wsConn)
}

func (m *WsMgr) Get(userID uint64) (conns []*websocket.Conn, ok bool) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	return m.get(userID)
}

func (m *WsMgr) Del(userID uint64, oldConn *websocket.Conn) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	fmt.Println("delconn", userID)
	conns, ok := m.get(userID)
	if ok {
		for i, conn := range conns {
			if conn.RemoteAddr() == oldConn.RemoteAddr() {
				oldConn.Close()
				conns = append(conns[:i], conns[i+1:]...)
			}
		}
	}
	if len(conns) == 0 {
		delete(m.WsConnMapping, userID)
	} else {
		m.WsConnMapping[userID] = conns
	}
}

func (m *WsMgr) get(userID uint64) (conns []*websocket.Conn, ok bool) {
	conns, ok = m.WsConnMapping[userID]
	return
}
