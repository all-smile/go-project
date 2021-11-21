package main

import (
	"fmt"
	"sort"
)

func main() {
	mymap := make(map[int]int, 10)
	mymap[8] = 64
	mymap[10] = 100
	mymap[2] = 4
	mymap[5] = 25
	mymap[0] = 0

	fmt.Println("mymap = ", mymap) //  map[0:0 2:4 5:25 10:100]

	/*
		排序
		1.遍历map,把key存放在切片中
		2.对切片排序
		3.遍历切片，按照key输出
	*/
	var keyList []int
	for k, _ := range mymap {
		keyList = append(keyList, k)
	}
	fmt.Println("keyList = ", keyList)
	sort.Ints(keyList)
	fmt.Println("排序后keyList = ", keyList)

	// 按照keyList遍历map, 输出
	for _, v1 := range keyList {
		fmt.Printf("map[%v]: %v \n", v1, mymap[v1])
	}
}
