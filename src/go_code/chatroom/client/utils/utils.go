package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

// 面向对象思想
// 封装

// 定义一个结构体， 封装方法
type Transfer struct {
	Conn net.Conn
	// 传输时，使用缓冲
	Buf [8096]byte
}

func (tf *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		return
	}
	fmt.Println("读取的buf: ", tf.Buf[:4])
	var pkgLen uint32 = binary.BigEndian.Uint32(tf.Buf[0:4])
	// 从套接字里面读取 pkgLen 字节， 放到 tf.Buf 里面去
	n1, err := tf.Conn.Read(tf.Buf[:pkgLen])
	if uint32(n1) != pkgLen || err != nil {
		err = errors.New("read pkg body err")
		return
	}
	// 反序列化 func Unmarshal(data []byte, v interface{}) error
	// Unmarshal函数解析json编码的数据并将结果存入v指向的值。
	err = json.Unmarshal(tf.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	return
}

func (tf *Transfer) WritePkg(data []byte) (err error) {
	// 1. 先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(tf.Buf[0:4], pkgLen)
	// 3.1 发送长度
	n, err := tf.Conn.Write(tf.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}
	fmt.Printf("客户端发送消息长度: %v, 内容： %s \n", len(data), string(data))
	// 发送 data 本身
	n, err = tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送 data 数据失败", err)
		return
	}
	return
}
