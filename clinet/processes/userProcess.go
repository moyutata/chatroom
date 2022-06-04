package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/clinet/utils"
	"go_code/chatroom/common/message"
	"net"
)

type UserProcess struct {
}

//

func (userp *UserProcess) Login(userId int, userPwd string) (err error) {
	//制定规则

	//1. 连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送信息
	var mes message.Message
	mes.Type = message.LoginMesType

	//3. 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5. 把data赋给mes.Data字段
	mes.Data = string(data)
	//6. 把mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7. 此时data使我们要发送的类型
	//7.1 首先发送data字节数
	//先获取data的长度-> 表示长度的byte切片
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], pkgLen)
	n, err := conn.Write(buf[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	fmt.Printf("clinet send mes len=%d data=%s\n", len(data), data)

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20秒....")

	//处理服务器返回消息
	//传建一个Transfer实例
	trans := &utils.Transfer{
		Conn: conn,
	}
	mes, err = trans.ReadPkg()

	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}
	//将mes 的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//启动一个协程保持和服务器端保持通讯
		//如果服务器有数据推送给客户端，则接受并显示在客户端终端

		go ServerProcessMes(conn)
		//1. 显示登录成功后的菜单
		ShowMenu()
		//2.
	} else {
		fmt.Println("错误代码:", loginResMes.Code)
		fmt.Println(loginResMes.Error)
	}
	return
}

func (userp *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	//1. 连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送信息
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3. 创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5. 把data赋给mes.Data字段
	mes.Data = string(data)
	//6. 把mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//传建一个Transfer实例
	trans := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端

	err = trans.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	mes, err = trans.ReadPkg() // mes 就是 registerMes

	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	//将mes 的data部分反序列化成LoginResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功~")
	} else {
		fmt.Println("错误代码:", registerResMes.Code)
		fmt.Println(registerResMes.Error)
	}
	return
}
