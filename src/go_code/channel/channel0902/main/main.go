// 互斥锁定义同步访问资源的代码临界区
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
	// mutex 用来定义一段代码临界区, (加锁、解锁)
	mutex sync.Mutex
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
		// 同一时刻，只能允许一个 goroutine 进入这个临界区
		mutex.Lock()
		{
			// 捕获 counter 的值
			value := counter

			// 当前 goroutine 从线程退出，并放回到队列
			runtime.Gosched()

			// 增加本地 value 的值， 并将值保存回 counter
			value++
			counter = value
		}
		// 释放锁，允许其它正在等待的 goroutine 进入临界区
		mutex.Unlock()
	}
}

// Final Counter: 4
