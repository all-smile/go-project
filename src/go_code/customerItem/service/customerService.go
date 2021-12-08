package service

import (
	"fmt"
	"go_code/customerItem/model"
)

// 封装对 Customer 的操作方法(增删改查)
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片有多少客户，
	// 该字段还可以作为客户的Id号
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := CustomerService{}
	customerService.customerNum = 1
	// id int, name string, gender string, age int, phone string, email string
	customer := model.NewCustomer(1, "tom", "男", 20, "16000000000", "16000000000@email.com")
	customerService.customers = append(customerService.customers, customer)
	return &customerService
}

// 编写 List() 方法, 返回客户切片
func (customerService *CustomerService) List() []model.Customer {
	return customerService.customers
}

// 编写 Add() 方法，新增客户
// *CustomerService 指针接收者
func (customerService *CustomerService) Add(customer model.Customer) bool {
	// 客户编号处理
	customerService.customerNum++
	customer.Id = customerService.customerNum
	customerService.customers = append(customerService.customers, customer)
	return true
}

// 编写 Delete() 方法，删除客户
func (customerService *CustomerService) Delete(id int) bool {
	index, customer := customerService.FindById(id)
	fmt.Println("customer = ", customer)
	if index != -1 {
		// 删除 id 对应的客户
		customerService.customers = append(customerService.customers[:index], customerService.customers[index+1:]...)
		return true
	} else {
		fmt.Printf("找不到编号: %v 的客户\n", id)
		return false
	}
}

// 编写 Update() 方法，修改客户
func (customerService *CustomerService) Update(id int, customer model.Customer) bool {
	index, _ := customerService.FindById(id)
	if index != -1 {
		// 修改 id 对应的客户
		customerService.customers[index] = customer
		return true
	} else {
		fmt.Printf("找不到编号: %v 的客户\n", id)
		return false
	}
}

// 编写 FindById(id) 方法，查找客户，并返回客户切片下标, 和对应的客户信息 (编号个底层切片下标存在不对应的情况)
func (customerService *CustomerService) FindById(id int) (int, model.Customer) {
	// 查找客户切片下标，有则返回，没有则返回 -1
	index := -1
	var customer model.Customer
	for i := 0; i < len(customerService.customers); i++ {
		if customerService.customers[i].Id == id {
			index = i
			customer = customerService.customers[i]
		}
	}
	return index, customer
}
