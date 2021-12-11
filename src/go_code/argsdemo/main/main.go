// 获取命令行参数 os.Args
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("命令行的参数有 %v 个 \n", len(os.Args))
	for i, v := range os.Args {
		i++
		fmt.Printf("第 %v 个 参数： %v \n", i, v)
	}
	/*
		命令行的参数有 5 个
		第 1 个 参数： D:\go-project\src\go_code\argsdemo\main\testArgs.exe
		第 2 个 参数： tom
		第 3 个 参数： aaa
		第 4 个 参数： 1111
		第 5 个 参数： 哈哈
	*/
}
