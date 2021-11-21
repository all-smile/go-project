package main

import "fmt"

func main() {
	// 定义map切片, 长度容量都是2
	var monsters []map[string]string = make([]map[string]string, 2)
	// 增加一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "小牛"
		monsters[0]["age"] = "10"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "小马"
		monsters[1]["age"] = "20"
	}

	// 继续增加妖怪
	// panic: runtime error: index out of range [2] with length 2\
	// 数组越界
	/* if monsters[2] == nil {
		monsters[2] = make(map[string]string)
		monsters[2]["name"] = "小猴子"
		monsters[2]["age"] = "22"
	} */

	// 这里使用切片的append函数增长
	var newMonster = map[string]string{
		"name": "新增加的：小羊",
		"age":  "20",
	}
	monsters = append(monsters, newMonster)
	fmt.Println("monsters长度 = ", len(monsters)) // 3
	fmt.Println("monsters = ", monsters)
	// [map[age:10 name:小牛] map[age:20 name:小马] map[age:20 name:新增加的：小羊]]

}
