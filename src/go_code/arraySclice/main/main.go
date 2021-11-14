package main

import "fmt"

func main() {
	// 求3个数的平均数
	n1 := 10.289
	n2 := 9.98
	n3 := 10.645
	var total = n1 + n2 + n3
	var avgNum = total / 3
	_avgNum := fmt.Sprintf("%.2f", avgNum) // 保留两位小数

	fmt.Println("avgNum = ", avgNum)   // avgNum =  10.304666666666666
	fmt.Println("_avgNum = ", _avgNum) // avgNum =  10.304666666666666
}
