package main

import "fmt"

func main() {
	// 定义数组
	var intArr [5]int = [...]int{10, 20, 30, 40, 50}
	fmt.Println("intArr = ", intArr) // [10 20 30 40 50]
	// 以intArr为底层数组，创建切片
	slice := intArr[2:4]                   // 长度为 4-2=2 ， 容量为 intArr底层数组的容量5 - 2 = 3
	fmt.Println("slice 元素 = ", slice)      // [30 40]
	fmt.Println("slice 长度 = ", len(slice)) // 2
	fmt.Println("slice 容量 = ", cap(slice)) // 3
	// 修改切片元素，会作用到底层数组
	slice[1] = 21
	fmt.Println("修改后slice 元素 = ", slice) // [30 21]
	fmt.Println("修改后intArr = ", intArr)  //  [10 20 30 21 50]

	// nil切片
	var slice01 []int
	fmt.Printf("slice01 类型： %T, 值： %v \n", slice01, slice01) // slice01 类型： []int, 值： []
	fmt.Println(slice01 == nil)                              // true

	var slice02 []int = make([]int, 3)
	fmt.Printf("slice02 类型： %T, 值： %v \n", slice02, slice02)

	// 空切片
	slice03 := make([]int, 0)
	fmt.Printf("slice03 类型： %T, 值： %v \n", slice03, slice03)

	// 常规for循环遍历切片
	var slice04 []int = []int{10, 20, 30}
	for i := 0; i < len(slice04); i++ {
		fmt.Printf("slice04[%v] = %v, 地址：%p \n", i, slice04[i], &slice04[i])
	}
	// range创建每个元素的副本，而不是直接返回对该元素的引用
	// 迭代返回变量是一个迭代过程中根据切片依次赋值的新变量
	// 所以v的地址总是相同的
	// 要想获取每个元素的地址，可以使用切片变量和索引值
	for i, v := range slice04 {
		fmt.Printf("range - slice04[%v] = %v, 地址：%p \n", i, v, &v)
	}

	// 切片，修改
	var slice05 []int = []int{5, 6, 7, 8, 9}
	fmt.Println("slice05 = ", slice05) // [5 6 7 8 9]
	slice06 := slice05[2:4]
	fmt.Println("slice06 = ", slice06) // [7 8]
	slice06[0] = 22
	fmt.Println("修改slice06后： slice05 = ", slice05) // slice05也发生改变 [5 6 22 8 9]
	fmt.Println("修改slice06后： slice06 = ", slice06) // [22 8]

	// append 追加
	var slice07 = make([]int, 3, 5)                               // 长度3，容量5
	fmt.Println("slice07 = ", slice07, "slice07容量", cap(slice07)) // slice07 =  [0 0 0] slice07容量 5
	// 追加元素
	slice07 = append(slice07, 2)
	fmt.Println("slice07 = ", slice07) // slice07 =  [0 0 0 2]
	// 追加切片 slice...
	slice07 = append(slice07, slice07...)
	fmt.Println("slice07 = ", slice07, "追加后slice07容量", cap(slice07)) // slice07 =  [0 0 0 2 0 0 0 2] 追加后slice07容量 10

	// 拷贝切片
	var a []int = []int{1, 2, 3}
	var b = make([]int, 5) // 长度 容量都是5
	c := copy(a, b)        // 返回被复制的元素个数，
	fmt.Println("c = ", c)
	fmt.Println("a = ", a, "容量=", cap(a)) // a =  [1 2 3] 容量= 3
	fmt.Println("b = ", b, "容量=", cap(b)) // b =  [1 2 3 0 0] 容量= 5
}
