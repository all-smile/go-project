// 方法练习
package main

import "fmt"

type Test struct {
}

// 编写 Test 的方法
func (t Test) print() {
	for i := 0; i < 10; i++ {
		for i := 0; i < 8; i++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 方法接收参数
func (t Test) print2(m int, n int) {
	for i := 0; i < m; i++ {
		for i := 0; i < n; i++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 方法接收参数，并返回值
func (t Test) print3(m float64, n float64) float64 {
	return m * n
}

// 指针接收者 方法里判断奇数偶数
func (t *Test) JudgeNum(n int) {
	if n%2 == 0 {
		fmt.Printf("%v是偶数 \n", n)
	} else {
		fmt.Printf("%v是奇数 \n", n)
	}
}

// 方法接收参数,并打印指定字符
func (t Test) print4(m int, n int, sign string) {
	for i := 0; i < m; i++ {
		for i := 0; i < n; i++ {
			fmt.Print(sign)
		}
		fmt.Println()
	}
}

// 实现一个四则运算的方法
type Calc struct {
	Num1 float64
	Num2 float64
}

func (calc *Calc) CalcMethod(opt string) float64 {
	var res float64
	switch opt {
	case "+":
		res = calc.Num1 + calc.Num2
	case "-":
		res = calc.Num1 - calc.Num2
	case "*":
		res = calc.Num1 * calc.Num2
	case "/":
		res = calc.Num1 / calc.Num2
	default:
		fmt.Println("运算符输入有误")
	}
	return res
}

type Person struct {
	Name string
}

// 实现打印人名的函数
// 普通函数，参数是值类型，传递的参数也必须是值类型；参数是指针类型，必须传递地址。
func printName(p Person) {
	fmt.Println("printName函数 人名 = ", p.Name) // tom
}
func printName2(p *Person) {
	fmt.Println("*printName函数 人名 = ", p.Name) // tom
}

// 值接收者Person的方法
func (p Person) printName() {
	p.Name = "jack"
	fmt.Println("Person 方法 人名 = ", p.Name) // jack
}

// 指针接收者Person的方法
func (p *Person) printName2() {
	p.Name = "bill"
	fmt.Println("*Person 方法 人名 = ", p.Name) // bill
}

func main() {
	var t Test
	t.print()

	// 方法传值
	var t2 Test
	t2.print2(2, 5)

	// 方法传值并返回
	var t3 Test
	multiply := t3.print3(2.2, 5.5)
	fmt.Println("multiply = ", multiply)

	// 实现方法判断奇数还是偶数
	var t4 Test
	t4.JudgeNum(4)
	(&t4).JudgeNum(5)

	// 打印指定字符
	var t5 Test
	t5.print4(3, 6, "$")

	// 四则运算
	var calc Calc = Calc{
		Num1: 10,
		Num2: 2,
	}
	res := calc.CalcMethod("/")
	fmt.Println("res = ", res)

	// 方法和函数的区别
	p := Person{"tom"}
	// 调用函数
	printName(p)
	printName2(&p) // 函数的形参类型必须跟实参类型一致，都为值类型，或者都为指针类型

	// 调用方法 (值接收者)
	// 值接收者 不论调用者变量是值还是地址，都不会改变原始数据
	p.printName()
	fmt.Println("值接收者1 p.Name = ", p.Name) // tom
	(&p).printName()
	fmt.Println("值接收者2 p.Name = ", p.Name) // tom

	// 调用方法 (指针接收者)
	// 指针接收者, 会改变原始数据
	p.printName2()
	fmt.Println("指针接收者1 p.Name = ", p.Name) // bill
	(&p).printName2()
	fmt.Println("指针接收者2 p.Name = ", p.Name) // bill
}
