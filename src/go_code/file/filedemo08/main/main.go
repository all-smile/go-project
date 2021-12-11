// 统计文件中英文、数字、空格和其它字符的数量
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，保存统计的字符以及数量
type CharCount struct {
	EnCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	// 思路：
	// 打开一个文件，创建一个 Reader
	// 每读取一行，就去统计不同的字符数量，将结构保存到结构体中

	// 定义 CharCount 实例
	charCount := CharCount{}

	filepath := "d:/test.txt"
	file, err := os.Open(filepath)
	var reader *bufio.Reader

	if err != nil {
		fmt.Printf("读取文件：%v 失败", filepath)
	} else {
		// 创建带缓存的 *reader
		reader = bufio.NewReader(file)
		// 循环读取内容
		for {
			// 读到一个换行就结束
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				// 读到文件末尾就退出
				break
			}
			// str2 := []rune(str) // 处理中文字符
			for _, v := range str {
				// fmt.Printf("%c", v) // 查看字符
				switch {
				case v >= 'a' && v <= 'z':
					fallthrough // 穿透
				case v >= 'A' && v <= 'Z':
					charCount.EnCount++
				case v == ' ' || v == '\t':
					charCount.SpaceCount++
				case v >= '0' && v <= '9':
					charCount.NumCount++
				default:
					charCount.OtherCount++
				}
			}
			// 输出内容
			fmt.Print(str)
		}
		fmt.Println("文件读取结束")

		// 输出统计的结果
		fmt.Printf("英文个数：%v, 数字个数：%v, 空格个数：%v, 其它字符个数：%v \n",
			charCount.EnCount, charCount.NumCount, charCount.SpaceCount, charCount.OtherCount)
	}

	// 关闭文件
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		} else {
			fmt.Printf("文件： %v 已关闭", filepath)
		}
	}()
}
