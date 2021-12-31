package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis" // 引入 redis 包
)

func main() {
	// 通过 go 向 redis 读写数据
	// 1. 连接到 redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	fmt.Println("redis 连接成功")
	defer conn.Close()

	// 2. 通过go向redis写入 string 数据
	_, err = conn.Do("Set", "name", "tom猫猫")
	if err != nil {
		fmt.Println("set 出错了")
	}
	// 3. 取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get 出错了")
		return
	}
	// 返回的 r 是interface{}, 需要类型断言
	fmt.Printf("v 类型： %T, 值： %v \n", r, r)

}
