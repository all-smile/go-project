// 迷宫

package main

import "fmt"

// 老鼠找路
func SetWay(mazeMap *[8][7]int, i int, j int) bool {
	var flag bool = false
	if mazeMap[6][5] == 2 {
		// 终点
		flag = true
	} else {
		if mazeMap[i][j] == 0 {
			// 标记可以走
			mazeMap[i][j] = 2

			// 这个点可以探测, 上下左右四个方向 实际走的路线是 "W"形状
			/* if SetWay(mazeMap, i-1, j) { // 上
				flag = true
			} else if SetWay(mazeMap, i+1, j) { // 下
				flag = true
			} else if SetWay(mazeMap, i, j-1) { // 左
				flag = true
			} else if SetWay(mazeMap, i, j+1) { // 右
				flag = true
			} else {
				mazeMap[i][j] = 3
				flag = false
			} */

			// 更换 "下右上左" 的策略
			if SetWay(mazeMap, i+1, j) { // 下
				flag = true
			} else if SetWay(mazeMap, i, j+1) { // 右
				flag = true
			} else if SetWay(mazeMap, i+1, j) { // 上
				flag = true
			} else if SetWay(mazeMap, i, j-1) { // 左
				flag = true
			} else {
				mazeMap[i][j] = 3
				flag = false
			}

		} else if mazeMap[i][j] == 1 {
			// 这个点是墙，不能探测
			flag = false
		}
	}
	return flag
}

func main() {
	// 创建二维数组 地图迷宫
	/*
	   1. 如果元素为 1 代表墙
	   2.           0 代表没有走过的点
	   3.           2 代表可以行走
	   4.           3 代表走过的路，但是没有走通
	*/

	var mazeMap [8][7]int // 8 行 7 列

	// 初始化地图上的墙
	// 上下的墙
	for i := 0; i < 7; i++ {
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
	}
	// 左右的墙
	for i := 0; i < 8; i++ {
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
	}

	// 设置墙
	mazeMap[3][1] = 1
	mazeMap[3][2] = 1
	// mazeMap[1][2] = 1
	// mazeMap[2][2] = 1

	// 输出查看
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j], " ")
		}
		fmt.Println()
	}
	/*
		1 1 1 1 1 1 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 1 1 1 1 1 1
	*/

	// 测试
	fmt.Println("探测")
	SetWay(&mazeMap, 1, 1)

	// 输出查看
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j], " ")
		}
		fmt.Println()
	}
}

/*
// 没有设置任何墙， 上下左右的策略
1 1 1 1 1 1 1
1 2 2 2 2 2 1
1 2 2 2 2 2 1
1 2 2 2 2 2 1
1 2 2 2 2 2 1
1 2 2 2 2 2 1
1 1 1 1 1 1 1

// 设置两个墙， 上下左右的策略
1 1 1 1 1 1 1
1 2 2 2 2 2 1
1 2 2 2 2 2 1
1 1 1 2 2 2 1
1 3 3 2 2 2 1
1 3 3 2 2 2 1
1 3 3 2 2 2 1
1 1 1 1 1 1 1

// 设置两个墙，下右上左的策略
1 1 1 1 1 1 1
1 2 0 0 0 0 1
1 2 2 2 0 0 1
1 1 1 2 0 0 1
1 0 0 2 0 0 1
1 0 0 2 0 0 1
1 0 0 2 2 2 1
1 1 1 1 1 1 1

// 设置墙堵死
1 1 1 1 1 1 1
1 3 1 0 0 0 1
1 3 1 0 0 0 1
1 1 1 0 0 0 1
1 0 0 0 0 0 1
1 0 0 0 0 0 1
1 0 0 0 0 0 1
1 1 1 1 1 1 1
*/
