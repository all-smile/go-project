package main

import (
	"fmt"
	"go_code/customerItem/model"
	"go_code/customerItem/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

// 编写 list() 方法，调用 CustomerService.List() 展示所有客户信息
func (customer *customerView) list() {
	// 获取当前客户信息
	customers := customer.customerService.List()
	// 显示出来
	fmt.Println("***********************客户列表***********************")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("*********************客户列表完成*********************")
}

// 编写 add() 方法，调用 CustomerService.Add() 方法，新增客户
func (customer *customerView) add() {
	fmt.Println("***********************添加客户***********************")
	fmt.Println("姓名: ")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话: ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱: ")
	email := ""
	fmt.Scanln(&email)
	newCustomer := model.NewCustomer2(name, gender, age, phone, email)
	isAdd := customer.customerService.Add(newCustomer)
	if isAdd {
		fmt.Println("*********************客户添加完成*********************")
	} else {
		fmt.Println("*********************客户添加失败*********************")
	}
}

// 编写 delete() 方法，调用 CustomerService.Delete() 方法，删除客户
func (customer *customerView) delete() {
	fmt.Println("***********************删除客户***********************")
	fmt.Println("请输入要删除的客户编号（-1：退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" {
		// 执行删除
		if customer.customerService.Delete(id) {
			fmt.Println("***********************删除成功***********************")
		} else {
			fmt.Println("***********************删除失败***********************")
		}
	}
}

// 编写 update() 方法，调用 CustomerService.Update() 方法，修改客户
func (customer *customerView) update() {
	fmt.Println("***********************修改客户***********************")
	fmt.Println("请输入要删除的客户编号（-1：退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	// 输入回车，不修改
	_, selectCustomer := customer.customerService.FindById(id)
	fmt.Printf("姓名(%v): ", selectCustomer.Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别(%v): ", selectCustomer.Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v): ", selectCustomer.Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话(%v): ", selectCustomer.Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v): ", selectCustomer.Email)
	email := ""
	fmt.Scanln(&email)
	// 回车表示不修改
	if name == "" {
		name = selectCustomer.Name
	}
	if gender == "" {
		gender = selectCustomer.Gender
	}
	if age == 0 {
		age = selectCustomer.Age
	}
	if phone == "" {
		phone = selectCustomer.Phone
	}
	if email == "" {
		email = selectCustomer.Email
	}
	newCustomer := model.NewCustomer(id, name, gender, age, phone, email)
	isUpdate := customer.customerService.Update(id, newCustomer)
	if isUpdate {
		fmt.Println("*********************客户修改完成*********************")
	} else {
		fmt.Println("*********************客户修改失败*********************")
	}
}

// 退出方法
func (customer *customerView) exit() {
	fmt.Println("你确定退出吗，请输入：Y(是)/N(否)")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "N" {
			break
		}
		fmt.Println("你的输入有误，请重新输入：Y/N")
	}
	if choice == "Y" {
		customer.loop = false
	}
}

// 显示主菜单
func (customer *customerView) mainView() {
	for {
		fmt.Println("***********************客户信息管理软件***********************")
		fmt.Println("                       1、添加客户")
		fmt.Println("                       2、修改客户")
		fmt.Println("                       3、删除客户")
		fmt.Println("                       4、客户列表")
		fmt.Println("                       5、退出")
		fmt.Print("请选择(1-5): ")
		fmt.Scanln(&customer.key)
		switch customer.key {
		case "1":
			customer.add()
		case "2":
			customer.update()
		case "3":
			customer.delete()
		case "4":
			customer.list()
		case "5":
			customer.exit()
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if !customer.loop {
			break
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("你已退出客户关系管理系统")
}

func main() {
	fmt.Println("ok")
	// 初始化创建一个客户
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainView()
}
