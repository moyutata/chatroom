package main

import (
	"fmt"
	"go_code/chatroom/client/processes"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	//接收操作指令
	var op int
	//循环退出标志
	loop := true

	//循环显示主菜单
	for loop {
		fmt.Println("-------------------欢迎使用多人聊天室-------------------")
		fmt.Println("\t\t1. 登录聊天系统")
		fmt.Println("\t\t2. 注册用户")
		fmt.Println("\t\t3. 退出系统")
		fmt.Print("\t\t请选择(1-3): ")
		fmt.Scanf("%d\n", &op)
		fmt.Println("-------------------------------------------------------")

		switch op {
		case 1:
			fmt.Println("登录聊天系统")
			loginMenu()
			// loop = false
		case 2:
			fmt.Println("注册用户")
			registerMenu()
			// loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("您的输入有误, 请重新输入~")
		}
	}
}

func loginMenu() {

	fmt.Print("请输入用户id: ")
	fmt.Scanf("%d\n", &userId)
	fmt.Print("请输入用户密码: ")
	fmt.Scanf("%s\n", &userPwd)

	//1. 创建一个UserProcess结构体
	userp := &processes.UserProcess{}

	//这里会重新调用
	userp.Login(userId, userPwd)

}

func registerMenu() {

	fmt.Print("请输入用户id: ")
	fmt.Scanf("%d\n", &userId)
	fmt.Print("请输入用户密码: ")
	fmt.Scanf("%s\n", &userPwd)
	fmt.Print("请输入用户昵称: ")
	fmt.Scanf("%s\n", &userName)

	userp := &processes.UserProcess{}

	userp.Register(userId, userPwd, userName)
}
