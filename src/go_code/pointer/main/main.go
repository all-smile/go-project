package main

import (
	"fmt"
	// 引包
	"go_code/pointer/model"
)

func test() int {
	return 10
}

func main() {
	fmt.Println(model.Name)
	// a = 9, b = 2 => a = 2, b = 9
	a := 9
	b := 2
	a = a + b
	b = a - b // b = a
	a = a - b // a = b
	fmt.Printf("a=%T %v \n", a, a)
	fmt.Printf("b=%T %v \n", b, b)

	var n int = test() + 2 // 右边是表达式
	fmt.Printf("n=%T %v \n", n, n)

	a1 := 100
	fmt.Println("a1的地址：", &a1)
	var ptr *int = &a1
	fmt.Println("ptr指向的值是：", *ptr)

}
