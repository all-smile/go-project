package model

import "fmt"

// 定义一个 account 结构体
type account struct {
	accountNo string
	pwd       string
	blance    float64
}

// 工厂模式的构造函数
func NewAccount(accountNo string, pwd string, blance float64) *account {
	// 数据校验
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不对")
		// 对于nil，一般通常指针类型和interface类型可以使用这样的返回值
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
		return nil
	}
	if blance < 20 {
		fmt.Println("余额不对")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		blance:    blance,
	}
}

// 绑定方法
// 1. 存款
func (account *account) Deposite(money float64, pwd string) {
	// 判断输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	// 校验金额
	if money < 0 {
		fmt.Println("金额不正确")
		return
	}
	account.blance += money
	fmt.Printf("存款成功，交易后余额： %v \n", account.blance)
}

// 2.取款
func (account *account) WithDraw(money float64, pwd string) {
	// 判断输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	// 校验金额
	if money < 0 || money > account.blance {
		fmt.Println("金额不正确")
		return
	}
	account.blance -= money
	fmt.Printf("取款成功，交易后余额： %v \n", account.blance)
}

// 2.查询余额
func (account *account) Query(pwd string) {
	// 判断输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	fmt.Printf("你的账号：%v, 余额： %v \n", account.accountNo, account.blance)
}
