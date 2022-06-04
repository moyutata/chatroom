package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 4096)
	// fmt.Println("reading message from clinet...")

	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header err~")
		return
	}

	//根据buf[:4]转成uint32类型
	var pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body err~")
		return
	}

	//将buf[:pkgLen]反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// func writePkg(conn net.Conn, data []byte) (err error) {

// 	//先发送一个长度给对方
// 	var pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[:], pkgLen)
// 	n, err := conn.Write(buf[:])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write err=", err)
// 		return
// 	}

// 	//发送data本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write err=", err)
// 		return
// 	}
// 	return
// }
