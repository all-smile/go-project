package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello")
	i := 0
	for ; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println("i = ", i)
	}
	fmt.Println("i循环结束之后的值 = ", i)
	var str string = "hello,world你好"
	str2 := []rune(str)
	fmt.Println(str2)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}
	fmt.Println("\n")
	var str02 = "xiao笑笑"
	for index, val := range str02 {
		fmt.Printf("index=%d val=%c \n", index, val)
	}

	var num int = 100
	var count int
	var sum int
	for i := 0; i < num; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%d sum=%d \n", count, sum)

	var num1 int = 6
	for i := 0; i < num1; i++ {
		fmt.Printf("%v + %v = %v \n", i, num1-i, num1)
	}
	var n int = 1
	for {
		if n > 5 {
			break
		}
		fmt.Println("hello~", n)
		n++
	}
	fmt.Println(n)

	// var classNum int = 3 // 3个班级
	// var totalScore float64
	// var pNum = 3
	// for n := 1; n <= classNum; n++ {
	// 	var sumScore float64
	// 	for i := 1; i <= pNum; i++ {
	// 		var score float64
	// 		fmt.Printf("请输入第%d个班的第%d个学生的成绩 \n", n, i)
	// 		fmt.Scanln(&score)
	// 		sumScore += score
	// 	}
	// 	fmt.Printf("第%d个班级的总分是%v \n", n, sumScore)
	// 	fmt.Printf("第%d个班级的平均分是%v \n", n, sumScore/float64(pNum))
	// 	totalScore += sumScore
	// }
	// fmt.Printf("%d个班级总分数是%v \n", classNum, totalScore)
	// fmt.Printf("%d个班级平均分数是%v \n", classNum, totalScore/(float64(classNum)*float64(pNum)))
	fmt.Println("-------------------------------------")
	// 打印金字塔
	var countRow int = 5
	for i := 1; i <= countRow; i++ {
		for n := 0; n < countRow-i; n++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (i*2 - 1); j++ {
			// 空心金字塔
			if (j > 1 && j < (i*2-1)) && i != countRow {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v  ", j, i, i*j)
		}
		fmt.Println()
	}

	// 随机数生成
	// 设置种子
	var times int = 0
	fmt.Println("time Unix==", time.Now().UnixNano())
	for {
		rand.Seed(time.Now().UnixNano())
		var r int = rand.Intn(100) + 1
		fmt.Printf("%T %v \n", r, r)
		times++
		if r == 99 {
			break
		}
	}
	fmt.Printf("生成99,一共用了%v次 \n", times)

	// break label
testlabel:
	for i := 0; i < 5; i++ {
		// label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break testlabel
			}
			fmt.Println("j=", j)
		}
	}

	// 登录验证，三次机会，提示剩余次数
	var name string
	var pwd string
	var loginChance int = 3
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name == "xiao" && pwd == "666" {
			fmt.Println("登录成功")
			break
		} else {
			loginChance--
			fmt.Printf("您还有%d次机会\n", loginChance)
		}
	}
	if loginChance == 0 {
		fmt.Println("机会用完，没有登录成功，明天再试")
	}
}
