// 约瑟夫问题 (丢手绢)， 单向环形链表

package main

import "fmt"

type Boy struct {
	No   int
	Next *Boy
}

// 编写函数 构成 单向环形链表
// num 表示小孩的个数
// *Boy 返回该
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	for i := 1; i <= num; i++ {
		// 生成 boy
		boy := &Boy{
			No: i,
		}

		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

// 显示单向环形链表
func show(first *Boy) {
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}

	curBoy := first
	for {
		fmt.Printf("[编号： %v] => ", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
	fmt.Println()
}

// 游戏开始了， 结束后 first 重新赋值， 状态归位（否则会引发问题）， 以便后续操作
func PlayGame(first *Boy, startNo int, countNum int) *Boy {
	if first.Next == nil {
		fmt.Println("空链表")
		return first
	}

	// 定义辅助指针, 指向最后一个节点
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	// 定位到初始节点 赋给first； tail 跟着移动到 first 后面
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// for循环处理方式
	// for {
	// 	// 找到要删除的节点,
	// 	for i := 1; i <= countNum-1; i++ {
	// 		first = first.Next
	// 		tail = tail.Next
	// 	}

	// 	// 执行删除
	// 	fmt.Printf("编号： %v 出列\n", first.No)
	// 	first = first.Next
	// 	tail.Next = first
	// 	// 可以编写回调函数

	// 	if first == tail {
	// 		break
	// 	}
	// }

	// 以下是递归调用的方式
	first = delNode(first, tail, countNum)
	fmt.Printf("编号： %v 是最后一个\n", first.No)
	return first
}

// 递归调用
func delNode(first *Boy, tail *Boy, countNum int) *Boy {
	// 1. 找到要删除的节点,
	for i := 1; i <= countNum-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 2. 执行删除
	fmt.Printf("编号： %v 出列\n", first.No)
	first = first.Next
	tail.Next = first

	// 向临界条件逼近 first.Next == first, 或者 first == tail
	if first.Next == first {
		// 表示已经剩最后一个了,
		return first
	}
	boy := delNode(first, tail, countNum)
	// 返回最后剩下的
	return boy
}

func main() {
	first := AddBoy(5)
	show(first)
	fmt.Println("开始...")
	first = PlayGame(first, 2, 3)
	fmt.Println("------ 结束了， 显示剩余的：------")
	show(first)
}

/*
[编号： 1] => [编号： 2] => [编号： 3] => [编号： 4] => [编号： 5] =>
开始...
编号： 4 出列
编号： 2 出列
编号： 1 出列
编号： 3 出列
编号： 5 是最后一个
------ 结束了， 显示剩余的：------
[编号： 5] =>
*/
