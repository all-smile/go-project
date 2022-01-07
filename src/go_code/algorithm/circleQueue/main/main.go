// 数组模拟环形队列(有缺口)
package main

import (
	"errors"
	"fmt"
	"os"
)

/*
分析思路:
1. 什么时候表示队列满
 (tail + 1) % maxSize == head
2. tail == head 表示空
3. 初始化时， tail = 0 head = 0
4. 怎么统计该队列有多少个元素
(tail + maxSize - head ) % maxSize
*/

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // 队首， 初始 0
	tail    int // 队尾， 初始 0
}

// 判断环形队列是否已满
func (que *CircleQueue) IsFull() bool {
	fmt.Println("IsFull --- head", que.head, "tail", que.tail)
	return (que.tail+1)%que.maxSize == que.head
}

// 判断环形队列是否为空
func (que *CircleQueue) IsEmpty() bool {
	return que.tail == que.head
}

// 环形队列所有元素
func (que *CircleQueue) Size() int {
	return (que.tail + que.maxSize - que.head) % que.maxSize
}

// 入队列
func (que *CircleQueue) Push(val int) (err error) {
	// 先判断队列是否已满
	if que.IsFull() {
		return errors.New("CircleQueue full")
	}
	que.array[que.tail] = val
	fmt.Println("head = ", que.head, "tail = ", que.tail, "array = ", que.array)
	// que.tail = (que.tail + 1) % que.maxSize
	que.tail = (que.tail + 1) % len(que.array)
	// que.tail++ // tail 后移
	return
}

// 出队列
func (que *CircleQueue) Pop() (val int, err error) {
	// 判断队列是否是空
	if que.IsEmpty() {
		return 0, errors.New("CircleQueue empty")
	}
	val = que.array[que.head]
	que.array[que.head] = 0
	fmt.Println("Pop --- head = ", que.head, "tail = ", que.tail, "array = ", que.array)
	que.head = (que.head + 1) % que.maxSize
	return val, err
}

// 显示队列
func (que *CircleQueue) ShowQueue() {
	fmt.Println("显示队列: ")
	size := que.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	tempHead := que.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, que.array[tempHead])
		tempHead = (tempHead + 1) % que.maxSize
	}
	fmt.Println()
	fmt.Println("head = ", que.head, "tail = ", que.tail, "array = ", que.array)
	fmt.Println()
}

func main() {
	circleQueue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入 add 表示添加数据到队列")
		fmt.Println("2. 输入 get 表示从队列中获取数据")
		fmt.Println("3. 输入 show 表示显示队列")
		fmt.Println("4. 输入 exit 表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入要入队列的数")
			fmt.Scanln(&val)
			err := circleQueue.Push(val)
			if err == nil {
				fmt.Println("加入队列成功")
			} else {
				fmt.Println(err.Error())
			}
		case "get":
			fmt.Println("取出")
			val, err := circleQueue.Pop()
			if err == nil {
				fmt.Println("取到的值：", val)
			} else {
				fmt.Println(err.Error())
			}
		case "show":
			circleQueue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}

/*
get
取出
Pop --- head =  1 tail =  0 array =  [0 0 3 4 5]
取到的值： 2
1. 输入 add 表示添加数据到队列
2. 输入 get 表示从队列中获取数据
3. 输入 show 表示显示队列
4. 输入 exit 表示退出
add
请输入要入队列的数
6
IsFull --- head 2 tail 0
head =  2 tail =  0 array =  [6 0 3 4 5]
加入队列成功
1. 输入 add 表示添加数据到队列
2. 输入 get 表示从队列中获取数据
3. 输入 show 表示显示队列
4. 输入 exit 表示退出
show
显示队列:
arr[2]=3        arr[3]=4        arr[4]=5        arr[0]=6
head =  2 tail =  1 array =  [6 0 3 4 5]
*/
