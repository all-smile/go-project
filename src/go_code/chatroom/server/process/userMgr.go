// 统计在线人数 - 增删改查
package process

import "fmt"

// 很多需要使用 UserMgr 的地方，定义全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		// 预置 1024 个客户端连接
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 增加
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

// 删除 - 单个
func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

// 统计在线人数
func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers
}

// 根据id返回用户
func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userId]
	if !ok {
		fmt.Printf("用户： %v 不在线或不存在 \n", userId)
		return
	}
	return
}

// 查询
