package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

const (
	USER_KEY = "users"
)

//服务器启动后，初始化一个userDao实例
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体
//完成对User结构体的所有操作

type UserDao struct {
	pool *redis.Pool
}

//Factory
func NewUserDao(pool *redis.Pool) (userd *UserDao) {
	userd = &UserDao{
		pool: pool,
	}
	return
}

//应该提供方法
//1. 根据userId 返回一个User和err
func (userd *UserDao) getUserById(conn redis.Conn, id int) (user User, err error) {
	//通过给定的id在Redis中查询用户
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在users未找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	//反序列化res -> User
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return

}

//2. 登录校验
//a. 用户的id和pwd都正确，则返回一个user实例
//b. 用户的id或pwd有错误，则返回对应的错误信息
func (userd *UserDao) Login(userId int, userPwd string) (user User, err error) {

	//1. 从redis连接池中取出一个连接
	conn := userd.pool.Get()
	defer conn.Close()
	user, err = userd.getUserById(conn, userId)
	if err != nil {
		return
	}

	//这时获取到该用户
	if user.UserPwd != userPwd {
		err = ERROR_PWD
		return
	}
	return
}

func (userd *UserDao) Register(user *message.User) (err error) {
	//1. 从redis连接池中取出一个连接
	conn := userd.pool.Get()
	defer conn.Close()
	_, err = userd.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//这时注册用户id不存在于数据库
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//存入数据库
	_, err = conn.Do("HSET", USER_KEY, user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误~ err=", err)
	}
	return
}
