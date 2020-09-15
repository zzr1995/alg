package stack

import (
	"fmt"
)

// 节点
type Node struct {
	data interface{}
	next *Node
}

// 链栈结构体
type LinkedStack struct {
	top    *Node // 栈顶指针
	length int   // 栈的长度
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		top:    nil,
		length: 0,
	}
}

// 压栈操作
func (s *LinkedStack) Push(e interface{}) {
	node := &Node{
		data: e,
		next: s.top, // 指向之前的栈顶元素
	}
	s.top = node
	s.length++
}

// 弹栈操作
func (s *LinkedStack) Pop() interface{} {

	if s.top == nil {
		fmt.Println("空栈无法pop")
		return nil
	}
	e :=s.top.data
	s.top = s.top.next
	s.length--
	return e
}

// 查看栈顶元素
func (s *LinkedStack) Top() interface{} {
	if s.top == nil {
		fmt.Println("空栈无栈顶元素")
		return nil
	}
	return s.top
}
