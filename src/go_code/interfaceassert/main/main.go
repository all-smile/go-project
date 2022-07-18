// 此案例演示类型断言 assert
package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	// 定义空接口
	var a interface{}
	fmt.Println("a = ", a)
	var point Point = Point{1, 2}
	a = point // OK
	fmt.Printf("a 类型： %T, 值： %v \n", a, a)

	// 怎么把 变量 赋值给 空接口
	var b Point
	fmt.Println("b = ", b)
	// cannot use a (type interface {}) as type Point in assignment: need type assertion
	// b = a // ERR
	// b = a.(Point) // 类型断言
	/*
		b = a.(Point) // 类型断言
		表示判断 a 是否是指向 Point 类型的变量，（上面	a = point 代码，a已经指向了Point）
		如果是，就把 a 转换成 Point 类型，并赋给 b 变量
		否则，就报错 (把上面 a = point 代码 去掉，就报错，类型推导不出来)
	*/
	fmt.Println("b = ", b)

	// 类型断言 float32 - interface{}
	var m float32
	var x interface{}
	x = m
	fmt.Printf("x 类型： %T, 值： %v \n", x, x)
	// x 类型： float32, 值： 0
	n := x.(float32)
	fmt.Printf("n 类型： %T, 值： %v \n", n, n)
	// n 类型： float32, 值： 0

	// 带检测的类型断言
	n2, ok := x.(float32)
	if ok {
		fmt.Println("convert success")
		fmt.Printf("n2 类型： %T, 值： %v \n", n2, n2)
	} else {
		fmt.Println("convert error")
	}
	// panic: interface conversion: interface {} is float32, not float64

	/* if n2, ok := x.(float64); ok {
	} */

	fmt.Println("panic 之后还会执行吗")
}
