package main

// import "unsafe"
// import "fmt"
import (
"fmt"
"unsafe"
)

func main() {
	var i int8 = 12
	fmt.Println("i= ", i)

	var i2 uint8 = 129
	fmt.Println("i2= ", i2)

	var n = 12
	fmt.Printf("n的类型是%T\n", n) // n的类型是int

	var name = "xiao"
	fmt.Printf("name的类型是%T\n", name) // name的类型是string

	// 查看变量占用的字节大小和数据类型
	fmt.Printf("name的类型是%T", name, unsafe.Sizeof(name))
}

