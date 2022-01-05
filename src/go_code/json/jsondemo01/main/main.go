// json 序列化
package main

import (
	"encoding/json"
	"fmt"
)

// 这就需要通过 tag 来转化大写的字段，同时又保证内部调用的时候可以暴露变量
// 反射原理
type People struct {
	Name   string  `json:"name01"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func testStruct() {
	p1 := People{
		Name:   "tom",
		Age:    18,
		Height: 178.5,
	}
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("testStruct = ", string(data)) // {"name01":"tom","age":18,"height":178.5}
		fmt.Println("Struct p1 = ", p1)            // {tom 18 178.5}
	}
}

// 结构体切片 序列化存在问题， 不建议
// 数组对象应该采用 map切片 []map[string]interface{}
func testSliceStruct() {
	type myObj struct {
		name string
		age  int
	}
	var jsonList []myObj = make([]myObj, 0)
	obj01 := myObj{
		name: "xiao",
		age:  12,
	}
	jsonList = append(jsonList, obj01)
	jsonList = append(jsonList, obj01)
	fmt.Println("jsonList = ", jsonList)
	data, err := json.Marshal(jsonList)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("testSliceStruct = ", string(data)) // {"剑豪":"索隆","海贼王":"luffy"}
	}
}

func testMap() {
	var dict map[string]interface{} = make(map[string]interface{})
	dict["海贼王"] = "luffy"
	dict["剑豪"] = "索隆"
	fmt.Printf("Map dict类型： %T, 值：%v \n", dict, dict)
	data, err := json.Marshal(dict)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("testMap = ", string(data)) // {"剑豪":"索隆","海贼王":"luffy"}
	}
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{} = make(map[string]interface{})
	m1["name"] = "xixi"
	m1["age"] = 7
	var m2 map[string]interface{} = make(map[string]interface{})
	m2["name"] = "haha"
	m2["age"] = 8
	m2["addr"] = [2]string{"北京", "上海"}
	slice = append(slice, m1, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("testSlice = ", string(data)) // [{"age":7,"name":"xixi"},{"addr":["北京","上海"],"age":8,"name":"haha"}]
	}
}

func testFloat64() {
	var n1 float64 = 12.21
	data, err := json.Marshal(n1)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("testFloat64 = ", string(data)) // 12.21
	}
}

func main() {
	// struct
	testStruct()
	// map
	testMap()
	// slice
	testSlice()
	// 基本数据类型序列化
	// 对基本数据类型序列化 意义不大
	testFloat64()
	testSliceStruct()
}
