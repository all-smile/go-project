// 二叉树

package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero // 指向左边节点的指针
	Right *Hero // 指向右边节点的指针
}

// 前序遍历 - 先输出根节点，然后输出左子树，最后输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		// 递归调用，先根再左后右
		fmt.Printf("No=%v, Name=%v \n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历 - 先输出左子树，然后输出根节点，最后输出右子树
func InfixOrder(node *Hero) {
	if node != nil {
		// 递归调用，先左再根后右
		InfixOrder(node.Left)
		fmt.Printf("No=%v, Name=%v \n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// 后序遍历 - 先输出左子树，然后输出右子树，最后输出根节点
func PostOrder(node *Hero) {
	if node != nil {
		// 递归调用，先左再右后根
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("No=%v, Name=%v \n", node.No, node.Name)
	}
}

/*
层序遍历：
借助队列先进先出的特性，从树的根结点开始，依次将其左子节点和右子节点入队。
而后每次队列中一个结点出队，都将其左子节点和右子节点入队，直到树中所有结点都出队，
出队结点的先后顺序就是层次遍历的最终结果
*/
func levelOrder(node *Hero) [][]int {
	ret := [][]int{} // 定义一个存放遍历之后的数组
	fmt.Printf("res类型： %T \n", ret)
	if node == nil {
		return ret
	}
	q := []*Hero{node} // 将整颗树怼进队列里
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*Hero{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.No)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	// 构建二叉树
	root := &Hero{
		No:   1,
		Name: "xixi",
	}
	left1 := &Hero{
		No:   2,
		Name: "haha",
	}
	left2 := &Hero{
		No:   10,
		Name: "1010",
	}
	right02 := &Hero{
		No:   11,
		Name: "1111",
	}
	right1 := &Hero{
		No:   3,
		Name: "hehe",
	}
	right2 := &Hero{
		No:   4,
		Name: "heihei",
	}
	right5 := &Hero{
		No:   5,
		Name: "5555",
	}
	root.Left = left1
	root.Right = right1
	right1.Right = right2

	left1.Left = left2
	left1.Right = right02

	right1.Left = right5

	fmt.Println("前序遍历---")
	// 前序遍历
	PreOrder(root)

	fmt.Println("中序遍历---")
	// 中序遍历
	InfixOrder(root)

	// 后续遍历
	fmt.Println("后续遍历---")
	PostOrder(root)

	fmt.Println("层序遍历---")
	arr := levelOrder(root)
	fmt.Println("arr = ", arr)
}

// 前序遍历
/*
No=1, Name=xixi
No=2, Name=haha
No=10, Name=1010
No=11, Name=1111
No=3, Name=hehe
No=5, Name=5555
No=4, Name=heihei
*/

// 中序遍历
/*
No=10, Name=1010
No=2, Name=haha
No=11, Name=1111
No=1, Name=xixi
No=5, Name=5555
No=3, Name=hehe
No=4, Name=heihei
*/

// 后序遍历
/*
 No=10, Name=1010
No=11, Name=1111
No=2, Name=haha
No=5, Name=5555
No=4, Name=heihei
No=3, Name=hehe
No=1, Name=xixi
*/

// 层序遍历---
/* res类型： [][]int
arr =  [[1] [2 3] [10 11 5 4]] */
