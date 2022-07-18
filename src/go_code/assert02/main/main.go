// 编写函数，判断参数是什么类型
package main

import (
	"fmt"
)

// 定义 Student 类型
type Student struct{}

// 可以接收多个不同类型的参数
func JudgeType(items ...interface{}) {
	for i, v := range items {
		i = i + 1
		// 类型断言
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值：%v \n", i, v)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值：%v \n", i, v)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值：%v \n", i, v)
		case int, int16, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值：%v \n", i, v)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值：%v \n", i, v)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值：%v \n", i, v)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值：%v \n", i, v)
		default:
			fmt.Printf("第%v个参数类型不确定，值：%v \n", i, v)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.1
	var name string = "tom"
	address := "北京"
	n3 := 64
	stu1 := Student{}
	stu2 := &Student{}
	var a interface{}
	var b *interface{}
	JudgeType(n1, n2, name, address, n3, stu1, stu2, a, b)

	fmt.Printf("a 类型： %T, 值： %v \n", a, a) // a 类型： <nil>, 值： <nil>
	fmt.Printf("b 类型： %T, 值： %v \n", b, b) // b 类型： *interface {}, 值： <nil>
}
