// 面向对象方式完成调用
package main

import (
	_ "fmt"
	"go_code/accountItem/familyAccount/utils"
)

func main() {
	// account := utils.NewFamilyAccount()
	// fmt.Println("account = ", *account)
	//  true 1000 0  false 收支    账户金额        收支金额        说明}
	utils.NewFamilyAccount().MainMenu()
}
