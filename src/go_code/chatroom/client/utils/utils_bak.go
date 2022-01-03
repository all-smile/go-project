package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		// info := fmt.Sprintf("read pkg header err from %v", conn.RemoteAddr().String())
		// err = errors.New(info)
		// fmt.Printf("服务端读取客户端： %v 信息失败= %v", conn.RemoteAddr().String(), err)
		return
	}
	fmt.Println("读取的buf: ", buf)
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[0:4])
	fmt.Printf("pkgLen 类型： %T, 值： %v \n", pkgLen, pkgLen)
	// 从套接字里面读取 pkgLen 字节， 放到 buf 里面去
	n1, err := conn.Read(buf[:pkgLen])
	if uint32(n1) != pkgLen || err != nil {
		err = errors.New("read pkg body err")
		return
	}
	// 反序列化 func Unmarshal(data []byte, v interface{}) error
	// Unmarshal函数解析json编码的数据并将结果存入v指向的值。
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// 1. 先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte // 数组
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 3.1 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}
	// 发送 data 本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送 data 数据失败", err)
		return
	}
	return
}
