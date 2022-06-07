package processes

import "fmt"

/*
	UserMgr在服务器端仅有一个
	经常被使用，定义为全局变量
*/

var (
	userManager *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userManager = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//对onlineUsers添加
func (userm *UserMgr) AddOnlineUser(userp *UserProcess) {
	userm.onlineUsers[userp.UserId] = userp
}

//对onlineUsers删除
func (userm *UserMgr) DeleteOnlineUser(userId int) {
	delete(userm.onlineUsers, userId)
}

//返回所有在线用户
func (userm *UserMgr) GetAddOnlineUsers() map[int]*UserProcess {
	return userm.onlineUsers
}

//根据id返回对应userProcess
func (userm *UserMgr) GetOnlineUserById(userId int) (userp *UserProcess, err error) {
	userp, ok := userm.onlineUsers[userId]
	if !ok {
		//此时该userId用户不在线
		err = fmt.Errorf("用户%d不在线", userId)
		return
	}
	return
}
