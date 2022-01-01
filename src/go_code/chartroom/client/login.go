package main

import "fmt"

func login(userId int, userPwd string) (err error) {
	// 定协议
	fmt.Printf("userId = %d, userPwd = %s \n", userId, userPwd)
	return nil
}
