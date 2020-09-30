package tree

import "fmt"

// init tree
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
	for currentNode != nil {
		// 被插入节点的父节点是当前节点
		if data < currentNode.data { // 左
			if currentNode.left == nil {
				insertNode.parent = currentNode
				currentNode.left = insertNode
				return
			} else {
				currentNode = currentNode.left
			}
		} else { // 右
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

// delete or update depend on find
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

func (t *bst) remove(data int) bool {
	delNode := t.find(data)
	if delNode == nil {
		fmt.Println("can not find this node")
		return false
	}
	if delNode.left != nil && delNode.right != nil {
		// two children
	} else if delNode.left == nil && delNode.right == nil {
		// no children
		t.removeZeroNode(delNode)
	} else {
		// one children
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
	// zero or two
	if (n.left == nil && n.right == nil) || (n.left != nil && n.right != nil) {
		panic("has two or no children")
	}

	// left or right
	var nextNode *node
	if n.left != nil {
		nextNode = n.left
	} else {
		nextNode = n.right
	}

	// delNode is root
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

// 前驱节点:以其左孩子为根的子树的最大结点
// 后继节点:以其右孩子为根的子树的最小结点
