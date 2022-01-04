package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	// 增加字段，表示该连接属于哪个用户
	UserId int
}

// 登录逻辑处理
func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("loginMes反序列化失败", err)
		return
	}
	// 定义返回消息类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	// 校验信息
	// redis数据库校验
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		// 组织错误信息
		if err == model.ERROR_USER_NOEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		// 放入在线用户
		up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(up)

		// 通知在线用户，我已上线
		up.NotifyOthersOnlineUser(loginMes.UserId, 0)

		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Printf("用户id: %v, 登录成功 \n", user.UserId)
	}

	// 序列化信息
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	// 发送
	/*
		MVC 分层
		创建 Transfer 实例， 进行读取
	*/
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 注册
func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}

	// 定义返回消息类型
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// 操作 redis
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		// 组织错误信息
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
		fmt.Printf("用户id: %v, 注册成功 \n", registerMes.User.UserId)
	}

	// 序列化信息
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	// 处理结果发送给客户端
	/*
		MVC 分层
		创建 Transfer 实例， 进行读取
	*/
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 退出登录
func (up *UserProcess) ServerProcessLogout(mes *message.Message) (err error) {
	fmt.Printf("服务端接收到客户端： %v 的退出登录请求 \n", up.Conn.RemoteAddr().String())
	// 反序列化，取出信息
	var logoutMes message.LogoutMes
	err = json.Unmarshal([]byte(mes.Data), &logoutMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}

	// 定义返回消息类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var logoutResMes message.LogoutResMes

	// 调用退出操作
	err = model.MyUserDao.Logout(logoutMes.UserId)
	if err != nil {
		// 组织错误信息
		if err == model.ERROR_USER_NOEXISTS {
			logoutResMes.Code = 500
			logoutResMes.Error = err.Error()
		} else {
			logoutResMes.Code = 505
			logoutResMes.Error = "服务器内部错误"
		}
	} else {
		logoutResMes.Code = 200
		// 下线用户
		userMgr.DelOnlineUser(logoutMes.UserId)

		logoutResMes.NotifyUserStatusMes.UserId = logoutMes.UserId
		logoutResMes.NotifyUserStatusMes.Status = message.UserOfline

		delete(userMgr.onlineUsers, logoutMes.UserId)
		// 通知在线用户，我已下线
		up.NotifyOthersOnlineUser(logoutMes.UserId, logoutMes.Status)
		for id, _ := range userMgr.onlineUsers {
			logoutResMes.UserIds = append(logoutResMes.UserIds, id)
		}
		info := fmt.Sprintf("用户id: %v, 退出成功 \n", logoutMes.UserId)
		fmt.Println(info)
		logoutResMes.Error = info
	}
	// 序列化信息
	data, err := json.Marshal(logoutResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resMes.Data = string(data)
	fmt.Println("服务端退出请求处理结果： ", resMes)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	// 发送
	/*
		MVC 分层
		创建 Transfer 实例， 进行读取
	*/
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 通知其它在线用户我上下线的方法
func (up *UserProcess) NotifyOthersOnlineUser(userId int, status int) {
	// 遍历 onlineUsers, 通知
	fmt.Println("status = ", status)
	fmt.Printf("🎑%v 将要通知在线的用户: %v \n", userId, userMgr.onlineUsers)
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		// 开始通知
		err := up.NotifyOnline(userId, status)
		if err != nil {
			fmt.Println("通知失败", err)
		}
	}
}

func (up *UserProcess) NotifyOnline(userId int, status int) (err error) {
	// 组装 NotifyUserStatusMes 消息类型
	var resMes message.Message
	resMes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusResMes message.NotifyUserStatusMes
	notifyUserStatusResMes.UserId = userId
	notifyUserStatusResMes.Status = status

	// 序列化信息
	data, err := json.Marshal(notifyUserStatusResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	// 发送给客户端
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}
