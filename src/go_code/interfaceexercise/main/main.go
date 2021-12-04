// 接口强大功能初体验
// 函数式编程，只需要实现方法，内部会根据方法的实现处理
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1、声明 Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// 2、声明 Hero 结构体切片，动态增长
type HeroSlice []Hero

// 3、实现 Sort() 接口接收者需要的三个方法 Len Less Swap
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less() 方法决定使用什么标准排序
// 按 hero 的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	// 按照 Name 排序
	// return hs[i].Name < hs[j].Name
}

// 交换顺序
func (hs HeroSlice) Swap(i, j int) {
	/* temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp */
	// 以下是Go语法交换顺序
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	// 切片排序
	// 先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	fmt.Println("intSlice = ", intSlice) // [0 -1 10 7 90]
	// 1、冒泡排序
	// 2、使用系统提供的方法
	// func Ints(a []int)
	sort.Ints(intSlice)                       // Ints函数将a排序为递增顺序。
	fmt.Println("sort intSlice = ", intSlice) // [-1 0 7 10 90]

	// 结构体切片排序
	// 1、冒泡排序
	// 2、使用系统提供的方法
	/*
		sort包
		func Sort(data Interface)
		Sort排序data。它调用1次 data.Len 确定长度，调用O(n*log(n))次 data.Less 和 data.Swap
		本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。
	*/
	// 定义 HeroSlice
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		// 1、生成 Hero 结构体实例
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 2、将生成的 hero 实例放入 heros切片实例
		heros = append(heros, hero)
	}
	// 排序前
	for _, v := range heros {
		fmt.Println(v)
	}
	fmt.Println("---------Sort() 方法排序后---------")
	// HeroSlice 结构体已全部实现接口的三个方法， 调用 Sort() 排序
	sort.Sort(heros) // 切片是引用类型，Sort() 方法内部修改 heros 会作用到源数据
	for _, v := range heros {
		fmt.Println(v)
	}

	// 交换值
	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i = ", i, "j = ", j) // i =  20 j =  10
}
