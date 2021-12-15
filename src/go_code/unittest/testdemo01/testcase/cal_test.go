package main

import (
	"testing" // 引入 testing 框架
)

// 编写测试用例

func TestAdd(t *testing.T) {
	// 调用
	res := add(5)
	if res != 10 {
		// fmt.Printf("add(5)执行错误, 返回值：%v, 期望值： %v \n", res, 10)
		t.Fatalf("add(5)执行错误, 返回值：%v, 期望值： %v \n", res, 10)
	} else {
		t.Logf("add(5)执行正确")
	}
}

func TestHello(t *testing.T) {
	// fmt.Printf("t类型： %T, 值： %v \n", t, t)
	t.Logf("TestHello被调用")
}

/*
go test -v

=== RUN   TestAdd
    cal_test.go:16: add(5)执行正确
--- PASS: TestAdd (0.00s)
=== RUN   TestHello
    cal_test.go:22: TestHello被调用
--- PASS: TestHello (0.00s)
PASS
ok      go_code/unittest/testdemo01/testcase    6.356s
*/
