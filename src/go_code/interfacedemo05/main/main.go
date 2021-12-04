// 接口继承，以及实现
package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}

// 要实现 AInterface 接口，必要同时实现 BInterface、CInterface接口
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Student struct {
}

func (stu Student) test01() {}
func (stu Student) test02() {}
func (stu Student) test03() {}

type blankInterface interface{}

func main() {
	var stu Student
	var a AInterface = stu
	a.test01()

	// 空接口
	var b blankInterface
	var blank blankInterface = 10
	fmt.Printf("b类型： %T, 值： %v \n", b, b)             // b类型： <nil>, 值： <nil>
	fmt.Printf("blank类型： %T, 值： %v \n", blank, blank) // blank类型： int, 值： 10
}
