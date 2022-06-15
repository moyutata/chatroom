package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct{}

func (sp *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	//1. 遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发出去

	//取出mes中的content
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.SmsResMesType

	var smsResMes message.SmsResMes
	smsResMes.Content = smsMes.Content
	smsResMes.User = smsMes.User

	//序列化smsResMes
	data, err := json.Marshal(smsResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	resMes.Data = string(data)

	//序列化mes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userManager.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		sp.sendToEach(data, up.Conn)
	}
	return
}

func (sp *SmsProcess) sendToEach(data []byte, conn net.Conn) {

	//创建transfer实例转发消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败, err=", err)
	}
}
