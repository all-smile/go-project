package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// Monster 绑定 Store 方法, 序列化 并 保存到文件中
func (monster *Monster) Store() bool {
	// 1、序列化
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}
	// 2、保存到文件中
	filePath := "d:/monster.txt"
	_err := ioutil.WriteFile(filePath, data, 0666)
	if _err != nil {
		fmt.Println("保存文件失败", _err)
		return false
	}
	return true
}

// Monster 绑定 ReStore 方法, 从文件中读取Json字符串， 并反序列化
func (monster *Monster) ReStore() bool {
	// 1、从文件中读取Json字符串
	filePath := "d:/monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return false
	}
	// 2、反序列化
	error := json.Unmarshal(data, &monster)
	if error != nil {
		fmt.Println("反序列化失败", error)
		return false
	} else {
		fmt.Println("monster=", monster)
	}
	return true
}
