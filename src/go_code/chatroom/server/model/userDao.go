// data access object
package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"

	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

// 定义 UserDao 结构体， 完成对 User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建 UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1. 根据用户 id 返回一个 User 实例
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过 id 去redis 查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEXISTS
			fmt.Printf("该用户: %v 不存在\n", id)
			return
		}
	}
	fmt.Println("res = ", res)
	// 反序列化 res
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	return
}

// 2. 完成登录校验
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 3. 完成注册逻辑
func (ud *UserDao) Register(user *message.User) (err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	_, err = ud.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data, err2 := json.Marshal(user)
	if err2 != nil {
		fmt.Println("序列化失败", err2)
		return
	}
	// 落库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("HSet 失败", err)
	}
	return
}
