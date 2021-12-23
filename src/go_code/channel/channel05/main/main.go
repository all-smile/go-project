// goroutine + channel 求素数
package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 800; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				// 不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
	fmt.Println("该协程退出 primeNum")
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	// start time
	start := time.Now().Unix()

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// end time
		end := time.Now().Unix()
		fmt.Println("协程耗费时间：", end-start)
		close(primeChan)
	}()

	for v := range primeChan {
		fmt.Println("素数： ", v)
	}
	fmt.Println("主线程退出")
}
