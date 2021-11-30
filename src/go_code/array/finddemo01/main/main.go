package main

import "fmt"

func main() {
	// 顺序查找
	arr := [4]string{"路飞", "索隆", "山治", "甚平"}
	var name string
	fmt.Println("请输入名字查找")
	fmt.Scanln(&name)

	// 第一种方式，直接输出结果
	for i := 0; i < len(arr); i++ {
		if arr[i] == name {
			fmt.Printf("找到%s了，是第%v个\n", name, i)
			break
		} else if i == len(arr)-1 {
			fmt.Printf("没有找到%s", name)
		}
	}
	// 上面的写法看起来 思路不是很清晰，代码不是很优雅，
	// 第二种方式
	// 定义下标, 找到之后，重新赋值下标，没有则是-1
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == name {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%s了，是第%v个\n", name, index)
	} else {
		fmt.Printf("没有找到%s", name)
	}
}
