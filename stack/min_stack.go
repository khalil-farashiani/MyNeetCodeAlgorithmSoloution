package main

import "fmt"

type element struct {
	data int
	min  int
	next *element
}

type MinStack struct {
	head *element
	size int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	element := new(element)
	element.data = val
	temp := this.head
	element.next = temp
	if this.head == nil {
		this.head = element
		this.head.min = val
		this.size++
		return
	} else if this.head.min > val {
		element.min = val
	} else {
		element.min = this.head.min
	}
	this.head = element
	this.size++
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
	this.size--
}

func (this *MinStack) Top() int {
	if this.head == nil {
		return 0
	}
	r := this.head.data
	return r
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

// https://leetcode.com/problems/min-stack/
func main() {
	//["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
	//[[],[2],[0],[3],[0],[],[],[],[],[],[],[]]
	obj := Constructor()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	min1 := obj.GetMin()
	fmt.Println(min1)
	obj.Pop()
	min2 := obj.GetMin()
	fmt.Println(min2)
	obj.Pop()
	min3 := obj.GetMin()
	fmt.Println(min3)
	obj.Pop()
	min4 := obj.GetMin()
	fmt.Println(min4)
}
