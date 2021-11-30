// 这个示例程序展示当内部类型和外部类型要实现同一个接口的做法
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

// 通过 user 类型的指针调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s> \n",
		u.name,
		u.email)
}

// admin 代表一个拥有权限得管理员用户
type admin struct {
	user  // 嵌入类型
	level string
}

// 通过 admin 类型的指针调用的方法
func (u *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s> \n",
		u.name,
		u.email)
}

func main() {
	// 创建一个 admin 用户
	ad := admin{
		user: user{
			name:  "john",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	// 给 admin 用户发送通知
	// 接口嵌入的内部类型实现并没有提升到外部类型
	sendNotification(&ad) // Sending admin email to john<john@yahoo.com>

	// 我们可以直接访问内部类型的方法
	ad.user.notify() // Sending user email to john<john@yahoo.com>

	// 内部类型的方法没有被提升
	ad.notify() // Sending admin email to john<john@yahoo.com>
}

// sendNotification 接收一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
