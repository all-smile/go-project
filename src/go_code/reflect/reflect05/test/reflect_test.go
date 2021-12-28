// reflect_test.go 测试用例
// 通过反射创建并操作结构体
package test

import (
	"reflect"
	"testing"
)

type User struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *User
		st    reflect.Type
		elem  reflect.Value
	)

	// 获取 *User 类型
	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf = ", st.Kind().String()) // ptr
	// st 真正指向的类型
	st = st.Elem()
	t.Log("reflect.TypeOf.Elem = ", st.Kind().String()) // struct
	// New 返回一个 Value 类型的值，该值持有一个指向类型为 typ 的新申请的零值的指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String())             // ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) // struct
	// model 就是创建的 User 结构体变量（实例）
	// 类型断言
	model = elem.Interface().(*User) // model 是 *User, 它的指向跟elem 是一样的
	elem = elem.Elem()               // 取得 elem 指向的值
	// 结构体赋值
	elem.FieldByName("UserId").SetString("123456")
	elem.FieldByName("Name").SetString("xiao")
	t.Log("model model.Name", model, model.Name)
}

/*
go test -v

=== RUN   TestReflectStructPtr
    reflect_test.go:24: reflect.TypeOf =  ptr
    reflect_test.go:27: reflect.TypeOf.Elem =  struct
    reflect_test.go:30: reflect.New ptr
    reflect_test.go:31: reflect.New.Elem struct
    reflect_test.go:39: model model.Name &{123456 xiao} xiao
--- PASS: TestReflectStructPtr (0.00s)
PASS
ok      go_code/reflect/reflect05/test  0.740s
*/
