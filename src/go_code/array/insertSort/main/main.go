// 插入排序

package main

import "fmt"

func insertSort(arr *[5]int) {

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]  // 要进行插入的值
		insertIndex := i - 1 // 插入位置的前一个下标

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}

		// 插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第 %v 次插入排序：%v \n", i, arr)
	}

	/* insertVal := arr[1]
	insertIndex := 1 - 1 // 插入位置的前一个下标

	// 从大到小
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex] // 数据后移
		insertIndex--
	}

	// 插入
	if insertIndex+1 != 1 {
		arr[insertIndex+1] = insertVal
	}
	fmt.Printf("第 %v 次插入排序：%v \n", 1, arr)

	insertVal = arr[2]
	insertIndex = 2 - 1

	// 从大到小
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	// 插入
	if insertIndex+1 != 2 {
		arr[insertIndex+1] = insertVal
	}
	fmt.Printf("第 %v 次插入排序：%v \n", 2, arr)

	insertVal = arr[3]
	insertIndex = 3 - 1

	// 从大到小
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	// 插入
	if insertIndex+1 != 3 {
		arr[insertIndex+1] = insertVal
	}
	fmt.Printf("第 %v 次插入排序：%v \n", 3, arr)

	insertVal = arr[4]
	insertIndex = 4 - 1

	// 从大到小
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	// 插入
	if insertIndex+1 != 4 {
		arr[insertIndex+1] = insertVal
	}
	fmt.Printf("第 %v 次插入排序：%v \n", 4, arr) */
}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println("arr = ", arr)
	insertSort(&arr)
	fmt.Println("排序后：", arr)
}
