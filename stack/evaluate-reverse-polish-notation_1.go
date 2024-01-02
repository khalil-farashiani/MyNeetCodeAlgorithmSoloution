package main

import (
	"fmt"
	"strconv"
)

type stackElement struct {
	data int
	next *stackElement
}

type MyStack struct {
	head *stackElement
}

func MyStackConstructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(val int) {
	element := new(stackElement)
	element.data = val
	temp := m.head
	element.next = temp
	m.head = element
}

func (m *MyStack) Pop() int {
	if m.head == nil {
		return 0
	}
	r := m.head.data
	m.head = m.head.next
	return r
}

func evalRPN(tokens []string) int {
	stack := MyStackConstructor()
	for _, token := range tokens {
		if token == "+" {
			num1 := stack.Pop()
			num2 := stack.Pop()
			res := num2 + num1
			stack.Push(res)
		} else if token == "-" {
			num1 := stack.Pop()
			num2 := stack.Pop()
			res := num2 - num1
			stack.Push(res)
		} else if token == "*" {
			num1 := stack.Pop()
			num2 := stack.Pop()
			res := num2 * num1
			stack.Push(res)
		} else if token == "/" {
			num1 := stack.Pop()
			num2 := stack.Pop()
			res := num2 / num1
			stack.Push(res)
		} else {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}
	return stack.Pop()
}

// https://leetcode.com/problems/evaluate-reverse-polish-notation
func main() {
	fmt.Printf("%v", evalRPN([]string{"2", "1", "+", "3", "*"}))
}
