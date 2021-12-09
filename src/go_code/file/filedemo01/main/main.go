package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		// 创建带缓冲的 *reader
		/* const (
			 // 默认缓冲区 4096 并不是一次读取完成，读一部分处理一部分，这样可以处理一些比较大的文件
			defaultBufSize = 4096
		) */
		reader := bufio.NewReader(file)
		// 循环读取内容
		for {
			// 读到一个换行就结束
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			// 输出内容
			fmt.Print(str)
		}
		fmt.Println("文件读取结束")
	}
	fmt.Printf("file: %v \n", file) // &{0xc00008e780} 文件指针/文件对象/文件句柄
	/*
		错误文件演示
		err =  open d:/test01.txt: The system cannot find the file specified.
		file: <nil>
		关闭文件错误
	*/

	// 案例
	// 当函数退出时，要及时关闭file，否则会有内存泄露
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		}
	}()
}
