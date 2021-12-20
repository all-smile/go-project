package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{} = make(chan interface{}, 10)
	allChan <- 10
	cat1 := Cat{Name: "tom", Age: 2}
	cat2 := Cat{Name: "tom02", Age: 3}
	allChan <- cat1
	allChan <- cat2
	allChan <- "jack"
	// 取出
	item := <-allChan
	fmt.Printf("item类型：%T, 值： %v \n", item, item)
	// item类型：int, 值： 10
	// 类型断言
	newItem, ok := item.(Cat)
	fmt.Printf("newItem 类型：%T, 值： %v, ok = %v", newItem, newItem, ok)
	// newItem 类型：main.Cat, 值： { 0}, ok = false
}
