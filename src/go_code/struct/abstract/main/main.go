package main

import "fmt"

// 定义一个 Account 结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 绑定方法
// 1. 存款
func (account *Account) Deposite(money float64, pwd string) {
	// 判断输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	// 校验金额
	if money < 0 {
		fmt.Println("金额不正确")
		return
	}
	account.Balance += money
	fmt.Printf("存款成功，交易后余额： %v \n", account.Balance)
}

// 2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	// 判断输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	// 校验金额
	if money < 0 || money > account.Balance {
		fmt.Println("金额不正确")
		return
	}
	account.Balance -= money
	fmt.Printf("取款成功，交易后余额： %v \n", account.Balance)
}

// 2.查询余额
func (account *Account) Query(pwd string) {
	// 判断输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	fmt.Printf("你的账号：%v, 余额： %v \n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "gs123456",
		Pwd:       "666666",
		Balance:   100.0,
	}
	// 查询余额
	account.Query("666666")
	// 存款
	account.Deposite(100.2, "666666")
	// 取款
	account.WithDraw(50, "666666")
}
