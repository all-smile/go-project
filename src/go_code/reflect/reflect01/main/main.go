// 反射 reflect
package main

import (
	"fmt"
	"reflect"
)

// 基本数据类型反射
func reflecttest01(b interface{}) {
	// 通过反射获取到变量的 Type 信息
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rTyp类型： %T, 值： %v \n", rTyp, rTyp) // rTyp类型： *reflect.rtype, 值： int
	fmt.Println("rTyp.Name", rTyp.Name())

	// 获取 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal类型： %T, 值： %v \n", rVal, rVal) // rVal类型： reflect.Value, 值： 100
	fmt.Println("rVal Kind : ", rVal.Kind())       // int
	// n2 := 10 + rVal.Int()
	// fmt.Println("原始值： ", rVal.Int()) // 原始值：  100
	// fmt.Println("n2 = ", n2)         // 110

	// rVal.SetInt(20) // panic: reflect: reflect.Value.SetInt using unaddressable value
	rVal.Elem().SetInt(20)

	// 将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV 类型： %T， 值： %v \n", iV, iV) // iV 类型： int， 值： 100
	// 断言
	num2, ok := iV.(int)
	if ok {
		fmt.Printf("num2 类型： %T， 值： %v \n", num2, num2) // num2 类型： int， 值： 100
	} else {
		fmt.Println("断言失败XXXX")
	}
}

// 结构体类型反射
func reflecttest02(b interface{}) {
	// 通过反射获取到变量的 Type 信息
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rTyp类型： %T, 值： %v \n", rTyp, rTyp) // rTyp类型： *reflect.rtype, 值： main.Stu
	fmt.Println("rTyp.Name", rTyp.Name())          // rTyp.Name Stu

	// 获取 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal类型： %T, 值： %v \n", rVal, rVal) // rVal类型： reflect.Value, 值： {tom 20}
	fmt.Println("rVal Kind : ", rVal.Kind())       // rVal Kind :  struct

	// 将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV 类型： %T， 值： %v \n", iV, iV) // iV 类型： main.Stu， 值： {tom 20}
	// 虽然 iv的 值是 {tom 20}， 但是仍然无法直接取出 Name,
	// 反射是运行时，不是编译时，仍需要断言
	// fmt.Printf("name = %v \n", iV.Name) // iV.Name undefined (type interface {} is interface with no methods)
	// 带检测的断言
	_stu, ok := iV.(Stu)
	if ok {
		fmt.Printf("_stu 类型： %T， 值： %v \n", _stu, _stu) // _stu 类型： main.Stu， 值： {tom 20}
		fmt.Printf("name = %v \n", _stu.Name)
	} else {
		fmt.Println("断言失败")
	}
}

type Stu struct {
	Name string
	Age  int
}

func main() {
	// 基本数据类型 - interface{} - reflect.ValueOf()
	var num int = 100
	reflecttest01(&num)
	fmt.Println("num ===== ", num)

	// 结构体数据类型 - interface{} - reflect.ValueOf()
	var stu Stu = Stu{
		Name: "tom",
		Age:  20,
	}
	reflecttest02(stu)
}
