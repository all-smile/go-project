package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monsters struct {
	// 字段必须大写，否则当使用别包的方法 json.Marshal 序列化的时候，读不到字段信息
	// 前端使用者，通常使用小写的字段名,
	// 这就需要通过 tag 来转化大写的字段，同时又保证内部调用的时候可以暴露变量
	// 反射原理
	Name  string `json:"name"`
	Age   int    `json:age`
	Skill string `json:skill`
}

func main() {
	// 结构体转化，必须字段和类型全部一致，才可以转化
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	// 后端开发者： 结构体序列化
	monster := Monsters{"小牛", 3, "芭蕉扇"}
	// 将monster变量序列化为json字符串格式
	monsterStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 序列化失败", err)
	}
	fmt.Printf("monsterStr = %s \n", monsterStr)     // {"Name":"小牛","Age":3,"Skill":"芭蕉扇"}
	fmt.Println("monsterStr = ", string(monsterStr)) // {"Name":"小牛","Age":3,"Skill":"芭蕉扇"}

	// 前端使用者，通常使用小写的字段名  Name  string `json:"name"`
	// 输出： {"name":"小牛","Age":3,"Skill":"芭蕉扇"}
}
