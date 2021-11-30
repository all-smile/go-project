package main

import "fmt"

func test(n1 int) {
	n1++
	fmt.Println("test() n1=", n1)
}

func sum(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)
	total, reduce := sum(10, 20)
	fmt.Printf("total%T %v \n", total, total)
	fmt.Printf("reduce%T %v \n", reduce, reduce)
}
