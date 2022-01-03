package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	// 发序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	info := fmt.Sprintf("用户id: %v , 对大家说： %v", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
