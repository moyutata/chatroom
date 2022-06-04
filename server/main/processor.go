package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/processes"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

//先创建一个Processor结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，决定有哪个函数来处理
func (processor *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		userp := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = userp.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("该消息类型不存在，无法处理...")
	}
	return
}

func (processor *Processor) ServerProcess() (err error) {

	//读客户端发送的信息
	for {
		//封装函数用于读取数据包
		//创建一个Transfer
		trans := &utils.Transfer{
			Conn: processor.Conn,
		}
		mes, err := trans.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		// fmt.Println("mes=", mes)
		err = processor.ServerProcessMes(&mes)
		if err != nil {
			return err
		}

	}
}
