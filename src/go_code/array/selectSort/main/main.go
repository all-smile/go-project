// 选择排序

package main

import "fmt"

func selectSort(arr *[5]int) {

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
		fmt.Printf("第 %v 次排序: %v\n", i+1, arr)
	}

	// max := arr[0]
	// maxIndex := 0
	// for i := 0 + 1; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 		maxIndex = i
	// 	}
	// }
	// if maxIndex != 0 {
	// 	arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	// }
	// fmt.Println("第1次排序", arr)

	// max = arr[1]
	// maxIndex = 1
	// for i := 0 + 2; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 		maxIndex = i
	// 	}
	// }
	// if maxIndex != 0 {
	// 	arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
	// }
	// fmt.Println("第2次排序", arr)

	// max = arr[2]
	// maxIndex = 2
	// for i := 0 + 3; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 		maxIndex = i
	// 	}
	// }
	// if maxIndex != 0 {
	// 	arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
	// }
	// fmt.Println("第3次排序", arr)

	// max = arr[3]
	// maxIndex = 3
	// for i := 0 + 4; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 		maxIndex = i
	// 	}
	// }
	// if maxIndex != 0 {
	// 	arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
	// }
	// fmt.Println("第4次排序", arr)
}

func main() {
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println("原数组")
	fmt.Println("arr = ", arr)
	selectSort(&arr)
	fmt.Println("排序后的数组")
	fmt.Println("arr = ", arr)
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
