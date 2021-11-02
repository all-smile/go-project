package main

import "fmt"

// 全局变量声明
var num1 = 100
var num2 = 200
var name3 = "tom"

// 上面的写法可以简化成下面的， 注意没有,号
// 全局多变量声明
var (
	x = 22
	y = 33
	z = "44"
)

// non-declaration statement outside function body
// 不使用var声明，只适用与函数内部，全局变量不行
// test := "123132"
// fmt.Println('test=', test)

func main() {
	// 局部变量声明

	// 局部单一变量
	// 1、指定变量类型，声明后不赋值，go会自动给定默认值
	var i int
	i = 10
	fmt.Println("i=", i)
	// 2、根据值，自行判定变量类型（类型推导）
	var n = 10.2
	fmt.Println("n=", n)
	// 3、省略var，:= 左边的变量不应该是已经声明过的，否则会报错
	name := "xiao"
	fmt.Println("name=", name)

	// 局部多变量声明
	// 1、同一种类型
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	// 2、不同类型
	var num, name2, sex = 12, "xiao", true
	fmt.Println("num=", num, "name2=", name2, "sex=", sex)
	// 3、类型推导
	a, b, c := 1, "str", false
	fmt.Println("a=", a, "b=", b, "c=", c)

	// 全局变量打印
	fmt.Println("num1=", num1, "num2=", num2, "name3=", name3)
	fmt.Println("x=", x, "y=", y, "z=", z)
}

// syntax error: non-declaration statement outside function body
// fmt.Println("x=", x, "y=", y, "z=", z)
