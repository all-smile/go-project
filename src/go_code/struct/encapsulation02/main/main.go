package main

import (
	"fmt"
	"go_code/struct/encapsulation02/model"
)

func main() {
	account := model.NewAccount("js123456", "888888", 22)
	if account != nil {
		fmt.Println("账户创建成功")
	} else {
		fmt.Println("账户创建失败")
	}
	fmt.Println("account = ", account)
}
