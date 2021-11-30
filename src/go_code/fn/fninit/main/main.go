package main

import (
	"fmt"
	"go_code/fninit/utils"
)

var a int = testVars()

func testVars() int {
	fmt.Println("testVars")
	return 90
}

func init() {
	fmt.Println("init")
}

// 全局匿名函数
var (
	fn01 = func(n1 int, n2 int) int {
		return n1 - n2
	}
	fn02 = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {
	fmt.Println("main", a)
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)

	// 匿名函数
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println("sub=", sub(20, 10))

	fmt.Println("fn01=", fn01(20, 10))
	fmt.Println("fn02=", fn02(20, 10))

}
