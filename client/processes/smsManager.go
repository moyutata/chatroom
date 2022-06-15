package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//mes为SmsResMes类型

	//显示
	//1. 反序列化mes
	var smsResMes message.SmsResMes
	err := json.Unmarshal([]byte(mes.Data), &smsResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	info := fmt.Sprintf("用户id:\t%d,对大家说:\t%s", smsResMes.User.UserId, smsResMes.Content)
	fmt.Println(info)
	fmt.Println()

}
