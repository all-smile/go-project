package main

import (
	"fmt"
	_ "strconv"
	"strings"
)

func main() {
	// 遍历字符串
	/* name := "xiao"
	name01 := "xiao肖"
	fmt.Println("name len=", len(name))     // name len= 4
	fmt.Println("name01 len=", len(name01)) // name len= 7 // 一个汉字占三个字节
	_name := []rune(name01)                 // str 转成 []rune 切片
	fmt.Printf("%T %v \n", _name, _name)
	for i := 0; i < len(_name); i++ {
		fmt.Printf("_name[%v] = %c \n", i, _name[i])
	} */

	// 字符串转数字
	/* n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转化成功")
		fmt.Println("n = ", n, " err = ", err)
	} */

	// 整数转字符串
	/* str := strconv.Itoa(123)
	fmt.Printf("str 类型：%T, 值：%s", str, str) */

	// 字符串转 []byte
	/* var bytes = []byte("hello go")
	fmt.Printf("bytes类型： %T, 值： %v \n", bytes, bytes) */

	// []byte转字符串
	/* var str = string([]byte{97, 98, 99})
	fmt.Printf("str类型： %T, 值： %s", str, str) */

	// 十进制转其它进制
	/* var n1 = strconv.FormatInt(8, 2)
	fmt.Printf("n1类型： %T, 值： %v", n1, n1) */

	// 查找字符串是否包含某个子串
	var isHasStr = strings.Contains("hello go", "go")
	fmt.Printf("isHasStr类型： %T, 值： %v \n", isHasStr, isHasStr)

	// 统计一个字符串中有几个指定的子串
	var n2 = strings.Count("xiaoxiaoxiao", "xiao")
	fmt.Println("n2 = ", n2)

	// 不区分大小写的字符串比较（===区分大小写）
	var b2 = strings.EqualFold("abc", "ABC")
	fmt.Println("b2 = ", b2)
	var b3 = ("abc" == "ABC")
	fmt.Println("b3 = ", b3)

	// 返回子串在字符串第一次出现的index的值
	var n3 = strings.Index("hello,你好", "好")
	fmt.Println("n3 = ", n3)

	// 返回子串在字符串最后一次出现的index的值
	var n4 = strings.LastIndex("hello,你好好", "好")
	fmt.Println("n4 = ", n4)

	// 替换字符串
	str02 := "go go hello"
	str03 := strings.Replace(str02, "go", "北京", 2)
	fmt.Println("str03= ", str03)
	fmt.Println("str02= ", str02)

	// 按照指定字符进行字符串分割
	var strArr = strings.Split("go,go,hello", ",")
	fmt.Printf("strArr类型： %T, 值： %v \n", strArr, strArr)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("第%v个，%v \n", i, strArr[i])
	}

}
