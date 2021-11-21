package main

import "fmt"

func main() {
	var arr [4][6]int
	fmt.Println("arr = ", arr) // arr =  [[0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0]]
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][2] = 3
	fmt.Println("arr = ", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("第%v个数组 = ", i+1)
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
	/*
		第1个数组 = 0 0 0 0 0 0
		第2个数组 = 0 0 1 0 0 0
		第3个数组 = 0 2 3 0 0 0
		第4个数组 = 0 0 0 0 0 0
	*/

	// 二维数组内存探究
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println("arr2 = ", arr2) // arr2 =  [[0 0 0] [0 10 0]]

	// 0xc00000a3f0 + 3(元素个数) * 8(int占8个字节) = 0xc00000a408

	fmt.Printf("arr2[0]的地址： %p\n", &arr2[0]) // 0xc00000a3f0
	fmt.Printf("arr2[1]的地址： %p\n", &arr2[1]) // 0xc00000a408

	// 上面的地址就是数组元素首地址
	fmt.Println("arr2[0][0]的地址： ", &arr2[0][0]) // 0xc00000a3f0
	fmt.Println("arr2[1][0]的地址： ", &arr2[1][0]) // 0xc00000a408

	// 声明并初始化
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3 = ", arr3) // arr3 =  [[1 2 3] [4 5 6]]

	fmt.Println()

	// 遍历
	var arr4 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("第%v层", i)
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Printf("%v ", arr4[i][j])
		}
		fmt.Println()
	}
	// for range
	for i, v := range arr4 {
		fmt.Printf("i = %v, v = %v\n", i, v)
		for j, v2 := range v {
			/*
				range创建每个元素的副本，而不是直接返回对该元素的引用
				迭代返回变量是一个迭代过程中根据切片依次赋值的新变量
				所以v的地址总是相同的
				要想获取每个元素的地址，可以使用切片变量和索引值
			*/
			fmt.Printf("arr4[%v][%v] = %v, 地址：%v \n", i, j, v2, &v2)
		}
		fmt.Println()
	}
}
