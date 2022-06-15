package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu() {
	var op int
	fmt.Print("\n**********************xxxx登录成功**********************\n\n")
	for {
		fmt.Println("-------------------欢迎使用多人聊天室-------------------")
		fmt.Println("\t\t1. 显示在线用户列表")
		fmt.Println("\t\t2. 发送消息")
		fmt.Println("\t\t3. 信息列表")
		fmt.Println("\t\t4. 退出系统")
		fmt.Print("\t\t请选择(1-4): ")
		fmt.Scanf("%d\n", &op)
		fmt.Println("-------------------------------------------------------")

		//smsProcess会被经常使用
		smsProcess := &SmsProcess{}
		switch op {
		case 1:
			showOnlineUser()
		case 2:
			sendMes(smsProcess)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("退出系统...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确，请重新输入...")
		}
	}

}

func ServerProcessMes(Conn net.Conn) (err error) {

	//传建一个Transfer实例, 不停地读取服务器发送的消息
	trans := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Printf("\n客户端 %s 正在读取服务器端发送的消息...\n", Conn.LocalAddr().String())
		mes, err := trans.ReadPkg()
		if err != nil {
			fmt.Println("trans.ReadPkg() err=", err)
			return err
		}
		// fmt.Printf("mes=%v\n", mes)

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//处理
			//1. 取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal err=", err)
			}
			//2. 把用户的状态保存到客户端map
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsResMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("Unknown Type!")
		}
	}
}

func sendMes(sp *SmsProcess) {
	var content string
	fmt.Println("-------------------欢迎使用多人聊天室-------------------")
	fmt.Println("请输入聊天信息:")
	fmt.Scanf("%s\n", &content)
	sp.SendGroupMes(content)
}
