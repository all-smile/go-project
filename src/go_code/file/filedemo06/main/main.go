// 文件导入
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	fmt.Println("fileInfo = ", fileInfo)
	if err == nil {
		// 文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		// 不存在
		return false, nil
	}
	return false, err
}

func main() {
	// 将 d:/abc.txt 文件内容 导入到 d:/test.txt

	filepath01 := "d:/abc.txt"
	filepath02 := "d:/test.txt"
	// 1、将 d:/abc.txt 内容 读取到内存中
	fileData, err := ioutil.ReadFile(filepath01)
	if err != nil {
		fmt.Println("ReadFile -err", err)
		return
	}
	// 2、将内存中的内容写入到 d:/test.txt
	err = ioutil.WriteFile(filepath02, fileData, 0666)
	if err != nil {
		fmt.Println("写文件出错了")
	}
	// 不需要显示的 Open 和 Close

	// 判断文件或者目录是否存在
	ok, err := PathExists(filepath01)
	if err != nil {
		fmt.Printf("文件或者目录： %v 不存在", filepath01)
		return
	}
	if ok {
		fmt.Printf("文件或者目录： %v 存在", filepath01)
	}
}
