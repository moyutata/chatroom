package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
)

type SmsProcess struct{}

//发送群聊消息
func (sp *SmsProcess) SendGroupMes(content string) (err error) {

	//1. 创建一个Message实例
	var mes message.Message
	mes.Type = message.SmsMesType

	//2. 创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurrentUser.UserId
	smsMes.UserStatus = CurrentUser.UserStatus

	//3. 序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	//4. 序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}

	//5. 将mes发送给服务器
	trans := &utils.Transfer{
		Conn: CurrentUser.Conn,
	}
	err = trans.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes trans.WritePkg err =", err)
		return
	}
	return

}
