package main

import (
	"fmt"
)

func test(char byte) byte {
	return char + 1
}

func main() {
	var key byte
	fmt.Println("请输入单个字母字符")
	fmt.Scanf("%c", &key) // %c 该值对应的unicode码值
	fmt.Printf("%T %v \n", key, key)
	switch test(key) {
	case 'a':
		fmt.Println("今天星期一")
	case 'b':
		fmt.Println("今天星期二")
	case 'c':
		fmt.Println("今天星期三")
	default:
		fmt.Println("今天可能是周末!")
	}

	var num int = 10
	switch num {
	case 10:
		fmt.Println("101010")
		fallthrough
	case 20:
		fmt.Println("202020")
	}
}
