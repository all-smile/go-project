// 拷贝文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 编写拷贝文件或者目录的函数
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// 需要实现一个 Writer 和一个 Reader
	// 1、 先读取源文件，得到一个 Reader
	var reader *bufio.Reader
	var writer *bufio.Writer
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("Open file: %v error: %v \n", srcFileName, err)
	} else {
		reader = bufio.NewReader(srcFile)
		// reader 类型： *bufio.Reader,
		fmt.Printf("reader 类型： %T, 值： %v \n", reader, reader)
	}
	defer srcFile.Close()
	// 2、打开 dstFileName ，得到一个 Writer
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open file: %v error: %v \n", dstFileName, err)
	} else {
		writer = bufio.NewWriter(dstFile)
		// writer 类型： *bufio.Writer,
		fmt.Printf("writer 类型： %T, 值： %v \n", writer, writer)
	}
	defer dstFile.Close()
	// 3、传入 Reader Writer , 调用 os.Copy(), 并返回
	return io.Copy(writer, reader)

}

func main() {
	filepath01 := "C:/Users/xiao/Pictures/iphone/baby.jpg"
	filepath02 := "C:/Users/xiao/Pictures/iphone/01.jpg"
	// 将 filepath02 拷贝到 filepath01, 并命名为 01.jpg
	_, err := CopyFile(filepath02, filepath01)
	if err != nil {
		fmt.Println("拷贝失败", err)
	} else {
		fmt.Println("拷贝成功")
	}
}
