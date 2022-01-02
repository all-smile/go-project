package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"io"
	"net"
)

func main() {
	fmt.Println("服务器监听 8889 端口")
	ln, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("服务器监听端口出错 = ", err)
		return
	}
	defer ln.Close()
	fmt.Println("ln = ", ln)
	for {
		fmt.Println("等待客户端连接")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ln.Accept() 出错了", err)
			continue
		}
		go handleConnection(conn)
	}
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		// info := fmt.Sprintf("read pkg header err from %v", conn.RemoteAddr().String())
		// err = errors.New(info)
		// fmt.Printf("服务端读取客户端： %v 信息失败= %v", conn.RemoteAddr().String(), err)
		return
	}
	fmt.Println("读取的buf: ", buf)
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[0:4])
	fmt.Printf("pkgLen 类型： %T, 值： %v \n", pkgLen, pkgLen)
	// 从套接字里面读取 pkgLen 字节， 放到 buf 里面去
	n1, err := conn.Read(buf[:pkgLen])
	if uint32(n1) != pkgLen || err != nil {
		err = errors.New("read pkg body err")
		return
	}
	// 反序列化 func Unmarshal(data []byte, v interface{}) error
	// Unmarshal函数解析json编码的数据并将结果存入v指向的值。
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// 1. 先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte // 数组
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 3.1 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}
	// 发送 data 本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送 data 数据失败", err)
		return
	}
	return
}

// 登录逻辑处理
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	// 定义返回消息类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	// 校验信息
	if loginMes.UserId == 100 && loginMes.UserPwd == "abc" {
		loginResMes.Code = 200
		loginResMes.Error = "登录成功"
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
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
	err = writePkg(conn, data)
	return
}

// 消息类型中转站： tcp 连接 分发 不同消息类型
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在")
	}
	return
}

func handleConnection(conn net.Conn) {
	fmt.Println("conn = ", conn)
	defer conn.Close()
	// 读取客户端发送的信息
	for {
		fmt.Println("读取客户端数据ing...")
		// 封装读取数据包的函数 readPkg() Message , Err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			} else {
				fmt.Println("readPkg - err: ", err)
				return
			}
		}
		fmt.Printf("服务端接收到客户端ip： %v 发来的消息：%v \n", conn.RemoteAddr().String(), mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes - err : ", err)
			return
		}
	}
}
