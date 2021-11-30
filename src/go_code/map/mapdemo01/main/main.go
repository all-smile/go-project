package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	/*
		如果name存在，则修改密码为 8888888
		如果不存在，则存储 nickname 和 password
	*/
	// val, ok = users[name]
	if users[name] != nil {
		users[name]["password"] = "8888888"
	} else {
		// 先make, 分配内存
		users[name] = make(map[string]string)
		// 存储信息
		users[name]["nickname"] = name
		users[name]["password"] = "8888888"
	}
}

func main() {
	var a map[string]string
	fmt.Printf("a类型： %T, 值： %v \n", a, a) // a类型： map[string]string, 值： map[]
	// make 作用就是分配空间
	a = make(map[string]string, 2) // 可以存放2个key-value
	a["test"] = "xiao"
	a["test"] = "xiao~"
	a["a"] = "xiao~"
	a["b"] = "xiao~"
	// key重复的话，value值就会覆盖调原先的，key是无序的
	fmt.Printf("a类型： %T, 值： %v \n", a, a) // a类型： map[string]string, 值： map[test:xiao~]

	var cities map[string]string = map[string]string{
		"beijing": "北京", // 必须要写 , 逗号
	}
	fmt.Printf("cities类型： %T, 值： %v \n", cities, cities)
	// cities类型： map[string]string, 值： map[beijing:北京]

	var cities01 = make(map[string]string)
	cities01["no1"] = "北京"
	fmt.Printf("cities01 类型： %T, 值： %v \n", cities01, cities01)
	// cities01 类型： map[string]string, 值： map[]

	// 练习
	// 存放学生信息，学生有name和sex字段
	studentMap := make(map[string]map[string]string)
	fmt.Printf("studentMap 类型： %T, 值： %v \n", studentMap, studentMap)
	studentMap["01"] = make(map[string]string)
	studentMap["01"]["name"] = "xiao"
	studentMap["01"]["sex"] = "男"

	studentMap["02"] = make(map[string]string)
	studentMap["02"]["name"] = "xiao"
	studentMap["02"]["sex"] = "男"
	studentMap["02"]["addr"] = "上海"
	fmt.Println("studentMap = ", studentMap)
	fmt.Printf("studentMap[01][addr] = %q \n", studentMap["01"]["addr"]) // studentMap[01][addr] = ""

	// 删除 studentMap["02"]["addr"]
	delete(studentMap["02"], "addr")
	fmt.Println("studentMap[02] = ", studentMap["02"])

	var bookMap map[string]string = map[string]string{
		"01": "Golang",
		"02": "GoWeb",
	}
	fmt.Println("bookMap = ", bookMap)
	// 重新赋值，删除
	// bookMap = make(map[string]string)
	// fmt.Println("bookMap = ", bookMap)

	// 查找
	val, ok := bookMap["03"]
	// 返回两个值，一个是key对应的value, 另一个是否有这个key
	if ok {
		fmt.Printf("存在bookMap[01]， 值为： %v \n", val)
	} else {
		fmt.Println("不存在bookMap[01]")
	}

	// for range 遍历map
	var bookMap01 map[string]string = map[string]string{
		"01": "Golang",
		"02": "GoWeb",
	}
	for k, v := range bookMap01 {
		fmt.Printf("k = %v, v= %v \n", k, v)
	}

	for k1, v1 := range studentMap {
		fmt.Println("k1 = ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t%v[%v] = %v \n", k1, k2, v2)
		}
		fmt.Println()
	}

	// 长度， key: value的个数
	fmt.Println("bookMap01的长度 = ", len(bookMap01))

	// map 结构体
	/*
		key是学号， 唯一的
		value是结构体
	*/
	// 自定义结构体类型 Stu
	type Stu struct {
		Name string
		Age  int
		Addr string
	}
	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 10, "上海"}
	fmt.Printf("stu1类型：%T, 值： %v \n", stu1, stu1) // stu1类型：main.Stu, 值： {tom 10 上海}
	stu2 := Stu{"jerry", 9, "上海01"}
	students["001"] = stu1
	students["002"] = stu2
	fmt.Println("students = ", students) // students =  map[001:{tom 10 上海} 002:{jerry 9 上海01}]

	// 遍历学生信息
	for k, v := range students {
		fmt.Printf("编号： %v\t", k)
		fmt.Printf("姓名： %v\t", v.Name)
		fmt.Printf("年龄： %v\t", v.Age)
		fmt.Printf("地址： %v\t", v.Addr)
		fmt.Println()
	}

	// 练习
	users := make(map[string]map[string]string, 10)
	fmt.Println("users = ", users)
	modifyUser(users, "xiao")
	fmt.Println("修改后users = ", users) // map[xiao:map[nickname:xiao password:8888888]]
}
