package main

import "testing"

// 编写测试用例

func TestGetSub(t *testing.T) {
	// 调用
	res := getSub(10, 3)
	if res != 7 {
		// fmt.Printf("add(5)执行错误, 返回值：%v, 期望值： %v \n", res, 10)
		t.Fatalf("getSub(10, 3)执行错误, 返回值：%v, 期望值： %v \n", res, 7)
	} else {
		t.Logf("getSub(10, 3)执行正确")
	}
}
