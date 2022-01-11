// 栈

package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    // 栈的容量
	Top    int    // 栈顶下标 默认 -1
	arr    [5]int // 数组模拟栈
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

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.List()
	fmt.Println("入栈")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	// if err != nil {
	// 	fmt.Println("push 出错了", err)
	// }

	fmt.Println("显示栈1")
	stack.List()
	fmt.Println("出栈")
	val, err := stack.Pop()
	if err == nil {
		fmt.Printf("出栈数据： %v \n", val)
	}
	val, err = stack.Pop()
	if err == nil {
		fmt.Printf("出栈数据： %v \n", val)
	}
	val, err = stack.Pop()
	if err == nil {
		fmt.Printf("出栈数据： %v \n", val)
	}
	val, err = stack.Pop()
	if err == nil {
		fmt.Printf("出栈数据： %v \n", val)
	}
	val, err = stack.Pop()
	if err == nil {
		fmt.Printf("出栈数据： %v \n", val)
	}
	fmt.Println("显示栈2")
	stack.List()
}

/*
空栈
入栈
stack full
显示栈1
arr[4] = 5
arr[3] = 4
arr[2] = 3
arr[1] = 2
arr[0] = 1
出栈
出栈数据： 5
出栈数据： 4
出栈数据： 3
出栈数据： 2
出栈数据： 1
显示栈2
空栈
*/
