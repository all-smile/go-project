// 该程序展示封装
package model

import "fmt"

type person struct {
	Name string
	// 以下字段不公开
	age int
	sal float64
}

// 工厂模式函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问 age 和 sal 我们编写一堆 SetXxx、GetXxx方法
func (p *person) SetAge(age int) {
	// 数据校验
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Printf("年龄%v不合适", age)
	}
}
func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	// 数据校验
	if sal > 3000 && sal < 80000 {
		p.sal = sal
	} else {
		fmt.Printf("薪水%v不合适", sal)
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
