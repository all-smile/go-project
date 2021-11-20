package main

import "fmt"

func main() {
	// string底层是一个 byte 数组，因此string也可以进行切片
	str := "xiaobaobao"
	// 使用切片 获取baobao
	slice := str[4:]
	fmt.Printf("slice类型： %T, 值： %v \n", slice, slice)
	fmt.Printf("str[2] = %c\n", str[2])
	fmt.Printf("slice[2] = %c\n", slice[2])

	// 修改字符串 xiaobaobao -> xiaotaotao
	// 1、string -> []byte 或者 []rune
	// 2、修改
	// 3、重写转成string
	arr := []byte(str) // 转成 []byte 切片 ，可以处理英文和数字，不能处理中文
	arr[4] = 't'
	arr[7] = 't'
	str2 := string(arr)
	fmt.Printf("str2类型： %T, 值： %v \n", str2, str2)

	str1 := "小笑笑"
	arr1 := []rune(str1)
	fmt.Printf("arr1类型： %T, 值： %v \n", arr1, arr1)
	arr1[1] = '涛'
	arr1[2] = '涛'
	str3 := string(arr1)
	fmt.Printf("str3类型： %T, 值： %v \n", str3, str3)
}
