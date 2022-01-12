// 快速排序

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(left int, right int, arr *[8000000]int) {

	// 保留 左右下标值，不参与运算, 只用来传递参数，定义运算边界，划分子问题
	L := left // 值拷贝
	R := right

	// 递归结束条件
	if left >= right {
		return
	}

	// 中轴
	pivot := arr[left]

	// 从小到大排
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right] // 移动值
		}

		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right] = arr[left] // 移动值
		}

		// 中轴复位
		if left >= right {
			arr[left] = pivot
		}
	}
	// fmt.Println(left, right, arr)

	// 划分子问题
	// 左递归
	quickSort(L, right-1, arr)
	// 右递归
	quickSort(right+1, R, arr)
}

func main() {
	// arr := [6]int{-9, 78, 0, 23, -567, 70}
	// arr := [...]int{5, 8, 2, 1, 3, 6, 2, 4, 7}

	var arr [8000000]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(90000000)
	}

	// fmt.Println("arr = ", arr)
	// quickSort(0, len(arr)-1, &arr)
	// fmt.Println("arr = ", arr)

	startTime := time.Now().Unix()
	quickSort(0, len(arr)-1, &arr)
	endTime := time.Now().Unix()

	fmt.Printf("排序耗费时间： %v 秒\n", endTime-startTime)
}

/*
arr =  [5 8 1 3 6 2 4 7]
4 4 &[4 2 1 3 5 6 8 7]
3 3 &[3 2 1 4 5 6 8 7]
2 2 &[1 2 3 4 5 6 8 7]
0 0 &[1 2 3 4 5 6 8 7]
5 5 &[1 2 3 4 5 6 8 7]
7 7 &[1 2 3 4 5 6 7 8]
arr =  [1 2 3 4 5 6 7 8]
*/
