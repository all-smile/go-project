package main

func add(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	return res
}

func main() {
	// 传统测试方式
	// var sum int = add(5)
	// if sum != 10 {
	// 	fmt.Printf("add函数错误, 返回值：%v, 期望值： %v \n", sum, 10)
	// } else {
	// 	fmt.Printf("add函数正确, 返回值：%v, 期望值： %v \n", sum, 10)
	// }
}
