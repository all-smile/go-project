// 哈希表
/*
Google的一个上机题：
有一个公司,当有新的员工来报道时,
要求将该员工的信息加入(id,性别,年龄,住址.）,
当输入该员工的 id 时,要求查到该员工的所有信息。
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

// 定义雇员
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 显示 emp 信息
func (emp *Emp) ShowEmpInfo() {
	fmt.Printf("链表： %v 内找到雇员，id=%v，name=%v \n", emp.Id%7, emp.Id, emp.Name)
}

// 定义雇员链, 不带表头
type EmpLink struct {
	Head *Emp
}

// 添加员工的方法, id从小到大
func (el *EmpLink) Insert(emp *Emp) {
	// 定义辅助指针
	cur := el.Head
	var pre *Emp = nil
	// 当前是空链表
	if cur == nil {
		el.Head = emp
		return
	}
	// 链首： 插入的元素要放到链表首
	if cur.Id > emp.Id {
		emp.Next = cur
		el.Head = emp
		return
	}
	// 中间+最后：不是空链，找到适当的位置插入, 要么 cur 前, 要么后
	for {
		if cur != nil {
			if cur.Id >= emp.Id {
				break
			}
		} else {
			// 找到最后， cur == nil
			break
		}
		pre = cur
		cur = cur.Next
	}
	emp.Next = cur
	pre.Next = emp
	/* if cur != nil {
		// 利用 pre 将 emp 插入到 cur 前面
		emp.Next = cur
		pre.Next = emp
	} else {
		emp.Next = cur
		pre.Next = emp // 这个写不写无所谓， cur就是nil
	} */
}

// 根据 id 查找 emp
func (el *EmpLink) FindById(id int) (emp *Emp) {
	cur := el.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			return nil
		}
		cur = cur.Next
	}
}

// 根据 id 删除 emp
func (el *EmpLink) DelById(id int) (err error) {
	// emp := el.FindById(id)
	// if emp == nil {
	// fmt.Printf("要删除的雇员id: %v 不存在\n", id)
	// return errors.New("id不存在")
	// }
	cur := el.Head
	var pre *Emp = nil
	fmt.Println("删除", id, cur)
	// 空链表
	if cur == nil {
		fmt.Println("空链表")
		return errors.New("空链表")
	}
	// 将删除的就是头节点
	if cur.Id == id {
		el.Head = nil
		return
	}
	// 要删除的不是头节点
	for {
		if cur != nil && cur.Id == id {
			pre.Next = cur.Next
			return
		} else if cur == nil {
			fmt.Printf("要删除的雇员id: %v 不存在\n", id)
			return errors.New("id不存在")
		}
		pre = cur
		cur = cur.Next
	}
}

// 定义 hash table, 链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 编写 哈希表 添加方法
func (ht *HashTable) Insert(emp *Emp) {
	// 1. 使用散列函数，确定添加到哪个链表里
	linkNo := ht.HashFun(emp.Id)
	// 2. 添加到对应的链表
	ht.LinkArr[linkNo].Insert(emp)
}

// 编写 哈希表 散列函数 (映射函数)
// 这里只展示按 id 散列， 没有二级链表（二级链表需要取模两次）
func (ht *HashTable) HashFun(id int) int {
	return id % 7
}

// 显示当前链表的 emp
func (el *EmpLink) ShowEmp(no int) {
	if el.Head == nil {
		fmt.Printf("链表： %v 为空\n", no)
		return
	}
	cur := el.Head
	fmt.Println("链表： ", no)
	for {
		if cur != nil {
			fmt.Printf("emp[id]=%v, emp[name]=%v => ", cur.Id, cur.Name)
		} else {
			break
		}
		cur = cur.Next
	}
	fmt.Println()
}

// 显示所有 emp
func (ht *HashTable) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowEmp(i)
	}
}

// 编写查找 emp 方法
func (ht *HashTable) FindEmpById(id int) *Emp {
	// 1. 使用散列函数，确定在哪个链表里
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].FindById(id)
}

// 编写删除 emp 方法
func (ht *HashTable) DelEmpById(id int) error {
	// 1. 使用散列函数，确定在哪个链表里
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].DelById(id)
}

func main() {
	var key int
	var id int
	var name string
	hashTable := &HashTable{}
	for {
		fmt.Println("*********** 雇员系统菜单 ***********")
		fmt.Println("*********** 1. 添加雇员 ***********")
		fmt.Println("*********** 2. 显示雇员 ***********")
		fmt.Println("*********** 3. 查找雇员 ***********")
		fmt.Println("*********** 4. 删除雇员 ***********")
		fmt.Println("*********** 5. 退出 ***********")
		fmt.Println("请输入(1-5):")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case 2:
			hashTable.ShowAll()
		case 3:
			fmt.Println("请输入要查找的雇员id")
			fmt.Scanln(&id)
			emp := hashTable.FindEmpById(id)
			if emp == nil {
				fmt.Printf("没有找到id:%v \n", id)
			} else {
				fmt.Println("找到了: ")
				emp.ShowEmpInfo()
			}
		case 4:
			fmt.Println("请输入要删除的雇员id")
			fmt.Scanln(&id)
			err := hashTable.DelEmpById(id)
			if err != nil {
				fmt.Printf("id: %v 删除失败, %v\n", id, err)
			} else {
				fmt.Printf("id: %v 已删除\n", id)
			}
		case 5:
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}

// 散列效果：
/*
链表：  0
emp[id]=14, emp[name]=wwwww =>
链表：  1
emp[id]=1, emp[name]=qqq => emp[id]=8, emp[name]=cccccc => emp[id]=15, emp[name]=ddd =>
链表： 2 为空
链表： 3 为空
链表： 4 为空
链表：  5
emp[id]=12, emp[name]=fffff =>
链表： 6 为空
*********** 雇员系统菜单 ***********
*********** 1. 添加雇员 ***********
*********** 2. 显示雇员 ***********
*********** 3. 找到雇员 ***********
*********** 4. 退出 ***********
请输入(1-4):
*/
