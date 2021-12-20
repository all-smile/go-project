// channel 的关闭和遍历
package main

import "fmt"

func main() {
	// 1、关闭
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// intChan <- 300
	// 关闭之后不能写入 panic: send on closed channel
	fmt.Println("intChan = ", intChan) // intChan =  0xc000118080
	// 关闭之后可以读取
	n1 := <-intChan
	fmt.Println("n1 = ", n1) // n1 =  100
	// 2、遍历
	intChan2 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan2 <- i // 放入100个数据
	}
	// 常规for循环只能遍历取出50个数据 - 因为每取一个，channel 的长度就 -1
	/* for i := 1; i <= len(intChan2); i++ {
		fmt.Printf("第%v个值：%v \n", i, <-intChan2)
	} */
	// for range 遍历, 没有 index
	close(intChan2)
	// 如果遍历时，没有关闭channel, 就会出现死锁 deadlock; 关闭后，则会正常遍历，遍历完之后，退出程序
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
