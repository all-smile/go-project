// 无缓冲的通道模拟网球比赛
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)

	// 启动两个选手
	go player("Nadal", court)
	go player("Dav", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won \n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed❗ \n", name)

			// 关闭通道，表示输了
			close(court)
			return
		}

		// 显示击球数，并将击球数 +1
		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		// 将球击向对手
		court <- ball
	}
}

/*
Player Dav Hit 1
Player Nadal Hit 2
Player Dav Missed❗
Player Nadal Won
*/
