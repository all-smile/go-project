package main

import "fmt"

/*
1、声明（定义）结构体，确定结构体名。
2、编写结构体字段
3、编写结构体方法
*/

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu *Student) Say() string {
	info := fmt.Sprintf("学生的信息：name=[%v], gender=[%v], age=[%v], id=[%v], score=[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return info
}

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

type Vistor struct {
	Name string
	Age  int
}

func (vistor *Vistor) showPrice() {
	if vistor.Age > 90 || vistor.Age < 8 {
		fmt.Println("考虑到安全，就不要玩了")
		return
	}
	if vistor.Age > 18 {
		fmt.Printf("游客姓名：%v, 年龄：%v, 收费：%v", vistor.Name, vistor.Age, 20)
	} else {
		fmt.Printf("游客姓名：%v, 年龄：%v, 免费", vistor.Name, vistor.Age)
	}
}

func main() {
	stu := Student{
		name:   "tom",
		gender: "male",
		age:    12,
		id:     001,
		score:  99.99,
	}
	info := stu.Say()
	fmt.Println("info = ", info)
	// 学生的信息：name=[tom], gender=[male], age=[12], id=[1], score=[99.99]

	box := Box{
		len:    6,
		width:  6,
		height: 6,
	}
	vol := box.getVolumn()
	fmt.Println("box的体积 = ", vol)

	vistor := Vistor{
		Name: "tom",
		Age:  22,
	}
	vistor.showPrice()
}
