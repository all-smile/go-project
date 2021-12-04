package main

import "fmt"

// 定义 Usb 接口， 有两个方法(行为)
type Usb interface {
	start()
	stop()
}

// 定义 Phone 结构体， 并实现 start, stop 方法
type Phone struct {
	Name string
}

// 值接收者
// Phone 类型实现了 Usb 接口 (就是实现了 Usb 接口的所有方法)
func (phone Phone) start() {
	fmt.Printf("%v开始--- \n", phone.Name)
}
func (phone Phone) stop() {
	fmt.Printf("%v停止--- \n", phone.Name)
}

// 定义 Camera 结构体， 并实现 start, stop 方法
type Camera struct {
	Name string
}

// 指针接收者
// Camera 类型实现了 Usb 接口 (就是实现了 Usb 接口的所有方法)
func (camera *Camera) start() {
	fmt.Printf("%v开始--- \n", camera.Name)
}
func (camera *Camera) stop() {
	fmt.Printf("%v停止--- \n", camera.Name)
}

// 定义 Computer 结构体, 并实现 Working 方法
type Computer struct {
}

// 编写 Working 方法，接收一个 Usb 接口类型的变量，
// 只要一个类型实现了某个接口，那么所有使用这个接口的地方，都支持这种类型的值
//  (比如： 实参 phone, camera, 都是 Usb 类型来接收， 这就是多态)
func (computer *Computer) Working(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {
	computer := Computer{}
	phone := Phone{"手机"}
	camera := Camera{"相机"}

	computer.Working(phone)   // 值接收者
	computer.Working(&camera) // 指针接收者
	/*
		手机开始---
		手机停止---
		相机开始---
		相机停止---
	*/
}
