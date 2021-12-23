// 反射 reflect
package main

import (
	"fmt"
	"reflect"
)

func reflecttest01(b interface{}) {
	// 通过反射获取到变量的 Type 信息
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rTyp类型： %T, 值： %v \n", rTyp, rTyp) // rTyp类型： *reflect.rtype, 值： int
	fmt.Println("rTyp.Name", rTyp.Name())

	// 获取 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal类型： %T, 值： %v \n", rVal, rVal) // rVal类型： reflect.Value, 值： 100
	fmt.Println("rVal Kind : ", rVal.Kind())       // int
	n2 := 10 + rVal.Int()
	fmt.Println("原始值： ", rVal.Int()) // 原始值：  100
	fmt.Println("n2 = ", n2)         // 110

	// 将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV 类型： %T， 值： %v \n", iV, iV) // iV 类型： int， 值： 100
	// 断言
	num2 := iV.(int)
	fmt.Printf("num2 类型： %T， 值： %v \n", num2, num2) // num2 类型： int， 值： 100
}

func main() {
	// 基本数据类型 - interface{} - reflect.ValueOf()
	var num int = 100
	reflecttest01(num)
}
