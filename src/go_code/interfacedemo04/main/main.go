package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Student struct {
	Name string
}

// 如果 Student 没有实现 Say() 方法，则调用时报错
// Student does not implement AInterface (missing Say method)
func (stu Student) Say() {
	fmt.Printf("Say name= %v \n", stu.Name)
}

// Student 结构体同时实现了 AInterface 和 BInterface 接口
func (stu Student) Hello() {
	fmt.Printf("Hello name= %v \n", stu.Name)
}

type cusInt int

// 数值类型实现接口
func (i cusInt) Say() {
	fmt.Println("我是数值类型")
}

func main() {
	var a AInterface
	fmt.Printf("a类型： %T, 值： %v \n", a, a)
	// 直接调用接口的方法 报错
	// 需要通过自定义类型实现方法
	// a.Say() // panic: runtime error: invalid memory address or nil pointer dereference

	// Student 类型实现了 AInterface 接口定义的所有方法(即实现了接口)
	// 所以， Student 的实例 stu 可以赋值给 AInterface 类型
	var stu Student = Student{"tom"}
	var a1 AInterface = stu
	a1.Say() // name= tom

	var b1 BInterface = stu
	b1.Hello()

	// 数值类型调用接口
	var i cusInt
	callSay(i)
}

func callSay(a AInterface) {
	a.Say()
}
