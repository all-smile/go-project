// 稀疏数组 - 保留棋盘
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. 原始数组, 二维数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 1 - 黑子
	chessMap[2][3] = 2 // 1 - 蓝子
	// chessMap[5][5] = 2 // 1 - 蓝子

	// 2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}

	// 3. 转成新数组 - 稀疏数组
	// 遍历数组，把值不为0的，创建一个 ValNode 结构体实例， 然后放入对应的切片

	// 定义切片
	var sparseArrObj []map[string]int

	// 标准的稀疏数组，含有记录二维数组的规模
	var obj map[string]int = make(map[string]int)
	obj["row"] = 11
	obj["col"] = 11
	obj["val"] = 0
	sparseArrObj = append(sparseArrObj, obj)

	for row, v := range chessMap {
		for col, v2 := range v {
			if v2 != 0 {
				var obj map[string]int = make(map[string]int)
				obj["row"] = row
				obj["col"] = col
				obj["val"] = v2
				sparseArrObj = append(sparseArrObj, obj)
			}
		}
	}
	fmt.Println("sparseArrObj", sparseArrObj)

	// 4. 输出稀疏数组
	fmt.Println("map切片：")
	for i, _map := range sparseArrObj {
		fmt.Printf("%d: %d\t%d\t%d\n", i, _map["row"], _map["col"], _map["val"])
	}

	// 存盘退出 续上盘
	// 5. 存盘 ： 稀松数组保存到文件 chessmap.data
	fmt.Println("存盘数据, 'd:/chessmap.data'")
	filepath := "d:/chessmap.data"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		}
	}()
	// 写入数据, 带缓冲， writer.Flush()
	writer := bufio.NewWriter(file)

	// 5.1 序列化数据，存入文件，
	data, err := json.Marshal(sparseArrObj)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	fmt.Println("sparseArrObj 序列化后： ", string(data))

	writer.WriteString(string(data))
	// 缓存落盘
	writer.Flush()

	// 6. 续上盘 ：读取 chessmap.data 文件，遍历每一行， 恢复数组
	_file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	// 6.1 读取内容
	fmt.Println("文件读取输出")
	var oldData string
	reader := bufio.NewReader(_file)
	for {
		// 读到一个换行就结束
		oldData, err = reader.ReadString('\n')
		// 输出内容
		fmt.Println("oldData=", oldData)
		if err == io.EOF {
			// 读到一行末尾
			break
		}
	}
	// 6.2 反序列化
	var _sparseArrObj []map[string]int
	err = json.Unmarshal([]byte(oldData), &_sparseArrObj)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	fmt.Println("反序列化结果： ", _sparseArrObj)

	// 7. 还原原始数组, 这里写死数组元素的个数， 应该使用切片，
	var newChessMap [11][11]int
	for i, _map := range _sparseArrObj {
		if i != 0 {
			newChessMap[_map["row"]][_map["col"]] = _map["val"]
		}
	}
	// 7.1 输出原始数组
	fmt.Println("======================")
	for _, v := range newChessMap {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}
	fmt.Println("======================")

}
