package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()

	//调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.ServerProcess()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 err=", err)
	}
}

func initUserDao() {
	//initPool在initUserDao之前
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//当服务器启动时，初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {

	//提示信息
	fmt.Println("服务器ver alpha 1.5 在8889端口监听...")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//延时关闭listener
	defer listen.Close()

	for {
		fmt.Println("wait for connnection...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//连接成功，启动协程与客户端保持通讯
		go process(conn)
	}

}
