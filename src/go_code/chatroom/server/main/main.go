package main

import (
	"fmt"
	"net"
)

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
