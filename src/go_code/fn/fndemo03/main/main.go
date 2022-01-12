// 函数递归调用 + 应用场景
package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n = ", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n = ", n)
	}
}

// 斐波那契数
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-2) + fbn(n-1)
	}
}

// 猴子吃桃子
// 每一天吃当天桃子数量的一半多1个，第十天的时候发现只剩下一个
// 求第一天的桃子数量
func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("您输入的天数不对")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	test(4)
	// test2(4)

	// 斐波那契数列
	// 1 ，1 ，2 ，3 ，5 ，8， 13 ，21 ，......
	// f(n) = f(n-2) + f(n-1)
	fmt.Println("value=", fbn(5)) // 第5个斐波那契数
	// fmt.Println("value=", fbn(6)) // 第6个斐波那契数
	// fmt.Println("value=", fbn(7))
	// fmt.Println("value=", fbn(8))

	fmt.Println("第1天的桃子数量", peach(1))
}
