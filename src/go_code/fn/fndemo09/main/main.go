package main

import "fmt"

// 全局变量
var age int = 18
var Name string = "xiao~"

// var Sex string
// Sex = "男"

// 局部变量
func test() {
	// age和Name的作用域只在test()函数内部有效
	age := 20
	Name := "tom"
	fmt.Println("test()age=", age)
	fmt.Println("test()Name=", Name)
}

// 练习
var name string = "jerry"

func test01() {
	fmt.Println("test01()name=", name)
}
func test02() {
	// name := "jack" // 声明一个新的变量
	name = "jack" // 变量赋值
	fmt.Println("test02()name=", name)
}

func main() {
	test()
	// fmt.Println("age=", age)
	// fmt.Println("Name=", Name)

	// 练习测试
	fmt.Println("main()name=", name)
	test01()
	test02()
	test01()

	// for i := 0; i < 1; i++ {
	// 	fmt.Println("i=", i)
	// }
	// for i := 0; i < 1; i++ {
	// 	fmt.Println("i=", i)
	// }
}
