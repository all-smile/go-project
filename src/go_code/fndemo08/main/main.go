package main

import "fmt"

func sum(n1 int, n2 int) int {
	// 当执行到 defer 时，暂时跳过不执行，将defer后面的语句压入一个新的栈中（defer栈）
	// 当函数其它代码执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	// 可以用来函数执行完毕之后，释放加载的资源
	defer fmt.Println("defer n1=", n1) // 3 defer n1= 10  // 值拷贝
	defer fmt.Println("defer n2=", n2) // 2 defer n2= 20

	// 测试值拷贝
	n1++
	n2++
	fmt.Println("n1=", n1) // n1= 11
	fmt.Println("n2=", n2) // n2= 21

	res := n1 + n2
	fmt.Println("res=", res) // 1 res= 32
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("main-res=", res) // 4 main-res= 32
}
