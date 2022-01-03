// 发消息
package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	// 1. 创建 message
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.Data = string(data)
	// 序列化 mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	// 调用封装的 writePkg() 函数， 向服务端发送数据
	// 创建 Transfer 实例
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("用户: %v 发送消息失败:%v \n", CurUser.UserId, err)
		return
	}
	return
}
