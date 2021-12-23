// goroutine ,channel 注意事项
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 管道可以声明只读或者只写
	// 1、 默认情况下，管道是双向的，可读可写
	// var chan1 chan int

	// 2、只写
	var chan2 chan<- int = make(chan<- int, 3)
	chan2 <- 20
	fmt.Println("chan2 = ", chan2)

	// 3、只读
	var chan3 <-chan int
	// num1 := <-chan3
	// chan3 <- 3 // invalid operation: chan3 <- 3 (send to receive-only type <-chan int)
	fmt.Println("chan3 = ", chan3)

	// select
	// 创建带缓冲的管道
	var intChan chan int = make(chan int, 5)
	for i := 0; i < 5; i++ {
		intChan <- i
	}
	// 创建带缓冲的管道
	var strChan chan string = make(chan string, 10)
	for i := 0; i < 10; i++ {
		strChan <- "hello" + strconv.Itoa(i)
	}

	// 传统方法遍历管道，如果不关闭会阻塞而导致 deadlock
	// 不确定什么时候关闭管道，可以使用 select 语法解决
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 中取到数据%d \n", v)
		case v := <-strChan:
			fmt.Printf("从 strChan 中取到数据%v \n", v)
		default:
			fmt.Println("都取不到，不玩了")
			return
		}
	}
	/*
		从 strChan 中取到数据hello0
		从 strChan 中取到数据hello1
		从 intChan 中取到数据0
		从 strChan 中取到数据hello2
		从 intChan 中取到数据1
		从 strChan 中取到数据hello3
		从 intChan 中取到数据2
		从 strChan 中取到数据hello4
		从 strChan 中取到数据hello5
		从 intChan 中取到数据3
		从 strChan 中取到数据hello6
		从 intChan 中取到数据4
		从 strChan 中取到数据hello7
		从 strChan 中取到数据hello8
		从 strChan 中取到数据hello9
		都取不到，不玩了
	*/

}
