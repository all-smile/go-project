package main

import "fmt"

// 引用类型
func test(slice []int) {
	slice[0] = 100
}

func main() {
	var slice = []int{1, 2, 3}
	fmt.Println("slice = ", slice) // slice =  [1 2 3]
	test(slice)
	fmt.Println("slice = ", slice) // slice =  [100 2 3]
}
