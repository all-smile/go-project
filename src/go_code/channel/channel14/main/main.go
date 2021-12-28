// 有缓冲的通道 固定数目
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的 goroutine 的数量
	taskLoad         = 10 // 要处理的工作的数量
)

var (
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 0; gr < numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有的 goroutine 退出
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down✔ \n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s \n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s \n", worker, task)
	}
}

/*
Worker: 3 : Started Task : 1
Worker: 1 : Started Task : 2
Worker: 0 : Started Task : 3
Worker: 2 : Started Task : 4
Worker: 3 : Completed Task : 1
Worker: 3 : Started Task : 5
Worker: 0 : Completed Task : 3
Worker: 0 : Started Task : 6
Worker: 1 : Completed Task : 2
Worker: 1 : Started Task : 7
Worker: 2 : Completed Task : 4
Worker: 2 : Started Task : 8
Worker: 2 : Completed Task : 8
Worker: 2 : Started Task : 9
Worker: 1 : Completed Task : 7
Worker: 1 : Started Task : 10
Worker: 2 : Completed Task : 9
Worker: 2 : Shutting Down✔
Worker: 1 : Completed Task : 10
Worker: 1 : Shutting Down✔
Worker: 3 : Completed Task : 5
Worker: 3 : Shutting Down✔
Worker: 0 : Completed Task : 6
Worker: 0 : Shutting Down✔
*/
