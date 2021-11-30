package main

import (
	"fmt"
)

// 累加器
// 闭包 - 函数柯里化
// 返回值类型： func(int) int
func AddUpper() func(int) int {
	var n int = 100
	var str = "hello"
	return func(i int) int {
		n = n + i
		fmt.Println("i=", i)
		str += string(36) // ascii 36 = '$'
		fmt.Printf("str==%s\n", str)
		return n
	}
}

func main() {
	f := AddUpper()
	// fmt.Println(AddUpper()(1)) // 101
	fmt.Println("f(1)=", f(1)) // 101
	fmt.Println("f(2)=", f(2)) // 103
	fmt.Println("f(3)=", f(3)) // 106
}
