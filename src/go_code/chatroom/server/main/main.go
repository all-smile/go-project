package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

// 编写函数，完成UserDao初始化
func initUserDao() {
	// 这里的 pool 是在 redis.go 定义的全局变量
	// 需要注意初始化顺序问题
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	// 初始化 redis 连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {
	fmt.Println("🎁新服务端监听 8889 端口")
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

func handleConnection(conn net.Conn) {
	fmt.Println("conn = ", conn)
	defer conn.Close()
	// 创建 Procssor 实例， 调用总控（转发中心）
	processor := &Processor{
		Conn: conn,
	}
	err := processor.InitProcess()
	if err != nil {
		fmt.Println("服务端处理失败 InitProcess : ", err)
		return
	}
}
