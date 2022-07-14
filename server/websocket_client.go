package server

import (
	"github.com/gorilla/websocket"
	"sync"
)

// DarlingBot 通信驱动websocket_client实现

type WebsocketClient struct {
	Conn        *websocket.Conn // websocket连接
	Mutex       sync.Mutex      // 写锁
	Url         string          // ws连接地址
	AccessToken string          // 访问令牌
	SelfID      int64           // 自身ID
	seq         uint64          // 序列号
}
