package main

import (
	"fmt"
	// 引包
	"go_code/pubsymbol/counters"
	"go_code/pubsymbol/model"
)

func main() {
	// 创建一个未公开类型的变量，并初始化为10
	counter := counters.New(10)
	fmt.Printf("counter类型：%T, 值：%v \n ", counter, counter) // counter类型：counters.AlertCounter, 值：10

	// 创建 Student 结构体变量
	/* stu := model.Student{
		Name: "tom",
		Age:  2,
	} */
	stu := model.NewStudent("tom02", 3)

	fmt.Println("stu = ", *stu)                 // {tom02 3}
	fmt.Println("stu.Name = ", stu.Name)        // tom02
	fmt.Println("stu.score = ", stu.GetScore()) // 没有赋值， 默认零值 0
}
