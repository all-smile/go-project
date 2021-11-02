package main

// import "unsafe"
// import "fmt"
import (
	"fmt"
	"strconv"
	_ "unsafe"
)

func main() {
	var str string = "true"
	var b bool
	// b, _ = strconv.ParseBool(str)
	// func strconv.ParseBool(str string) (bool, error)
	// 1、返回两个值，一个是转换的bool值，一个是error
	// 2、我们只需要拿到第一个返回值， 第二个忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("%T %v\n", b, b)

	var str2 string = "123"
	var n int64
	var n2 int
	n, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n)
	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("%T %v\n", f1, f1)

	var str4 string = "hello"
	var f2 float64
	var b2 bool = true
	f2, _ = strconv.ParseFloat(str4, 64)
	fmt.Printf("%T %v\n", f2, f2)
	b2, _ = strconv.ParseBool(str4)
	fmt.Printf("%T %v\n", b2, b2)
}
