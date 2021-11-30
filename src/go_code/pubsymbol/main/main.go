package main

import (
	"fmt"
	// 引包
	"go_code/pubsymbol/counters"
)

func main() {
	// 创建一个未公开类型的变量，并初始化为10
	counter := counters.New(10)
	fmt.Printf("counter类型：%T, 值：%v \n ", counter, counter) // counter类型：counters.AlertCounter, 值：10

}
