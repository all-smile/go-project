package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件，并往里面写入5个 hello, Golang
	// 1、打开 d:/abc.txt
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	// 2、写入5个 hello, Golang ， 使用带缓存的 *Write
	str := "hello, Golang\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为 writer 是带缓存的，因此在调用 WriteString() 时，其实内容是先写入到缓存了，
	// 因此需要再 Flush 一下，将缓存的数据真正写到文件中，否则文件中会没有数据
	writer.Flush()

	// 操作完后，及时关闭文件，防止内存泄漏
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		}
	}()

}
