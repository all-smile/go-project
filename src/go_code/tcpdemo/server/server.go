package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听...")
	ln, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败")
		return
	} else {
		fmt.Printf("监听成功, ln类型： %T, 值： %v \n", ln, ln)
		// ln类型： *net.TCPListener, 值： &{0xc00008ea00 {<nil> 0}}
	}
	defer ln.Close()

	// 等待客户端连接
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("n.Accept() err = ", err)
			continue
		} else {
			fmt.Printf("n.Accept() suc  conn: %v, 连接的客户端ip: %v \n", conn, conn.RemoteAddr().String())
			// n.Accept() suc  conn: &{{0xc00010ec80}}, 连接的客户端ip: 172.15.37.225:63367
		}
		// 这里准备一个协程，与客户端交互
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// 循环接收客户端连接
	defer conn.Close()
	for {
		// 创建切片
		buf := make([]byte, 1024)
		// 1.等待客户端通过 conn 发送信息
		// 2.如果客户端没有 发送信息 ，那么就阻塞
		fmt.Printf("服务端在等待客户端： %v 输入信息 \n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		/* if err == io.EOF {
			fmt.Printf("客户端： %v 退出\n", conn.RemoteAddr().String())
			return
		} */
		if err != nil {
			fmt.Println("服务端conn.Reader出错", err)
			return
		}
		// 3.显示服务端接收到的客户端信息
		fmt.Print(string(buf[:n]))
	}
}
