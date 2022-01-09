// 环形单向链表
package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func insertNode(head *CatNode, newCatNode *CatNode) {
	// 1. 处理第一个节点(头节点有值)
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 构成环形
		fmt.Printf("第一个节点： 猫： %v 加入环形链表 \n", newCatNode.name)
		return
	}

	// 2. 处理别的节点
	// 找到最后的节点
	// 定义临时变量
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表
	temp.next = newCatNode
	newCatNode.next = head
}

// 按顺序插入
func insertNode02(head *CatNode, newCatNode *CatNode) {
	fmt.Println("insertNode02 - head: ", head)
	// 1. 处理第一个节点(头节点有值)
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 构成环形
		fmt.Printf("第一个节点： 猫： %v 加入环形链表 \n", newCatNode.name)
		return
	}

	// 2. 处理别的节点
	// 定义临时变量
	temp := head
	flag := true
	for {
		if temp.no == newCatNode.no {
			fmt.Printf("节点： %v ,  已经存在，不允许加入", newCatNode.no)
			flag = false
			break
		} else if temp.next == head {
			break
		} else if temp.next.no > newCatNode.no {
			break
		} else if temp.next.no == newCatNode.no {
			fmt.Printf("节点： %v ,  已经存在，不允许加入\n", newCatNode.no)
			flag = false
			break
		}
		temp = temp.next
	}
	// 加入到链表
	// temp.next = newCatNode
	// newCatNode.next = head

	if !flag {
		fmt.Println("不能插入节点")
		return
	} else {

		// 最后一个节点后插入
		if temp.next == head {
			temp.next = newCatNode
			newCatNode.next = head
		} else {
			// 中间节点插入
			newCatNode.next = temp.next
			temp.next = newCatNode
		}

	}
}

// 删除节点
func delNode(head *CatNode, id int) *CatNode {
	// 这里的头节点是有值的
	fmt.Println("delNode - head: ", head)
	temp := head
	helper := head // helper 指向 head 的上一个节点， 即尾节点

	// 1. 空链表
	if temp.next == nil {
		fmt.Println("空链表")
		return head
	}

	// 2. 只有一个
	if temp.next == head {
		temp.next = nil // 没有任何指向就会被gc回收，即删除
		return head
	}

	// 3. 将 helper 指向最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 4. 两个及以上节点
	flag := true
	for {
		if temp.next == head {
			// 只是比较，还没有执行删除操作
			break
		}
		if temp.no == id {
			if temp == head {
				// 删除的是头节点, 修改头节点的指向
				head = temp.next
			}
			helper.next = temp.next
			flag = false
			fmt.Printf("猫: no: %v 被移除链表了 \n", id)
			break
		}
		temp = temp.next     // 移动 - 比较
		helper = helper.next // 移动 - 执行删除
	}
	// 这里要增加一次比较
	if flag {
		if temp.no == id {
			// temp 就是 要删除的节点
			helper.next = temp.next
		} else {
			fmt.Println("没有找到要删除的节点id: ", id)
			return head
		}
	}
	return head
}

// 显示所有节点
func showNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[猫: no: %v, name: %v] => ", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	// 初始化环形链表头节点
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tom1",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	insertNode02(head, cat1)
	insertNode02(head, cat3)
	insertNode02(head, cat2)
	insertNode02(head, cat3)

	fmt.Println("显示")
	showNode(head)

	fmt.Println("删除 1")
	// 删除之后，重新制定 头节点 head
	head = delNode(head, 1)

	showNode(head)
}

/*
insertNode02 - head:  &{0  <nil>}
第一个节点： 猫： tom1 加入环形链表
insertNode02 - head:  &{1 tom1 0xc00004e3c0}
insertNode02 - head:  &{1 tom1 0xc00004e420}
insertNode02 - head:  &{1 tom1 0xc00004e400}
节点： 3 ,  已经存在，不允许加入不能插入节点
显示
[猫: no: 1, name: tom1] => [猫: no: 2, name: tom2] => [猫: no: 3, name: tom3] =>
删除 1
delNode - head:  &{1 tom1 0xc00004e400}
猫: no: 1 被移除链表了
[猫: no: 2, name: tom2] => [猫: no: 3, name: tom3] =>
*/
