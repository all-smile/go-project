package utils

import "fmt"

// 1、打印金字塔函数封装
// level 层级
func Pyramid(level int) {
	for i := 1; i <= level; i++ {
		for n := 0; n < level-i; n++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (i*2 - 1); j++ {
			// 空心金字塔
			if (j > 1 && j < (i*2-1)) && i != level {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

// 打印九九乘法表
func Mtable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v  ", j, i, i*j)
		}
		fmt.Println()
	}
}
