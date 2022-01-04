package message

// 定义消息类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	LogoutMesType           = "LogoutMes"
	LogoutResMesType        = "LogoutResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

type Message struct {
	// 消息类型
	Type string `json:"type"`
	// 消息内容
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	// 状态码： 500-未注册 200-成功 403-密码不正确
	Code int `json:"code"`
	// 具体错误信息
	Error string `json:"error"`
	// 增加返回在线人数字段
	UserIds []int `json:"userIds"`
}

// 注册类型
type RegisterMes struct {
	User User `json:"user"`
	// UserName string
	// UserPwd  string
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// 定义常量， 表示在线状态
const (
	UserOnline = iota
	UserOfline
	UserBusyStatus
)

// 配合服务端推送 用户状态变化 （在离线）
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 退出登录消息类型
type LogoutMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type LogoutResMes struct {
	// 状态码： 500-失败 200-成功 505-未知错误
	Code int `json:"code"`
	// 具体错误信息
	Error string `json:"error"`
	// 增加返回在线人数字段
	UserIds []int `json:"userIds"`
	// 增加 当前用户的状态 字段
	NotifyUserStatusMes `json:"notifyUserStatusMes"`
}

// 增加发送消息的类型
type SmsMes struct {
	Content string `json:"content"`
	// 匿名结构体 继承
	User `json:"user"`
}
