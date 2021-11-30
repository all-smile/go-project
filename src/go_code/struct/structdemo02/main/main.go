package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type People Person

type integer int

func main() {
	// 结构体变量定义方式

	// 方式1
	var p1 Person
	fmt.Println("p1 = ", p1)

	// 方式2
	// var p2 = Person{}
	// var p2 = Person{"tom", 10}
	var p2 = Person{
		Name: "tom",
		Age:  20,
	}
	fmt.Println("p2 = ", p2)

	// 方式3
	var p3 *Person = new(Person)
	// 指针结构体变量，标准赋值方式
	(*p3).Name = "abc" // 这个写法不方便
	// (*p3).Name = "abc" 等价于  p3.Name = "abc"
	// go 底层会做指针取值处理
	p3.Age = 3
	fmt.Println("p3 = ", *p3)

	// 方式4
	var p4 *Person = &Person{}
	(*p4).Name = "hhh"
	p4.Age = 5
	fmt.Println("p4 = ", p4)

	// 结构体类型转化
	var p5 Person
	var people1 People
	// people1 = p5 // 错误
	people1 = People(p5)               // 强转类型
	fmt.Println("people1 = ", people1) // people1 =  { 0}

	var n1 integer = 20
	var n2 int = 20
	// n1 = n2 // 错误 integer 和 int golang认为不是同一个类型
	fmt.Printf("n1 类型： %T, 值： %v \n", n1, n1) // n1 类型： main.integer, 值： 20
	fmt.Printf("n2 类型： %T, 值： %v \n", n2, n2) // n2 类型： int, 值： 20
}
