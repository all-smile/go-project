package message

type User struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	// 用户状态 ： 在线 离线
	UserStatus int    `json:"userStatus"`
	Sex        string `json:"sex"`
}
