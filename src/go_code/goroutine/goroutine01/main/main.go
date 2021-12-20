// 协程 goroutinie 案例
package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
主线程开启一个 goroutine 该协程每隔1秒输出 “hello goroutine”
主线程每隔1秒输出 “hello golang” 输出10次后退出程序
主线程和 goroutine 同时执行
*/

// 编写函数
func test() {
	for i := 1; i <= 2; i++ {
		fmt.Println("test() hello goroutine: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	// test() // 执行完 test() 函数之后才会继续执行下面的程序
	go test() // 开启一个协程: goroutine
	for i := 1; i <= 5; i++ {
		fmt.Println("main() hello golang: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	cpuNum := runtime.NumCPU()
	fmt.Println("逻辑cpu个数： ", cpuNum) // 逻辑cpu个数：  12
	runtime.GOMAXPROCS(cpuNum - 1)
}
