package main

import (
	"fmt"
	"math"
)

func main() {
	// if age := 20; age > 18 {
	// 	fmt.Println("你已经大于18岁了")
	// } else {
	// 	fmt.Println("你还是个未成年人")
	// }

	// var score int
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&score)
	// if score == 100 {
	// 	fmt.Println("奖励mac电脑")
	// } else if score > 90 && score < 100 {
	// 	fmt.Println("奖励iphone手机")
	// } else {
	// 	fmt.Println("奖励ipad")
	// }

	// 求一元二次方程ax2+bx+c=0的根，b2-4ac判断根的个数
	// 引入math包， math.Sqrt(x), x的平方根
	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0
	m := b*b - 4*a*c
	var x1, x2 float64
	if m > 0 {
		x1 = (-b + math.Sqrt(m)) / 2 * a
		x2 = (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("有两个根，分别是%v %v\n", x1, x2)
	} else if m == 0 {
		x1 = (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("有一个根，是 %v\n", x1)
	} else {
		fmt.Printf("无解")
	}
}
