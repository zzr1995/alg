package tree

import (
	"container/list"
	"fmt"
	"math"
)

// 前序遍历
func preOrderTraverse(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.data) // 前序遍历就是从node开始遍历，所以要先打印
	preOrderTraverse(n.left)
	preOrderTraverse(n.right)
}

// 中序遍历
func inOrderTraverse(n *node) {
	if n == nil {
		return
	}
	inOrderTraverse(n.left)
	fmt.Printf("%v ", n.data)
	inOrderTraverse(n.right)
}

// 后序遍历
func postOrderTraverse(n *node) {
	if n == nil {
		return
	}
	postOrderTraverse(n.left)
	postOrderTraverse(n.right)
	fmt.Printf("%v ", n.data)
}

// 层序遍历：附带换行打印
/**
1 将各节点入队
2 循环执行以下操作，直到队列为空
	取出队头节点出队，进行访问
	将队头节点的左子节点入队
	将队头节点的右子节点入队
*/
func levelOrderTraverse(n *node) {

	if n == nil {
		fmt.Println("传入节点为空")
		return
	}

	// the first row line is root
	queue := list.New()
	queue.PushBack(n)
	levelLength := 1

	for queue.Len() != 0 {

		// 出队并打印
		queueHead := queue.Remove(queue.Front())
		tempNode := queueHead.(*node) // 类型断言
		fmt.Printf("%v ", tempNode.data)
		levelLength--
		// 入队
		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}
		// 更新节点的长度
		if levelLength == 0 {
			levelLength = queue.Len()
			fmt.Println() // 完成一行了，打个空行
		}
	}
}

// 插入时的递归函数
func insertRC(currentNode *node, insertNode *node) {

	if currentNode == nil || insertNode == nil {
		panic("非法节点")
	}

	// 如果插入数据和当前节点数据相同
	if currentNode.data == insertNode.data {
		currentNode = insertNode
		return
	}

	parentNode := currentNode

	if currentNode.data > insertNode.data { // 向左查找
		if currentNode.left == nil {
			insertNode.parent = parentNode
			currentNode.left = insertNode
		} else {
			parentNode = currentNode.left
			insertRC(currentNode.left, insertNode)
		}
	}

	if currentNode.data < insertNode.data { // 向右查找
		if currentNode.right == nil {
			insertNode.parent = parentNode
			currentNode.right = insertNode
		} else {
			parentNode = currentNode.right
			insertRC(currentNode.right, insertNode)
		}
	}
}

// 高度   math.Max
func heightByGe(n *node) int {
	if n == nil {
		return 0
	}
	r := 1 + math.Max(float64(heightByGe(n.left)), float64(heightByGe(n.right)))
	return int(r)
}

// 计算二叉树的高度：迭代方式
func heightByRC(n *node) int {

	if n == nil {
		return 0
	}

	// 层序遍历
	queue := list.New() // 制作一个队列
	queue.PushBack(n)

	height := 0      // 树的高度
	levelLength := 1 // 每层存储的元素个数

	for queue.Len() != 0 {

		queueHead := queue.Remove(queue.Front()) // 队首出队
		tempNode := queueHead.(*node)            // 类型断言

		levelLength--

		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}

		if levelLength == 0 { // 准备访问下一层
			levelLength = queue.Len()
			height++
		}
	}
	return height
}

// 判断一棵树是否是完全二叉树：即最后一层的叶子节点是向左对其
func isCompleteTree(n *node) bool {
	if n == nil {
		return false
	}

	// 层序遍历
	queue := list.New() // 制作一个队列
	queue.PushBack(n)

	isLeaf := false // 当前节点是否是叶子节点
	for queue.Len() != 0 {

		queueHead := queue.Remove(queue.Front()) // 队首出队
		tempNode := queueHead.(*node)            // 类型断言

		if tempNode.left == nil && tempNode.right != nil {
			return false
		}

		// 如果已经要求是叶子节点，但是又不是叶子节点
		if isLeaf && (tempNode.left == nil && tempNode.right == nil) {
			return false
		}

		// 要求下次遍历必须是叶子节点的情况
		if tempNode.right == nil || (tempNode.left != nil && tempNode.right == nil) {
			isLeaf = true
		}

		// 原层序遍历代码
		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}

	}

	return true
}

// 在遍历中加上一些操作
func invertBinaryTree(n *node) *node {

	if n == nil {
		return n
	}

	// 交换
	tempNode := n.left
	n.left = n.right
	n.right = tempNode

	invertBinaryTree(n.left)
	invertBinaryTree(n.right)

	return n
}

// 查找前驱节点：即二叉树在中序遍历时，当前元素的前一个节点，二叉搜索树的前驱节点也是当前节点前的最大节点
func findPrevNode(n *node) *node {

	// 情况一：当前节点为空
	if n == nil {
		fmt.Println("传入节点为空")
		return nil
	}

	// 情况二：当前节点的左子节点为空，父节点也为空
	if n.left == nil && n.parent == nil {
		fmt.Println("无前驱节点")
		return nil
	}

	// 定义当前循环到了哪个节点
	var currentNode *node

	// 情况三：当前节点的左子节点为空，父节点不为空，前驱在在parent的右子树中
	if n.left == nil && n.parent != nil {
		currentNode = n
		for currentNode.parent != nil && currentNode == currentNode.parent.left {
			currentNode = currentNode.parent
		}
	} else { // 最后一个情况： n.left != nil
		currentNode = n.left
		for currentNode.left != nil {
			currentNode = currentNode.right
		}
	}

	return currentNode
}

// 查找后继节点：即二叉树在中序遍历时，当前元素的后一个节点，二叉搜索树的后继节点也是当前节点后的最小节点
func findNextNode(n *node) *node {

	// 情况一：当前节点为空
	if n == nil {
		fmt.Println("传入节点为空")
		return nil
	}

	// 情况二：当前节点的左子节点为空，父节点也为空
	if n.right == nil && n.parent == nil {
		fmt.Println("无前驱节点")
		return nil
	}

	// 定义当前循环到了哪个节点
	var currentNode *node

	// 情况三：当前节点的左子节点为空，父节点不为空，前驱在在parent的右子树中
	if n.right == nil && n.parent != nil {
		currentNode = n
		for currentNode.parent != nil && currentNode == currentNode.parent.right {
			currentNode = currentNode.parent
		}
	} else { // 最后一个情况： n.right != nil
		currentNode = n.right
		for currentNode.right != nil {
			currentNode = currentNode.left
		}
	}

	return currentNode
}
