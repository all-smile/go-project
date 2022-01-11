// 栈的运用 - 计算表达式

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int     // 栈的容量
	Top    int     // 栈顶下标 默认 -1
	arr    [20]int // 数组模拟栈
}

// 入栈
func (stack *Stack) Push(val int) (err error) {
	// 栈满
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	stack.Top++
	// 放入数据
	stack.arr[stack.Top] = val
	return
}

// 出栈
func (stack *Stack) Pop() (val int, err error) {
	if stack.Top == -1 {
		fmt.Println("空栈")
		return 0, errors.New("空栈")
	}
	// 取出数据
	val = stack.arr[stack.Top]
	stack.Top--
	return val, nil
}

// 遍历栈
func (stack *Stack) List() {
	if stack.Top == -1 {
		fmt.Println("空栈")
		return
	}

	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %v \n", i, stack.arr[i])
	}

}

// 判断是运算符的函数 + - * / , 利用 ASCII 码
func (stack *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (stack *Stack) Cal(num1 int, num2 int, oper int) int {
	// 注意运算的顺序
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// 返回运算符的优先级： * / => 1 ; + - => 0
func (stack *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	// 数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	// 符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	// 表达式
	exp := "33+2*6/1-2" // 字符串本身也是切片

	// 定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	keepNum := "" // 字符串 用于拼接多位数

	// 定义索引，扫描表达式
	index := 0
	for {
		ch := exp[index : index+1] // 每次拿到一个字符， 如果数字大于1位就会出问题, 在非符号字符入栈时处理
		temp := int([]byte(ch)[0]) // 字符对应的 ASCII 码

		if operStack.IsOper(temp) {
			// 符号
			if operStack.Top == -1 {
				// 1. 符号栈为空
				operStack.Push(temp)
			} else {
				// 2. 符号栈不为空
				// 2.1 栈中的符号优先级大于等于即将入栈的符号
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result := operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					// 2.2 栈中的符号优先级小于即将入栈的符号
					operStack.Push(temp)
				}
			}
		} else {
			// 数字

			// 多位数拼接
			keepNum += ch
			// 判断扫描的下一位是不是运算符， 边界问题（表达式最后， 预防溢出）
			if index == len(exp)-1 {
				// 字符 转数字
				val, _ := strconv.Atoi(keepNum) // Atoi是ParseInt(s, 10, 0)的简写。 返回的就是 int 类型
				numStack.Push(val)
			} else {
				// 向后看一位
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.Atoi(keepNum) // Atoi是ParseInt(s, 10, 0)的简写。 返回的就是 int 类型
					numStack.Push(val)
					keepNum = ""
				}
				// 其它情况不做处理，index++, 判断下一个字符
			}

			/*
				// 字符 转数字
				// val, _ := strconv.ParseInt(ch, 10, 0) // 返回 int64, 需要转 int(val)
				val, _ := strconv.Atoi(ch) // Atoi是ParseInt(s, 10, 0)的简写。 返回的就是 int 类型
				numStack.Push(val)
			*/
		}

		// 判断扫描位置
		if index+1 == len(exp) {
			fmt.Println("扫描到最后了")
			break
		}
		index++
	}

	// 扫描完毕，之后对数栈和符号栈剩余元素处理
	for {
		if operStack.Top == -1 {
			fmt.Println("没有运算符了， 数栈中的值就是计算结果")
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result := operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}

	// 如果计算正确，数栈中只剩下一个元素，就是计算结果
	res, _ := numStack.Pop()
	fmt.Printf("表达式： %s = %v \n", exp, res)
}
