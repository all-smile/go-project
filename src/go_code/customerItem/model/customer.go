package model

import "fmt"

// 定义 Customer 结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式的函数，创建 Customer 实例， 并返回 Customer 实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 第二种创建 Customer 实例的方法，不穿 Id， 由系统分配
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 编写 GetInfo() 方法，返回当前用户的信息
func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		customer.Id,
		customer.Name,
		customer.Gender,
		customer.Age,
		customer.Phone,
		customer.Email)
	return info
}
