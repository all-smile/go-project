package main

import "fmt"

func main() {
	fmt.Printf("2&3=%v\n", 2&3)   // 2
	fmt.Printf("2|3=%v\n", 2|3)   // 3
	fmt.Printf("2^3=%v\n", 2^3)   // 1
	fmt.Printf("-2^2=%v\n", -2^2) // -4

	a := 1 >> 2
	b := 1 << 2
	fmt.Printf("1 >> 2 = %v \n", a) // 0
	fmt.Printf("1 << 2 = %v \n", b) // 4
}
