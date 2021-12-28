// atomic 包 原子函数加锁
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int64
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
		// 安全的对 counter + 1
		// AddInt64() 同步整型值的加法
		atomic.AddInt64(&counter, 1)

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}

// Final Counter: 4
