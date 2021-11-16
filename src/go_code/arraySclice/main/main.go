package main

import "fmt"

func main() {
	// 求3个数的平均数
	/* n1 := 10.289
	n2 := 9.98
	n3 := 10.645
	var total = n1 + n2 + n3
	var avgNum = total / 3
	_avgNum := fmt.Sprintf("%.2f", avgNum) // 保留两位小数

	fmt.Println("avgNum = ", avgNum)   // avgNum =  10.304666666666666
	fmt.Println("_avgNum = ", _avgNum) // avgNum =  10.304666666666666 */

	// 数组
	/* var list [5]float64
	list[0] = 10.289
	list[1] = 9.98
	list[2] = 10.645
	list[3] = 10
	list[4] = 9
	fmt.Printf("list 类型: %T, 值： %v \n", list, list)
	var t float64
	for i := 0; i < len(list); i++ {
		t += list[i]
	}
	fmt.Println("t = ", t/3)
	var avg = t / float64(len(list))
	var _avg = fmt.Sprintf("%.2f", avg)
	fmt.Println("_avg = ", _avg) */

	/* var arrInt [3]int32
	fmt.Printf("arrInt 类型: %T, 值： %v \n", arrInt, arrInt)
	fmt.Printf("arrInt 的内存地址 = %p \n", &arrInt)
	for i := 0; i < len(arrInt); i++ {
		fmt.Printf("list[%v]的地址是：%p, 值：%v \n", i, &arrInt[i], arrInt[i])
	} */

	// 初始化数组
	/* var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1 = ", arr1)
	var arr2 = [3]int{10, 20, 30}
	fmt.Println("arr2 = ", arr2)
	var arr3 = [...]int{0, 1, 2}
	fmt.Println("arr3 = ", arr3)
	var arr4 = [...]int{1: 100, 3: 500, 0: 11}
	fmt.Println("arr4 = ", arr4) */

	var arr02 = [3]string{"luffy01", "luffy02", "luffy03"}
	for index, value := range arr02 {
		fmt.Println("index = ", index, "value = ", value)
	}

}
