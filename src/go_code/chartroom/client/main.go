package main

import "fmt"

var (
	userId  int
	userPwd string
)

func main() {
	// 接收用户选择
	var key int
	var loop bool
	for {
		fmt.Println("************************* 欢迎登录多人聊天系统 *************************")
		fmt.Println("\t\t\t 1. 登录聊天室")
		fmt.Println("\t\t\t 2. 注册用户")
		fmt.Println("\t\t\t 3. 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanln(&key)
		fmt.Printf("key 类型： %T, 值： %v \n", key, key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，重新输入")
		}
		if !loop {
			return
		}
		if key == 1 {
			fmt.Println("请输入用户 ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%d\n", &userPwd)
			err := login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败")
			} else {
				fmt.Println("登陆成功")
			}
		} else if key == 2 {
			fmt.Println("用户注册")
		}
	}
}
