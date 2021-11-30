package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// byte类型的数组，长度为26，存放英文大写字母 A-Z
	// 字符可以参与运算： 'A' + 1 = 'B' // ascii
	/* var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i)
		fmt.Printf("%c ", myChars[i])
	}
	fmt.Println("myChars = ", myChars) */

	// 求最大值, 并输出对应下标
	/* var list = [5]int{1, 0, -5, 6, 2}
	var maxV int = list[0]
	var maxIndex int = 0
	for i := 0; i < len(list); i++ {
		if maxV < list[i] {
			maxV = list[i]
			maxIndex = i
		}
	}
	fmt.Println("list = ", list)
	fmt.Println("maxV = ", maxV, "maxIndex = ", maxIndex) */

	// 求平均值
	/* var list01 [5]int = [...]int{1, -1, 20, 40, 61}
	sum := 0
	for _, v := range list01 {
		sum += v
	}

	fmt.Printf("sum = %v, list01平均值：%v \n", sum, sum/len(list01))
	// 平均值保留小数
	fmt.Printf("sum = %v, list01平均值：%v \n", sum, float64(sum)/float64(len(list01))) */

	// 随机生成5个数，并反转打印
	var intList [5]int
	len := len(intList)
	// 为了每次生成的随机数不一样，需要设置种子 Seed, 通过时间戳的方式生成种子，以保证每次都不一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intList[i] = rand.Intn(100)
	}
	fmt.Println("intList = ", intList) // intList =  [81 87 47 59 81]
	// 反转
	// for i := len - 1; i >= 0; i-- {
	// 	fmt.Printf("intList[%v] = %v \n", i, intList[i])
	// }
	fmt.Println("交换前：", intList)
	var tmp int
	for i := 0; i < len/2; i++ {
		tmp = intList[len-1-i]
		intList[len-1-i] = intList[i]
		intList[i] = tmp
	}
	fmt.Println("交换后：", intList)
}
