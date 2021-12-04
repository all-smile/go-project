package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (monkey *Monkey) climping() {
	fmt.Printf("猴子：%v, 天生会爬树 \n", monkey.Name)
}

// 小猴子继承了老猴子会爬树的技能
// 小猴子天赋异禀学会了老猴子不会的飞行技能, 以及游泳技能
// 不能破坏老猴子的属性方法，所以需要接口来对继承做补充
type LittleMonkey struct {
	Monkey
}

func (littleMonkey LittleMonkey) Flying() {
	fmt.Printf("名字：%v, 通过学习，会飞了！！！ \n", littleMonkey.Name)
}
func (littleMonkey LittleMonkey) Swimming() {
	fmt.Printf("名字：%v, 通过学习，会游泳了！！！ \n", littleMonkey.Name)
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climping()
	// 实现了接口的行为，但是并没有采用接口的调用方式
	// 如下这样仅仅是调用结构体绑定的方法
	/* monkey.Flying()
	monkey.Swimming() */

	// 调用接口
	fly(monkey)
	swim(monkey)
}
func fly(bird BirdAble) {
	bird.Flying()
}

func swim(fish FishAble) {
	fish.Swimming()
}
