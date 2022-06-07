package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//应该有哪些字段
	Conn net.Conn
	//增加一个字段, 表示该Conn是属于哪个用户的
	UserId int
}

//通知所有在线用户userId上线
func (userp *UserProcess) NotifyOtherOnlineUser(userId int) {
	//遍历onlineUsers, 发送
	for id, up := range userp.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (userp *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.UserOnline

	//序列化notifyUserStatusMes
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//data赋值给mes
	mes.Data = string(data)

	//序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送
	trans := &utils.Transfer{
		Conn: userp.Conn,
	}

	err = trans.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (userp *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmashal() err=", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 504
			registerResMes.Error = "Unknown Error!"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功~")
	}

	//3. 将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4. 将data赋值给resMes
	resMes.Data = string(data)

	//5. 将resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//6. 发送data
	//使用了mvc模式, 先创建一个Transfer实例
	trans := &utils.Transfer{
		Conn: userp.Conn,
	}
	err = trans.WritePkg(data)
	return

}

//编写一个函数ServerProcessLogin，专门处理登录请求
func (userp *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data, 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明一个loginResMes
	var loginResMes message.LoginResMes

	//登录验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 404
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		//此时用户登录成功，将登录成功的用户放入到userManager中
		//将登录成功的userId存入到userp
		userp.UserId = loginMes.UserId
		userManager.AddOnlineUser(userp)
		//通知其他在线的用户上线
		userp.NotifyOtherOnlineUser(loginMes.UserId)
		//将当前在线用户的id放入loginResMes中
		for id := range userManager.onlineUsers {
			loginResMes.Users = append(loginResMes.Users, id)
		}

		fmt.Println(user.UserName, "登录成功")
	}

	//2. 如果用户id=100，密码=123456, 则合法，否则不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册后再使用~"
	// }

	//3. 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//4. 将data赋值给resMes
	resMes.Data = string(data)

	//5. 将resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//6. 发送data
	//使用了mvc模式, 先创建一个Transfer实例
	trans := &utils.Transfer{
		Conn: userp.Conn,
	}
	err = trans.WritePkg(data)
	return
}
