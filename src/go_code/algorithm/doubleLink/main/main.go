// 双向链表 - 增删改查
package main

import (
	"fmt"
)

// 定义节点
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 表示指向前一个节点
	next     *HeroNode // 表示指向下一个节点
}

// 添加节点
func insertNode(head *HeroNode, newNode *HeroNode) {
	// 1. 找到链表最后的节点
	// 2. 创建辅助节点， 跑龙套的
	temp := head
	for {
		if temp.next == nil {
			// 最后的节点
			break
		}
		temp = temp.next // temp 一直后移, 直到指向链表最后的节点
	}
	// 3. 修改 temp 节点的 next 指向 newNode, 并将新节点的 pre 指向 temp
	temp.next = newNode
	newNode.pre = temp
}

// 根据 编号 no 插入到指定位置（从小到大）
func insertNode02(head *HeroNode, newNode *HeroNode) {
	// 1. 找到链表适当的节点
	// 2. 创建辅助节点， 跑龙套的
	temp := head
	flag := true
	for {
		if temp.next == nil {
			// 最后的节点
			break
		} else if temp.next.no > newNode.no {
			// 升降序， 只需要修改 temp.next.no < newNode.no
			break
		} else if temp.next.no == newNode.no {
			fmt.Println("已经存在，不允许加入")
			flag = false
			break
		}
		temp = temp.next // temp 一直后移,
	}
	if !flag {
		fmt.Println("不能插入节点")
		return
	} else {
		// newNode 应该插入 temp.next 前面， temp后面
		newNode.next = temp.next
		newNode.pre = temp

		if temp.next != nil {
			temp.next.pre = newNode
		}

		temp.next = newNode

	}
}

// 删除节点
func delNode(head *HeroNode, id int) {
	// 1. 先来一个跑龙套的
	temp := head

	flag := false

	// 2. 找到要删除的节点
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			// 找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		// 删除， 更改节点指向
		temp.next = temp.next.next
		// 判断是否是最后一个
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的节点id不存在")
	}
}

// 显示链表所有节点信息
func showNode(head *HeroNode) {
	// 1. 创建辅助节点， 跑龙套的
	temp := head
	// 2. 判断该链表是不是空的
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	// 2. 遍历链表
	for {
		fmt.Printf("[%d , %s , %s] => ", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

// 逆序 显示链表所有节点信息
func showNode02(head *HeroNode) {
	// 1. 创建辅助节点， 跑龙套的
	temp := head

	// 2. 判断该链表是不是空的
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	// temp 定位到最后节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 2. 遍历链表
	for {
		fmt.Printf("[%d , %s , %s] => ", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

func main() {
	// 1. 创建头节点
	head := &HeroNode{}

	// 2. 创建 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "luffy",
		nickname: "海贼王",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "索隆",
		nickname: "大剑豪",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "山治",
		nickname: "allblue",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "甚平",
		nickname: "海侠",
	}

	insertNode02(head, hero1)
	insertNode02(head, hero3)
	insertNode02(head, hero2)
	insertNode02(head, hero4)

	fmt.Println("正序显示")
	showNode(head)

	fmt.Println("删除指定节点")
	delNode(head, 3)

	fmt.Println("逆序显示")
	showNode02(head)

}

/*
正序显示
[1 , luffy , 海贼王] => [2 , 索隆 , 大剑豪] => [3 , 山治 , allblue] => [4 , 甚平 , 海侠] =>
删除指定节点
逆序显示
[4 , 甚平 , 海侠] => [2 , 索隆 , 大剑豪] => [1 , luffy , 海贼王] =>
*/
