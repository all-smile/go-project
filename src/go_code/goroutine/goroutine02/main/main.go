// channel 引出
package main

import (
	"fmt"
	"sync"
	"time"
)

// 计算 1 - 200 各个数的阶乘，然后存放到map中, 并输出

// 思路
// 1、编写一个函数 计算各个数的阶层，并放入到 map 中
// 2、启动多个协程，将统计结果放入 map 中
// 3、map 应该是全局的

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	lock sync.Mutex
)

// 编写计算阶乘的函数，并将结果放入 map
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 操作 map 前加锁
	lock.Lock()
	// 放入 map
	myMap[n] = res
	// 操作完之后解锁
	lock.Unlock()
}

func main() {
	for i := 0; i < 20; i++ {
		go test(i)
	}

	// 休眠 10s 中 ?
	time.Sleep(10 * time.Second)

	// 遍历输出结果 主线程瞬间结束，程序退出，无输出
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%v]=%v \n", i, v)
	}
	lock.Unlock()
}
