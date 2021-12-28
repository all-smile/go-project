// goroutine + channel 用共享的方式，消除竞争状态
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter chan int = make(chan int, 1) // 缓冲1个数据的通道，goroutine 更新其值
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	counter <- 0

	wg.Wait()
	close(counter)
	for v := range counter {
		fmt.Println("Final Counter:=", v)
	}
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加报里 counter 变量的值
func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {

		// 从通道取值 counter 的值
		value, ok := <-counter
		if !ok {
			fmt.Println("err")
			return
		}

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()

		// 增加本地 value 的值， 写入 通道 counter
		value++
		counter <- value
	}
}

// Final Counter:= 4
// Final Counter: 0xc00001a0e0
