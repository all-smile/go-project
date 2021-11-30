package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("hello1")
	n := 20
	if n > 20 {
		goto here
	}
	fmt.Println("hello2")
	fmt.Println("hello3")
	fmt.Println("hello4")
	fmt.Println("hello5")
here:
	fmt.Println("hello6")
	fmt.Println("hello7")

	// return
	i := 0
	for ; i < 10; i++ {
		if i == 5 {
			return
		}
		fmt.Println("0_0", i)
	}
	fmt.Println("i=", i)
}
