package main

import (
	"fmt"
	util "go_code/fndemo1/utils"
)

// func calc(n1 float64, n2 float64, opt byte) float64 {
// 	var res float64
// 	switch opt {
// 	case '+':
// 		res = n1 + n2
// 	case '-':
// 		res = n1 - n2
// 	case '*':
// 		res = n1 * n2
// 	case '/':
// 		res = n1 / n2
// 	default:
// 		fmt.Println("不合法的运算符")
// 	}
// 	return res
// }

func main() {
	fmt.Println("hello fn")
	var n1 float64 = 1.233
	var n2 float64 = 2.111
	var opt byte = '+'
	var result float64 = util.Calc(n1, n2, opt)
	fmt.Printf("%T %.2f \n", result, result)

	str := util.TestFn()
	fmt.Println("str=", str)
}
