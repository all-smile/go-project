// 这个示例程序展示如何将一个类型嵌入另一个类型
// 内部类型和外部类型之间的关系
package main

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 实现了一个可以通过user类型值得指针调用得方法
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
	// 直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法也被提升到外部类型
	ad.notify()

	/*
		输出
		Sending user email to join<join@yahoo.com>
		Sending user email to join<join@yahoo.com>
	*/
}
