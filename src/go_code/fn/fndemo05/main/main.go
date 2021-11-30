package main

import "fmt"

func swap(n1 *int, n2 *int) {
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2 // *n1
	*n1 = *n1 - *n2 // *n2
}

func main() {
	a := 12
	b := 20
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
}
