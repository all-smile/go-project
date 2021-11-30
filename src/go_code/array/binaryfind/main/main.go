package main

import "fmt"

// 二分查找函数
// 递归， 边界条件
func binaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 找不到的情况
	// 退出递归的条件
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	// 递归条件
	// 先找到中间下标，折半
	mid := (leftIndex + rightIndex) / 2
	if (*arr)[mid] > findVal {
		// findVal 在左边 leftIndex - (mid - 1)
		binaryFind(arr, leftIndex, mid-1, findVal)
	} else if (*arr)[mid] < findVal {
		// findVal 在右边 (mid + 1) - rightIndex
		binaryFind(arr, mid+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Printf("找到了，下标为%v", mid)
	}
}

func main() {
	// 有序数组
	var arr [6]int = [6]int{1, 8, 10, 89, 1000, 1234}
	fmt.Println("arr = ", arr)
	// 二分查找/折半查找
	binaryFind(&arr, 0, len(arr), 8)
}
