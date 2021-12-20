package main

import "fmt"

func main() {
	// 1、声明 chan 类型的数据，存放 3 个int
	var intChan chan int
	fmt.Println("intChan =", intChan) // nil

	// 2、初始化
	intChan = make(chan int, 3)
	// intChan类型： chan int, 值： 0xc00001c100 , 本身的地址： 0xc000006028
	fmt.Printf("intChan类型： %T, 值： %v , 本身的地址： %v\n", intChan, intChan, &intChan)
	// 3、向管道（channel）写入数据
	// 注意： 写入数据时，不能超过其容量，fatal error: all goroutines are asleep - deadlock!
	// 致命错误：所有goroutine都处于休眠状态-死锁！
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 12
	// 4、查看channel的长度和容量
	fmt.Printf("intChan长度： %v, 容量： %v \n", len(intChan), cap(intChan))
	// 5、读取 channel - 队列（先进先出FIFO）
	var num2 int = <-intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("intChan长度： %v, 容量： %v \n", len(intChan), cap(intChan))
	// 6、在没有使用协程的情况下，如果我们的管道数据已经全部取出，再继续取就会报 deadlock (死锁)
	num3 := <-intChan
	num4 := <-intChan
	intChan <- 000
	<-intChan
	// num5 := <-intChan // channel 里已经没有数据，不能再取了 fatal error: all goroutines are asleep - deadlock!
	fmt.Println("num3 = ", num3, "num4 = ", num4)
}
