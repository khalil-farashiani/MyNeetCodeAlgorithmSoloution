package main

import "fmt"

type temperatureElement struct {
	data int
	min  int
	next *temperatureElement
}

type temperatureMinStack struct {
	head *temperatureElement
	size int
}

func temperatureConstructor() temperatureMinStack {
	return temperatureMinStack{}
}

func (this *temperatureMinStack) Push(val int) {
	element := new(temperatureElement)
	element.data = val
	temp := this.head
	element.next = temp
	if this.head == nil {
		this.head = element
		this.head.min = val
		this.size++
	} else if this.head.min > val {
		element.min = val
	} else {
		element.min = this.head.min
	}
	this.head = element
	this.size++
}

func (this *temperatureMinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
	this.size--
}

func (this *temperatureMinStack) Top() int {
	if this.head == nil {
		return 0
	}
	r := this.head.data
	return r
}

func (this *temperatureMinStack) GetMin() int {
	return this.head.min
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{} // a stack of indices
	for i, temp := range temperatures {
		// pop the stack until we find a smaller temperature or the stack is empty
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			// update the result array with the difference between the current index and the top of the stack
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// pop the stack
			stack = stack[:len(stack)-1]
		}
		// push the current index to the stack
		stack = append(stack, i)
	}
	return res
}

// https://leetcode.com/problems/daily-temperatures/submissions
func main() {
	fmt.Printf("%v", dailyTemperatures([]int{50, 40, 60, 30, 20, 90}))
}
