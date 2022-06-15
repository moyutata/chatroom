package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SmsResMesType           = "SmsResMes"
)

//用户状态的常量
const (
	UserOffline = iota
	UserOnline
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//定义两个消息, 后面按需添加

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码 200表示登录成功 500表示用户未注册
	Users []int  `json:"users"` //增加字段，保存用户id的切片
	Error string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码 400表示该用户已被占用 200表示用户注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化消息
type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"status"`
}

//增加一个smsMes结构体
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体
}

type SmsResMes struct {
	User    User   `json:"user"`
	Content string `json:"content"`
}
