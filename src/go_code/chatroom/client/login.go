package main

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// 定协议
	/* fmt.Printf("userId = %d, userPwd = %s \n", userId, userPwd)
	return nil */
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接失败： ", err)
		return
	}
	// 延时关闭
	defer conn.Close()
	fmt.Println("conn = ", conn)
	// 2. 组织消息
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 2.1. 序列化 loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.Data = string(data)
	// 2.1. 序列化 mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	// 调用封装的 writePkg() 函数， 向服务端发送数据
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("客户端发送消息失败", err)
		return
	}
	/* // 3. 发送
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte // 数组
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 3.1 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}
	fmt.Printf("客户端发送消息长度: %v, 内容： %s \n", len(data), string(data))

	// 4. 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write 失败=", err)
		return
	} */

	/* // 休眠 5 秒
	time.Sleep(5 * time.Second)
	fmt.Println("这里休眠了 5 秒") */

	// 5. 处理服务端返回
	mes, err = readPkg(conn)
	fmt.Printf("mes 类型： %T, 值： %v \n", mes, mes)
	if err != nil {
		fmt.Println("服务器处理失败")
		return
	}
	// 反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
