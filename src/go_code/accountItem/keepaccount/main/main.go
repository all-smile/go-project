// 家庭收支记账软件 - 面向过程实现
package main

import (
	"fmt"
)

func main() {
	// 声明变量，接收用户输入菜单项
	key := ""
	// 声明变量，控制是否退出for
	loop := true

	// 定义账户余额
	blance := 1000.0
	// 每次收支的金额
	money := 0.0
	//收支说明
	note := ""
	// 定义一个变量，表示是否存在收支记录
	flag := false
	// 收支详情，使用字符串记录
	// 当有收支记录的时候，只需要对 details 拼接即可
	details := "收支\t账户金额\t收支金额\t说明"

	// 显示主菜单
	for {
		fmt.Println("***********************家庭收支记账软件***********************")
		fmt.Println("                       1、收支明细")
		fmt.Println("                       2、登记收入")
		fmt.Println("                       3、登记支出")
		fmt.Println("                       4、退出软件")
		fmt.Print("请选择(1-4): ")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if !flag {
				fmt.Println("目前为止，你还没有收支记录，来一笔吧🎈")
			} else {
				fmt.Println("***********************当前收支明细记录***********************")
				fmt.Println(details)
			}
		case "2":
			fmt.Println("***********************登记收入***********************")
			fmt.Println("本次收入金额: ")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明: ")
			fmt.Scanln(&note)
			blance += money
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", blance, money, note)
			flag = true
		case "3":
			fmt.Println("***********************登记支出***********************")
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)
			if money > blance {
				fmt.Println("余额不足啦")
				break
			}
			blance -= money
			fmt.Println("本次收入说明: ")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", blance, money, note)
			flag = true
		case "4":
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
				loop = false
			}
		}
		if !loop {
			break
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("你退出了家庭记账软件🤞")
}
