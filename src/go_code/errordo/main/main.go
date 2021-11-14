package main

import (
	"errors"
	"fmt"
)

func test() {
	// defer + recover 捕获处理异常
	defer func() {
		/* err := recover() // recover内置函数，可以捕获异常
		if err != nil {
			// 有异常
			fmt.Println("err = ", err)
			// 做一些别的操作， 错误监控，通知工作人员等等
		} */
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	n1 := 100
	n2 := 0
	res := n1 / n2
	fmt.Println("res = ", res)
}

// 自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil
	} else {
		// 返回错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.init")
	if err != nil {
		// 如果读取错误，就输出错误信息，并退出程序
		panic(err)
	}
	fmt.Println("test02()后面的代码")
}

func main() {
	test()
	fmt.Println("接下来...")

	// 自定义错误测试
	test02()
	fmt.Println("main-test02...")
}
