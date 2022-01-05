// 稀疏数组 - 保留棋盘
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ValNode struct {
	row int
	col int
	val int // int
}

func main() {
	// 1. 原始数组, 二维数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 1 - 黑子
	chessMap[2][3] = 2 // 1 - 蓝子
	fmt.Println("原始数组 chessMap： ", chessMap)

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
	var sparseArr []ValNode

	// 标准的稀疏数组，含有记录二维数组的规模
	var oriNode ValNode = ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, oriNode)

	for row, v := range chessMap {
		for col, v2 := range v {
			if v2 != 0 {
				var node ValNode = ValNode{
					row: row,
					col: col,
					val: v2,
				}
				sparseArr = append(sparseArr, node)
			}
		}
	}

	fmt.Println("sparseArr = ", sparseArr)
	// 输出稀疏数组
	fmt.Println("存盘数据：")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d\t%d\t%d\n", i, valNode.row, valNode.col, valNode.val)
	}

	// 存盘退出 续上盘
	// 存盘 ： 稀松数组保存到文件 chessmap.data
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

	// 序列化数据，存入文件，
	data, err := json.Marshal(sparseArr)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	fmt.Println("sparseArr 序列化后： ", string(data))

	writer.WriteString(string(data))
	/* for _, valNode := range sparseArr {
		str := fmt.Sprintf("%d\t%d\t%d\n", valNode.row, valNode.col, valNode.val)
		writer.WriteString(string(str))
	} */
	// 缓存落盘
	writer.Flush()

	// 续上盘 ：读取 chessmap.data 文件，遍历每一行， 恢复数组
	_file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	// 读取内容
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
	/* // 反序列化
	var _sparseArrObj []map[string]int
	err = json.Unmarshal([]byte(oldData), &_sparseArrObj)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	fmt.Println("反序列化结果： ", _sparseArrObj)

	// 还原原始数组, 这里写死数组元素的个数， 应该使用切片，
	var newChessMap [11][11]int
	for i, _map := range _sparseArrObj {
		if i != 0 {
			newChessMap[_map["row"]][_map["col"]] = _map["val"]
		}
	}
	// 输出原始数组
	fmt.Println("======================")
	for _, v := range newChessMap {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}
	fmt.Println("======================") */

	var chessMap2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	// 输出原始的数组
	fmt.Println("恢复后的原始数组")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}
}
