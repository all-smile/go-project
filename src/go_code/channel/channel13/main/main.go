// 无缓冲的通道模拟4个 goroutine 间的接力比赛
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一位跑步者将计数 +1
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待游戏结束
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	// 等待接力棒
	runner := <-baton

	// 开始跑
	fmt.Printf("Runner %d Running Width Baton \n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line🏃‍♂️ \n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑

	time.Sleep(100 * time.Microsecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over🎈\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange Width Runner %d👉 \n", runner, newRunner)
	baton <- newRunner
}

/*
Runner 1 Running Width Baton
Runner 2 To The Line🏃‍♂️
Runner 1 Exchange Width Runner 2👉
Runner 2 Running Width Baton
Runner 3 To The Line🏃‍♂️
Runner 2 Exchange Width Runner 3👉
Runner 3 Running Width Baton
Runner 4 To The Line🏃‍♂️
Runner 3 Exchange Width Runner 4👉
Runner 4 Running Width Baton
Runner 4 Finished, Race Over🎈
*/
