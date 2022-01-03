package main

import (
	"fmt"
	"go_code/chatroom/client/process"
	"os"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	// 接收用户选择
	var key int
	var loop bool = true
	for loop {
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
			fmt.Println("请输入用户 ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)

			// 创建 UserProcess 实例, 调用 Login
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败", err)
			}

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%v\n", &userPwd)
			fmt.Println("请输入用户昵称:")
			fmt.Scanf("%v\n", &userName)

			// 创建 UserProcess 实例, 调用 Register
			up := &process.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("注册失败", err)
			}

		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0) // 退出系统
		default:
			fmt.Println("输入有误，重新输入")
		}
		// if !loop {
		// 	return
		// }
	}
}
