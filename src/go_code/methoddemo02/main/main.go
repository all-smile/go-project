package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 给 Student 结构体 实现 String() 方法
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	// 定义 Student 变量
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println("stu = ", stu) // {tom 20}

	// 如果实现了 *Student类型的 String() 方法，fmt.Println就会自动调用
	fmt.Println("stu = ", &stu) // Name=[tom] Age=[20]
}
