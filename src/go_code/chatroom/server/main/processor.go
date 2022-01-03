package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

// 创建结构体
type Processor struct {
	Conn net.Conn
}

// 消息类型中转站： tcp 连接 分发 不同消息类型
func (proc *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建 UserProcess 实例
		userP := &process.UserProcess{
			Conn: proc.Conn,
		}
		err = userP.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在")
	}
	return
}

func (proc *Processor) InitProcess() (err error) {
	// 读取客户端发送的信息
	for {
		fmt.Println("读取客户端数据ing...👉")
		// 封装读取数据包的函数 readPkg() Message , Err
		// 创建 Transfer 实例，读取
		tf := &utils.Transfer{
			Conn: proc.Conn,
		}
		mes, err2 := tf.ReadPkg()
		if err2 != nil {
			if err2 == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return err2
			} else {
				fmt.Println("readPkg - err2: ", err2)
				return err2
			}
		}
		fmt.Printf("服务端接收到客户端ip： %v 发来的消息：%v \n", proc.Conn.RemoteAddr().String(), mes)
		err = proc.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes - err : ", err)
			return
		}
	}
}
