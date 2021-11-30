package main

import "fmt"

func main() {
	num01 := 100
	fmt.Printf("num01的类型： %T, num01的值： %v, num01的地址： %v \n", num01, num01, &num01)

	num02 := new(int) // *int
	fmt.Printf("num02的类型： %T, num02的值： %v, num02的地址： %v \n", num02, num02, &num02)
	fmt.Printf("num02指针的真正的值： %v \n", *num02)

	name := new(string)
	fmt.Printf("name指针的真正的值： %q \n", *name)
}
