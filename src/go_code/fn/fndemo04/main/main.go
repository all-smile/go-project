package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 10
	fmt.Println("test n1=", n1)
}

func test02(n2 *int) {
	*n2 = *n2 + 10
	fmt.Println("test02 n2=", *n2)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func getSums(n1 int, n2 int, n3 int) int {
	return n1 + n2 + n3
}

func sumSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 自定义函数数据类型， 相当于起了一个别名
type mySum func(int, int) int

func testFn(fnVar mySum, num1 int, num2 int) int {
	type myInt int
	var num3 myInt = 2
	fmt.Printf("num3的类型：%T 值：%v \n", num3, num3)
	return fnVar(num1, num2)
}

// func testFn(fnVar func(int, int) int, num1 int, num2 int) int {
// 	return fnVar(num1, num2)
// }

// 参数个数可变
func sumV2(n1 int, args ...int) int {
	sum := n1
	fmt.Printf("args类型是：%T\n", args)
	fmt.Printf("args[0]类型是：%T 值：%v\n", args[0], args[0])
	// 遍历args切片
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	n1 := 20
	n2 := 2
	test(n1)
	test02(&n2)
	fmt.Println("main n1=", n1)
	fmt.Println("main n2=", n2)

	sumFn := getSum
	res := sumFn(10, 20)
	fmt.Printf("%T %v\n", res, res)
	fmt.Printf("sumFn %T \n", sumFn)

	sumsFn := getSums
	result := sumsFn(10, 20, 30)
	fmt.Printf("result : %T %v\n", result, result)
	fmt.Printf("sumsFn类型：%T \n", sumFn)

	// 函数类型形参
	total := testFn(sumFn, 1, 2)
	fmt.Println("total==", total)

	// 自定义数据类型
	type myInt int
	var num1 myInt = 2
	var num2 int = int(num1)
	fmt.Printf("num1的类型：%T 值：%v \n", num1, num1)
	fmt.Printf("num2的类型：%T 值：%v \n", num2, num2)

	// 返回参数命名
	sum, sub := sumSub(9, 8)
	fmt.Println("sum=", sum, "sub=", sub)

	// 参数可变
	total02 := sumV2(1, 2, 3, 4)
	fmt.Println("total02=", total02)
}
