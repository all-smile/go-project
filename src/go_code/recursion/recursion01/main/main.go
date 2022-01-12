// 递归

package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("test n = ", n)
}

func test02(n int) {
	if n > 2 {
		n--
		test02(n)
	} else {
		fmt.Println("test02 n = ", n)
	}
}

func main() {
	test(4)
	test02(4)
}

/*
test n =  2
test n =  2
test n =  3
test02 n =  2
*/
