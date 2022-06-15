package model

import (
	"go_code/chatroom/common/message"
	"net"
)

//客户端会频繁使用
type CurrentUser struct {
	Conn net.Conn
	message.User
}
