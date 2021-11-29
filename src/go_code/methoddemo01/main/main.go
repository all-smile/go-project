package main

import "fmt"

// user 在程序里定义了一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是主程序入口
func main() {
	// user 类型的值可以用来调用 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指向 user 类型值的指针也可以用来调用 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user 类型的值可以用来调用 使用指针接收者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// user 类型的指针可以用来调用 使用指针接收者声明的方法
	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()

	/*
		输出
		Sending User Email To Bill<bill@email.com>
		Sending User Email To Lisa<lisa@email.com>
		Sending User Email To Bill<bill@newdomain.com>
		Sending User Email To Lisa<lisa@comcast.com>
	*/
}
