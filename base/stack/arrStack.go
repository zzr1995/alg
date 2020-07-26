package main

import (
"errors"
"fmt"
)

//使用结构体管理栈
type sliceStack struct {
	arr       []interface{} //切片
	stackSize int           //栈中元素的个数
}

//创建栈
func newSliceStack() *sliceStack {
	return &sliceStack{arr: make([]interface{}, 0)}
}

//判断是否为空
func (p *sliceStack) isEmpty() bool {
	return p.stackSize == 0
}

//栈的大小
func (p *sliceStack) size() int {
	return p.stackSize
}

//返回栈顶元素
func (p *sliceStack) top() interface{} {
	if p.isEmpty() { //说明栈为空
		return nil
	}
	return p.arr[p.stackSize-1] //栈顶元素
}

//push栈元素
func (p *sliceStack) push(t interface{}) {
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}

//pop栈元素
func (p *sliceStack) pop() interface{} {
	if p.stackSize > 0 { //栈不为空时
		p.stackSize--
		element := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return element
	}
	return errors.New("栈为空")
}

//显示栈的元素
func (p *sliceStack) print() {
	fmt.Println("栈当前的情况是")
	if p.isEmpty() {
		fmt.Println("当前栈为空")
	}
	for i := p.stackSize - 1; i > 0; i-- {
		fmt.Println(i, "->", p.arr[i])
	}
}
func main() {
	var stack = newSliceStack()
	if stack.isEmpty() {
		fmt.Println("Stack为空! ")
	} else {
		fmt.Println("Stack不为空! ", stack.stackSize)
	}

	fmt.Println("\nstack增加字符串元素: ")
	stack.push("中文元素")
	stack.push("elem1")
	stack.push(10)
	stack.push(20)
	stack.push(30)

	stack.print()

	fmt.Println("当前Size = ", stack.size())

	fmt.Println("当前Top = ", stack.top())
	fmt.Println("取出元素", stack.pop())
	stack.print()
	//结果：
	//Stack为空!
	//
	//stack增加字符串元素:
	//栈当前的情况是
	//4 -> 30
	//3 -> 20
	//2 -> 10
	//1 -> elem1
	//当前Size =  5
	//当前Top =  30
	//取出元素 30
	//栈当前的情况是
	//3 -> 20
	//2 -> 10
	//1 -> elem1
}
