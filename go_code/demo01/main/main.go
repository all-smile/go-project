package main

import "fmt"

func main() {
	// 二进制
	// go不能直接输出二进制，可以使用 %b 格式化
	// fmt %b	表示为二进制
	var n int = 5
	fmt.Printf("%T %b\n", n, n) // 101

	// 8进制, 0-7
	// 以数字0开头表示
	var n1 int = 011
	fmt.Printf("%T %v\n", n1, n1) // 9

	// 16进制，0-9A-F,
	// 以0x或者0X开头表示
	var n2 int = 0x11
	fmt.Printf("%T %v\n", n2, n2) // 17
}
