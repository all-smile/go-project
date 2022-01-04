// 服务端转发消息
package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

// 服务端转发客户端消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 取出 mes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
	}

	for id, up := range userMgr.onlineUsers {
		// 过滤自己
		if id == smsMes.UserId {
			continue
		}
		err = sp.SendMesToUser(data, up.Conn)
	}
	return
}

// 发送
func (sp *SmsProcess) SendMesToUser(data []byte, conn net.Conn) (err error) {
	// 发送
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	return
}
