// goroutine + channel 案例
package main

import (
	"fmt"
	"time"
)

// write data
func writeData(intChan chan int) {
	for i := 1; i <= 10; i++ {
		intChan <- i
		fmt.Printf("写入数据： %v \n", i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}

// read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		} else {
			time.Sleep(time.Second)
			fmt.Printf("读到数据了： %v \n", v)
		}
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)
	// time.Sleep(time.Second * 10)
	// 阻塞的方式
	for {
		v, ok := <-exitChan
		if !ok {
			break
		} else {
			fmt.Printf("结束了： %v \n", v)
		}
	}
	fmt.Println("over~")
}
