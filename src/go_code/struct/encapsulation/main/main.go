package main

import (
	"fmt"
	"go_code/struct/encapsulation/model"
)

func main() {
	p := model.NewPerson("tom")
	// p.age = 18 // p.age undefined (cannot refer to unexported field or method age)
	p.SetAge(18)
	p.SetSal(20000)
	fmt.Println("*p = ", *p)
	fmt.Println("p.Name = ", p.Name)
	fmt.Println("p.age = ", p.GetAge())
	fmt.Println("p.sal = ", p.GetSal())
}
