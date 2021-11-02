package main

import "fmt"

func main() {
	fmt.Println("hello")
	var name string // 姓名
	var age int     // 年龄
	var sal float32 // 薪水
	var isPass bool // 是否通过考试

	// 1、Scanln 程序执行到这里，会卡住，等待用户输入，并回车
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name) // 传地址， 相当于是引用传递

	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试")
	// fmt.Scanln(&isPass)

	// fmt.Printf("name=%v\nage=%v\nsal=%v\nisPass=%v\n", name, age, sal, isPass)
	fmt.Println("请输入你的姓名、年龄、薪水、是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("name=%v\nage=%v\nsal=%v\nisPass=%v\n", name, age, sal, isPass)
}
