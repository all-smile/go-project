// 打开已存在的文件，把内容显示到终端，并追加内容 “hello, beijing”
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 1、打开 d:/abc.txt
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	// 2、读取内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		// 读到一个换行就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			// 读到一行末尾
			break
		}
		// 输出内容
		fmt.Print(str)
	}

	// 3、写入5个 hello, Golang ， 使用带缓存的 *Write
	str := "hello, beijing!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	// 4、将内容写入磁盘文件
	// 因为 writer 是带缓存的，因此在调用 WriteString() 时，其实内容是先写入到缓存了，
	// 因此需要再 Flush 一下，将缓存的数据真正写到文件中，否则文件中会没有数据
	writer.Flush()

	// 5、操作完后，及时关闭文件，防止内存泄漏
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		}
	}()

}
