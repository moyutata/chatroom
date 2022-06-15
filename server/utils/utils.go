package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析应该有哪些字段
	Conn net.Conn
	Buf  [4096]byte
}

func (trans *Transfer) ReadPkg() (mes message.Message, err error) {

	// buf := make([]byte, 4096)
	fmt.Println("reading message from client...")

	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则不会阻塞
	_, err = trans.Conn.Read(trans.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header err~")
		return
	}

	//根据buf[:4]转成uint32类型
	var pkgLen = binary.BigEndian.Uint32(trans.Buf[:4])

	//根据pkgLen 读取消息内容
	n, err := trans.Conn.Read(trans.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body err~")
		return
	}

	//将buf[:pkgLen]反序列化成 -> message.Message
	err = json.Unmarshal(trans.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (trans *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(trans.Buf[:4], pkgLen)
	n, err := trans.Conn.Write(trans.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//发送data本身
	n, err = trans.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}
