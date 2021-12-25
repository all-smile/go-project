// 结构体 反射 案例
// 反射可以用来开发底层框架
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 这就需要通过 tag 来转化大写的字段，同时又保证内部调用的时候可以暴露变量
// 反射原理
type People struct {
	Name   string  `json:"name01"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Action string
}

// 给结构体绑定方法

func (p People) Print() {
	fmt.Println("print() = ", p)
}

func (p People) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (p People) Set(name string, age int, height float64, action string) {
	p.Name = name
	p.Age = age
	p.Height = height
	p.Action = action
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("except struct, but its not")
		return
	}
	// 结构体字段
	// 字段的顺序跟定义结构体时字段的先后顺序一致
	num := val.NumField()
	fmt.Printf("struct has %v fields \n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("第%v个字段：%v \n", i+1, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("第%v个字段, tag ： %v \n", i+1, tagVal)
		}
	}
	// 结构体方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct 有 %v 个方法 \n", numOfMethod)

	// 内部方法名顺序，按照函数名的 ascii 码排序的
	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res = ", res[0].Int())
}

func main() {
	// struct

	p1 := People{
		Name:   "tom",
		Age:    18,
		Height: 178.5,
		Action: "衣食住行",
	}

	// 序列化就是通过反射机制获取到 tag 值
	p1Str, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("序列化失败", err)
	} else {
		fmt.Println("p1Str = ", string(p1Str))
	}
	TestStruct(p1)
	fmt.Println("p1 = ", p1)
}
