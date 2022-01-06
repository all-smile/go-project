// 数组模拟队列
// 非环队列
package main

import (
	"errors"
	"fmt"
	"os"
)

/*
1. 创建一个数组arrary，是作为队列的一个字段
2. front初始化为-1
3. rear, 表示队列尾部，初始化为-1
4. 完成队列的基本查找
AddQueue // 加入数据到队列
GetQueue // 从队列取出数据
ShowQueue // 显示队列
*/

type Queue struct {
	maxSize int
	array   [4]int // 数组模拟队列
	front   int    // 队列首(下标index), 不包含队首的元素
	rear    int    // 队列尾(下标index)
}

// 添加数据到队列 - 入队列
func (que *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if que.rear == que.maxSize-1 {
		return errors.New("queue full")
	}
	que.rear++ // rear 后移
	que.array[que.rear] = val
	fmt.Println("front = ", que.front, "rear = ", que.rear, "array = ", que.array)
	return
}

// 显示队列
func (que *Queue) ShowQueue() {
	fmt.Println("显示队列: ")
	fmt.Println("front = ", que.front, "rear = ", que.rear, "array = ", que.array)
	// front初始值是 -1， 不包含队首的元素
	for i := que.front + 1; i <= que.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, que.array[i])
	}
	fmt.Println()
}

// 出队列
func (que *Queue) GetQueue() (val int, err error) {
	// 判断队列是否是空
	if que.rear == que.front {
		return -1, errors.New("queue empty")
	}
	que.front++
	val = que.array[que.front]
	que.array[que.front] = 0
	fmt.Println("front = ", que.front, "rear = ", que.rear, "array = ", que.array)
	return val, err
}

func main() {
	queue := &Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err == nil {
				fmt.Println("加入队列成功")
			} else {
				fmt.Println(err.Error())
			}
		case "get":
			val, err := queue.GetQueue()
			if err == nil {
				fmt.Println("取到的值：", val)
			} else {
				fmt.Println(err.Error())
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}

/*
get
front =  2 rear =  3 array =  [0 0 0 4]
取到的值： 3
1. 输入 add 表示添加数据到队列
2. 输入 get 表示从队列中获取数据
3. 输入 show 表示显示队列
4. 输入 exit 表示退出
get
front =  3 rear =  3 array =  [0 0 0 0]
取到的值： 4
1. 输入 add 表示添加数据到队列
2. 输入 get 表示从队列中获取数据
3. 输入 show 表示显示队列
4. 输入 exit 表示退出
get
queue empty
*/
