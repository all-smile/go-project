// flag 包命令行参数解析
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义变量，用于接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int
	// &user 就是接收命令行里 -u 后面的值
	// u 就是指 -u 这个参数
	// "" 默认值
	// "描述：用户名，默认为空" 说明
	flag.StringVar(&user, "u", "", "描述：用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "描述：密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "描述：主机名，默认为空")
	flag.IntVar(&port, "port", 3306, "描述：端口，默认为空")

	// 转换， 必须调用
	flag.Parse()

	// 输出结果
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v \n",
		user, pwd, host, port)
}
