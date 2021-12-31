package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
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

	// hash 单个读取
	_, err = conn.Do("Hset", "cat", "name", "tom")
	if err != nil {
		fmt.Println("Hset 出错了")
	}
	_, err = conn.Do("Hset", "cat", "age", "12")
	if err != nil {
		fmt.Println("Hset 出错了")
	}
	// 3. 取数据
	h, err := redis.String(conn.Do("Hget", "cat", "name"))
	if err != nil {
		fmt.Println("get 出错了")
		return
	}
	fmt.Printf("hash 类型： %T, 值： %v \n", h, h)
	h2, err := redis.String(conn.Do("Hget", "cat", "age"))
	if err != nil {
		fmt.Println("get 出错了")
		return
	}
	fmt.Printf("hash 类型： %T, 值： %v \n", h2, h2)
	// hash 一次读写多个
	_, err = conn.Do("HMSet", "phone1", "name", "apple", "action", "call")
	if err != nil {
		fmt.Println("HMSet 出错了", err)
		return
	}
	// redis.Strings 处理多个值
	h3, err := redis.Strings(conn.Do("HMGet", "phone1", "name", "action"))
	if err != nil {
		fmt.Println("HMGet 出错了")
		return
	}
	fmt.Printf("hash 类型： %T, 值： %v \n", h3, h3)
	for _, v := range h3 {
		fmt.Println(v)
	}
}
