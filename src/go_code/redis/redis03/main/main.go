// redis 连接池
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的 pool
var pool *redis.Pool

// 当启动程序时，初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	fmt.Printf("pool 类型： %T, 值： %v \n", pool, pool)
	// 1. 先从 pool 里面取出连接
	conn := pool.Get()
	// conn 类型： *redis.activeConn, 值： &{0xc000070500 0xc0000aa000 0}
	fmt.Printf("conn 类型： %T, 值： %v \n", conn, conn)
	defer conn.Close()

	// 2. 操作
	_, err := conn.Do("Set", "watch", "xiaomi")
	if err != nil {
		fmt.Println("操作失败", err)
		return
	}
	val, err := redis.String(conn.Do("Get", "watch"))
	if err != nil {
		fmt.Println("操作失败", err)
		return
	}
	fmt.Println("val =", val)

	// 关闭连接池, 不能再继续读写了
	pool.Close()
	conn2 := pool.Get()
	fmt.Println("conn2=", conn2)

	_, err = conn2.Do("Set", "watch2", "xiaomi222")
	if err != nil {
		// Set()操作失败 redigo: get on closed pool
		fmt.Println("Set()操作失败", err)
		return
	}
	val, err = redis.String(conn2.Do("Get", "watch2"))
	if err != nil {
		fmt.Println("get操作失败", err)
		return
	}
	fmt.Println("val2 =", val)
}
