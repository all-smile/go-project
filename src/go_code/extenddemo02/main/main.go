package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
	Age  int
}

type C struct {
	A
	B
	// Name string
}

type D struct {
	a A
}

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name string
	addr string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int
}

func main() {
	var c C
	// 如果C没有Name字段，而A和B有Name, 这时就必须通过指定匿名结构体名字来区分
	// c.Name = "tom" // ambiguous selector c.Name 模糊的选择器
	c.A.Name = "tom"
	fmt.Println("c = ", c)

	//嵌入 有名结构体 => 组合
	// D 中存在有名结构体，在访问有名结构体的字段或者方法的时候，必须带上有名结构体的名字
	var d D
	d.a.Name = "xixi"
	fmt.Println("d = ", d)

	// 字面量形式创建组合实例
	d2 := D{
		a: A{
			Name: "jerry",
			Age:  2,
		},
	}
	fmt.Println("d2 = ", d2)

	// 创建 TV 实例
	tv := TV{Goods{"小米电视", 200}, Brand{"xiaomi", "武汉"}}
	fmt.Println("tv = ", tv)
	tv02 := TV{
		Goods{
			Name:  "华为电视",
			Price: 220,
		},
		Brand{
			Name: "huawei",
			addr: "深圳",
		},
	}
	fmt.Println("tv02 = ", tv02)

	// 内层结构是地址形式时，创建实例
	tv03 := TV2{&Goods{"熊猫电视", 180}, &Brand{"熊猫", "中国"}}
	// fmt.Println("tv03 = ", tv03) // {0xc0000040a8 0xc00004e3e0}
	fmt.Println("tv03 = ", tv03.Goods, tv03.Brand)   // &{熊猫电视 180} &{熊猫 中国}
	fmt.Println("tv03 = ", *tv03.Goods, *tv03.Brand) // {熊猫电视 180} {熊猫 中国}

	// int 匿名字段是基本数据类型
	var e E
	e.Name = "牛牛"
	e.Age = 3
	e.int = 20
	fmt.Println("e = ", e) // {{牛牛 3} 20}
}
