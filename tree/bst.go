package tree

import "fmt"

// 二叉搜索树
func NewBst(data int) *bst {
	return &bst{
		root: nil,
	}
}

// insert
func (t *bst) insert(data int) {
	if t.root == nil {
		root := &node{
			data:   data,
			left:   nil,
			right:  nil,
			parent: nil,
		}
		t.root = root
		return
	}
	// 开始插入
	currentNode := t.root
	insertNode := &node{
		data:   data,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	// 插左还是插右
	for currentNode != nil {
		// 被插入节点的父节点是当前节点
		if data < currentNode.data {
			if currentNode.left == nil {
				insertNode.parent = currentNode
				currentNode.left = insertNode
				return
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				insertNode.parent = currentNode
				currentNode.right = insertNode
				return
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

// find is easy
func (t *bst) find(data int) *node {
	currentNode := t.root
	for currentNode != nil {
		if data == currentNode.data {
			return currentNode
		}
		if data < currentNode.data {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return currentNode
}

// 删除零个孩子和一个孩子的节点都比较简单，删除有两个孩子的节点比较复杂
func (t *bst) remove(data int) bool {
	// 先找出要删除的节点
	delNode := t.find(data)
	// 删除的为空节点
	if delNode == nil {
		fmt.Println("can not find this node")
		return false
	}
	// 根据要删除的节点有多个孩子执行不用操作
	if delNode.left == nil && delNode.right == nil {
		t.removeZeroNode(delNode)
	} else if delNode.left != nil && delNode.right != nil {
		// 两个孩子则比较复杂

	} else {
		t.removeOneNode(delNode)
	}

	return true
}

// delete node has no children
func (t *bst) removeZeroNode(n *node) {
	// if node has children
	if n.left != nil || n.right != nil {
		panic("has children!!")
	}

	// node is root
	if n == t.root {
		t.root = nil
	}

	// left or right
	if n.parent.left == n {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
}

// delete node has one children
func (t *bst) removeOneNode(n *node) {
	// 并非只有一个孩子
	if (n.left == nil && n.right == nil) || (n.left != nil && n.right != nil) {
		panic("has two or no children")
	}

	// 判断该节点的孩子在左还是右
	var nextNode *node
	if n.left != nil {
		nextNode = n.left
	} else {
		nextNode = n.right
	}

	// 删除root节点
	if n == t.root {
		t.root = nextNode
	}

	// left or right
	if n.parent.left == n {
		n.parent.left = nextNode
	} else {
		n.parent.right = nextNode
	}

}
