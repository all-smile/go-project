// 选择排序

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectSort(arr *[80000]int) {

	for i := 0; i < len(arr)-1; i++ {
		max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
		// fmt.Printf("第 %v 次排序: %v\n", i+1, arr)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// arr := [5]int{10, 34, 19, 100, 80}

	var arr [80000]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(9000000)
	}

	fmt.Println("原数组")
	// fmt.Println("arr = ", arr)

	startTime := time.Now().Unix()
	selectSort(&arr)
	endTime := time.Now().Unix()

	fmt.Printf("排序耗费时间： %v 秒\n", endTime-startTime)

	fmt.Println("排序后的数组")
	// fmt.Println("arr = ", arr)
}

/*
原数组
arr =  [10 34 19 100 80]
第 1 次排序: &[100 34 19 10 80]
第 2 次排序: &[100 80 19 10 34]
第 3 次排序: &[100 80 34 10 19]
第 4 次排序: &[100 80 34 19 10]
排序后的数组
arr =  [100 80 34 19 10]
*/
