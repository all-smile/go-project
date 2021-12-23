// recover
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
}

func test() {
	// defer + recover 捕获处理异常
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			// 有异常
			fmt.Println("test() err = ", err)
			// test() err =  assignment to entry in nil map
		}
	}()
	var testmap map[int]string
	// var testmap = make(map[int]string)
	testmap[0] = "北京" // 没有分配空间 （make） error
}

func main() {
	fmt.Println("程序开始了")
	wg.Add(2)
	// recover
	go sayHello()
	go test()

	wg.Wait()

	fmt.Println("程序结束了")
}
