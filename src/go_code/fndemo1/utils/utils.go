package utils

import "fmt"

func TestFn() string {
	fmt.Println("TestFn 函数被调用")
	return "hahaha"
}

func Calc(n1 float64, n2 float64, opt byte) float64 {
	var res float64
	switch opt {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("不合法的运算符")
	}
	return res
}
