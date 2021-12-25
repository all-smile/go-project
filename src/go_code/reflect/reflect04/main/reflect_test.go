// 使用反射机制，编写适配器函数 桥连接
package test

import (
	"reflect"
	"testing"
)

// type User struct {
// 	UserId string
// 	Name   string
// }

func TestReflectFunc(t *testing.T) {
	// 定义两个不同参数的函数
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		// len(in) == 3, v.Call(in)代表调用v(in[0], in[1], in[2])
		function.Call(inValue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

/*
go test -v

=== RUN   TestReflectFunc
    reflect_test.go:17: 1 2
    reflect_test.go:20: 1 2 test2
--- PASS: TestReflectFunc (0.00s)
PASS
ok      go_code/reflect/reflect04/main  0.545s
*/
