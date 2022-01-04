package process

import (
	"fmt"
	"go_code/chatroom/client/model"
	"go_code/chatroom/common/message"
)

// 客户端要维护的 map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 登录成功后 初始化 CurUser
var CurUser model.CurUser

// 处理返回的 NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		// 创建 实例
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	if notifyUserStatusMes.Status == 1 {
		// 不在线
		delete(onlineUsers, notifyUserStatusMes.UserId)
	} else {
		// 更新状态
		user.UserStatus = notifyUserStatusMes.Status
		onlineUsers[notifyUserStatusMes.UserId] = user
	}
	outputOnlineUser()
}

// 客户端显示在线用户
func outputOnlineUser() {
	fmt.Println("--------------- 当前在线用户列表: ---------------")
	for id, _ := range onlineUsers {
		fmt.Printf("用户id: \t%v \n", id)
	}
}
