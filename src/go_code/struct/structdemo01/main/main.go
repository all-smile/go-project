package main

import "fmt"

// 字段类型默认值：指针、slice、map的零值都是nil，即还没有分配空间
type Person struct {
	Name  string
	Age   int
	Score float64
	Ptr   *int
	Slice []int
	Map   map[string]string
}

type Monster struct {
	Name string
	Age  int
}

func main() {

	// 定义结构体
	type Cat struct {
		Name  string
		Age   int
		Color string
	}

	// 使用
	var cat1 Cat // 像使用其它变量一样

	fmt.Printf("cat1 地址 = %p \n", &cat1)

	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("cat1 = ", cat1) // {小白 3 白色}
	fmt.Printf("cat1 地址 = %p \n", &cat1)

	//
	var p1 Person
	fmt.Println("p1 = ", p1)
	fmt.Println("p1.Ptr == nil", p1.Ptr == nil)     // true
	fmt.Println("p1.Slice == nil", p1.Slice == nil) // true
	fmt.Println("p1.Map == nil", p1.Map == nil)     // true
	// p1.Slice[0] = 100 // panic: runtime error: index out of range [0] with length 0
	// 使用Slice、Map 先 make
	p1.Slice = make([]int, 5) // slice需要指定长度(容量)
	fmt.Println("make后 p1 = ", p1)
	p1.Slice[0] = 100
	p1.Map = make(map[string]string) // map会自动增长，不需要指定长度和容量
	p1.Map["hobby"] = "coding"
	fmt.Println("p1 = ", p1) // { 0 0 <nil> [100 0 0 0 0] map[hobby:coding]}

	// 不同结构体变量的字段互不影响
	var monster1 Monster
	monster1.Name = "小牛"
	monster1.Age = 2
	monster2 := monster1
	monster2.Name = "小马"
	fmt.Println("monster1 = ", monster1)
	fmt.Println("monster2 = ", monster2)
}
