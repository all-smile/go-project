package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "172.15.37.225:8888")
	if err != nil {
		// handle error
		fmt.Println("client connect err", err)
		return
	}
	fmt.Println("连接成功", conn)

	// 客户端发送数据给服务端
	reader := bufio.NewReader(os.Stdin) // 从标准输入【终端】拿到信息
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("文件读取错误", err)
			return
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		fmt.Println("line = ", line)
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write() err = ", err)
			return
		}
		fmt.Printf("客户端发送 %v 个字节 \n", n)
	}
}

/*
连接成功 &{{0xc00008ea00}}
xixi
line =  xixi
客户端发送 4 个字节
*/
