// 保持跟服务端的通信
package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("***************** 欢迎XXX登录多人聊天系统 *****************")
	fmt.Println("***************** 1. 显示在线用户列表 *****************")
	fmt.Println("***************** 2. 发送消息 *****************")
	fmt.Println("***************** 3. 信息列表 *****************")
	fmt.Println("***************** 4. 退出系统 *****************")
	var key int
	fmt.Println("\t\t\t 请选择(1-4):")
	fmt.Scanf("%d \n", &key)
	var content string

	// 总会存在 发送消息的 情况， 因此，在switch 外部 创建 SmsProcess 实例
	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		fmt.Println("在线用户列表--")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息--")
		fmt.Println("你想说点什么:")
		fmt.Scanln(&content) // 如果输入的有空格，匹配出错？
		err := smsProcess.SendGroupMes(content)
		if err != nil {
			fmt.Println("发送失败", err)
		}
	case 3:
		fmt.Println("信息列表--")
	case 4:
		fmt.Println("退出系统--")

		os.Exit(0)
	default:
		fmt.Println("输入有误，重新输入")
	}
}

// 监听 服务端消息 并显示
func ServerProcessMes(conn net.Conn) {
	// 创建 Transfer 实例， 等待信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端： %v ,正在监听服务端消息 \n", conn.RemoteAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端读取信息失败---", err)
			return
		}
		fmt.Println("服务端推送的消息 = ", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 有人上线了
			// 1. 取出 NotifyUserStatusMes
			// 2. 客户端保存状态 定义 map[int]User
			// 反序列化
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("反序列化失败", err)
				return
			}
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			fmt.Println("接收消息啦： ", mes)
			outputGroupMes(&mes)
		default:
			fmt.Println("服务端消息类型未知")
		}
	}
}
