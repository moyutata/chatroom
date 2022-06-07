package message

//定义一个用户的结构体

type User struct {
	//需要给字段加上tag
	//保证序列化和反序列化成功
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
}
