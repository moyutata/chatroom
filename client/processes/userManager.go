package processes

import (
	"fmt"
	"go_code/chatroom/client/model"
	"go_code/chatroom/common/message"
)

//客户端维护的map

var onlineUsers map[int]*message.User = make(map[int]*message.User, 1024)

//用户登录成功后对CurrentUser初始化
var CurrentUser model.CurrentUser

//客户端显示当前在线
func showOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

//处理NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.UserStatus
	onlineUsers[notifyUserStatusMes.UserId] = user

	showOnlineUser()
}
