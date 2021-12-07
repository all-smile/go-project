// 面向对象的方式实现 家庭记账软件
// 封装、继承、多态
package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明必须的字段

	// 声明字段，接收用户输入菜单项
	key string
	// 声明字段，控制是否退出for
	loop bool

	// 定义账户余额
	blance float64
	// 每次收支的金额
	money float64
	//收支说明
	note string
	// 定义一个字段，表示是否存在收支记录
	flag bool
	// 收支详情，使用字符串记录
	// 当有收支记录的时候，只需要对 details 拼接即可
	details string
	// details := "收支\t账户金额\t收支金额\t说明"
}

// 编写一个工厂模式的构造函数，返回 FamilyAccount
func NewFamilyAccount() *FamilyAccount {
	// 给定初始值
	return &FamilyAccount{
		key:     "",
		loop:    true,
		blance:  1000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

// 将每个菜单的功能封装成一个方法
// 明细
func (account *FamilyAccount) showDetails() {
	if !account.flag {
		fmt.Println("目前为止，你还没有收支记录，来一笔吧🎈")
	} else {
		fmt.Println("***********************当前收支明细记录***********************")
		fmt.Println(account.details)
	}
}

// 登记收入方法
func (account *FamilyAccount) income() {
	fmt.Println("***********************登记收入***********************")
	fmt.Println("本次收入金额: ")
	fmt.Scanln(&account.money)
	fmt.Println("本次收入说明: ")
	fmt.Scanln(&account.note)
	account.blance += account.money
	account.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", account.blance, account.money, account.note)
	account.flag = true
}

// 登记支出方法
func (account *FamilyAccount) pay() {
	fmt.Println("***********************登记支出***********************")
	fmt.Println("本次支出金额")
	fmt.Scanln(&account.money)
	if account.money > account.blance {
		fmt.Println("余额不足啦")
		// break
	}
	account.blance -= account.money
	fmt.Println("本次收入说明: ")
	fmt.Scanln(&account.note)
	account.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", account.blance, account.money, account.note)
	account.flag = true
}

// 退出方法
func (account *FamilyAccount) exit() {
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
		account.loop = false
	}
}

// 给 FamilyAccount 结构体绑定方法
// 显示主菜单
func (account *FamilyAccount) MainMenu() {
	for {
		fmt.Println("***********************家庭收支记账软件***********************")
		fmt.Println("                       1、收支明细")
		fmt.Println("                       2、登记收入")
		fmt.Println("                       3、登记支出")
		fmt.Println("                       4、退出软件")
		fmt.Print("请选择(1-4): ")
		fmt.Scanln(&account.key)
		switch account.key {
		case "1":
			account.showDetails()
		case "2":
			account.income()
		case "3":
			account.pay()
		case "4":
			account.exit()
		}
		if !account.loop {
			break
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("你退出了家庭记账软件🤞")
}
