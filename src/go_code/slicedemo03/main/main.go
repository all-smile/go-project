package main

import "fmt"

func fbn(n int) []uint64 {
	// 声明切片， 长度和容量都是 n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	/* 需求：
	1、输入 n int
	2、能够将对应的斐波那契数列存放到切片中
	思路分析
	1、声明函数 fbn(n int) ([]uint64)
	2、编程 fbn(n int), for循环存放斐波那契数列 */

	// 调用生成斐波那契数列
	var slice = fbn(5)
	fmt.Println("slice = ", slice)

}
