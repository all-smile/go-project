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

func main() {
	// 创建一个 user 类型的值，并发送通知
	u := user{"Bill", "bill@email.com"}
	// sendNotification(u)
	/*
		cannot use u (type user) as type notifier in argument to sendNotification:
		user does not implement notifier (notify method has pointer receiver)
	*/
	sendNotification(&u)
}

// sendNotification 接收一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
