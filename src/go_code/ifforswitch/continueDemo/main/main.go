package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 5; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}
here:
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 5; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}
