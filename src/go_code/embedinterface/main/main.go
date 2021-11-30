// 演示如何将嵌入类型应用于接口
package main

import "fmt"

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的notifier接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s> \n",
		u.name,
		u.email)
}

// admin 代表一个拥有权限得管理员用户
// 要嵌入一个类型，只需要声明这个类型的名字就可以了
// user 是外部类型 admin 的内部类型
type admin struct {
	user  // 嵌入类型
	level string
}

func main() {
	// 创建一个 admin 用户
	// 内部类型的初始化采用结构字面量的方式完成
	ad := admin{
		user: user{
			name:  "join",
			email: "join@yahoo.com",
		},
		level: "super",
	}

	// 给 admin 用户发送一个通知
	// 用于实现接口的内部类型的方法，被提升到外部类型
	sendNotification(&ad)
}

// sendNotification 接收一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
