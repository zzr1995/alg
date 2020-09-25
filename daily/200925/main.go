package main

import "fmt"

// 普通二叉树，需要手动插
func main() {
	root := Node{Value: 3}
	root.Left = &Node{}
	root.Left.SetValue(0)
	root.Left.Right = CreateNode(2)
	root.Right = &Node{5, nil, nil}
	root.Right.Left = CreateNode(4)
	root.Order()
}

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Value = v
}

// 生左右孩子
func CreateNode(v int) *Node {
	return &Node{Value: v}
}

// dfs
func (node *Node) Order() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.Order()
	node.Right.Order()
}
