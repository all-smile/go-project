// 该程序展示竞争状态（实际中不希望出现这种状态）
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

// incCounter 增加报里 counter 变量的值
func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		// 捕获 counter 的值
		value := counter

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()

		// 增加本地 value 的值， 并将值保存回 counter
		value++
		counter = value

	}
}

// 官方示例 结果是
// Final Counter: 2

// 这里结果是： Final Counter: 4
// go version go1.17.2 windows/amd64
