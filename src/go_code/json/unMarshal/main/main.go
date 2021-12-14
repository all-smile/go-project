// json 反序列化
package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name   string  `json:"name01"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func unMarshalStruct() {
	var jsonStr string = "{\"name01\":\"tom\",\"age\":18,\"height\":178.5}"
	var p1 People
	// 要传入 &p1 指针类型，要写入的变量实例（可修改）
	// 用 &p1 接收反序列化后的结果
	err := json.Unmarshal([]byte(jsonStr), &p1)
	if err != nil {
		fmt.Println("反序列化失败", err)
	} else {
		fmt.Println("p1 = ", p1) // {tom 18 178.5}
	}
}

// 测试序列化和反序列化
func testMap() string {
	var dict map[string]interface{} = make(map[string]interface{})
	dict["海贼王"] = "luffy"
	dict["剑豪"] = "索隆"
	fmt.Printf("dict类型： %T, 值：%v \n", dict, dict)
	data, err := json.Marshal(dict)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("data = ", string(data)) // {"剑豪":"索隆","海贼王":"luffy"}
	}
	return string(data)
}

func unMarshalMap() {
	// var jsonStr string = "{\"剑豪\":\"索隆\",\"海贼王\":\"luffy\"}"
	var jsonStr string = testMap()
	var dict map[string]interface{}
	// 注意： 反序列化map不需要 make 操作
	// make 操作被封装到 json.Unmarshal() 里面
	err := json.Unmarshal([]byte(jsonStr), &dict)
	if err != nil {
		fmt.Println("反序列化失败", err)
	} else {
		fmt.Println("dict = ", dict) // map[剑豪:索隆 海贼王:luffy]
	}
}

func unMarshalSlice() {
	var jsonStr = "[{\"age\":7,\"name\":\"xixi\"}," +
		"{\"addr\":[\"北京\",\"上海\"],\"age\":8,\"name\":\"haha\"}]"
	var slice []map[string]interface{}
	// 同样不需要 make
	err := json.Unmarshal([]byte(jsonStr), &slice)
	if err != nil {
		fmt.Println("反序列化失败", err)
	} else {
		fmt.Println("slice = ", slice) // [map[age:7 name:xixi] map[addr:[北京 上海] age:8 name:haha]]
	}
}

func add(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	return res
}

func main() {
	unMarshalStruct()
	unMarshalMap()
	unMarshalSlice()

	var sum int = add(5)
	if sum != 10 {
		fmt.Printf("add函数错误, 返回值：%v, 期望值： %v \n", sum, 10)
	} else {
		fmt.Printf("add函数正确, 返回值：%v, 期望值： %v \n", sum, 10)
	}
}
