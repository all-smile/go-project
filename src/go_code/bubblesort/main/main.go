package main

import "fmt"

func sortBubble(arr *[5]int) {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// (*arr)[j] > (*arr)[j+1] 从小到大 交换位置
				// (*arr)[j] < (*arr)[j+1] 从大到小 交换位置
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
		fmt.Printf("第%v轮排序后的结果： %v \n", i+1, *arr)

	}
	/* for j := 0; j < len(arr)-1; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换位置
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第1轮排序后的结果： ", *arr)

	for j := 0; j < len(arr)-2; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换位置
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第2轮排序后的结果： ", *arr)

	for j := 0; j < len(arr)-3; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换位置
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第3轮排序后的结果： ", *arr)

	for j := 0; j < len(arr)-4; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			// 交换位置
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第4轮排序后的结果： ", *arr) */
}

func main() {
	var arr [5]int = [5]int{24, 69, 80, 57, 13}
	fmt.Println("排序前 arr = ", arr)
	// 冒泡排序
	sortBubble(&arr)
	fmt.Println("排序后 arr = ", arr)
}
