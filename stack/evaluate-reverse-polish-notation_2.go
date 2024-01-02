package main

import (
	"fmt"
	"strconv"
)

type myStack []int

func (s *myStack) Push(v int) {
	*s = append(*s, v)
}

func (s *myStack) Pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return ret
}

func evalRPN2(tokens []string) int {
	var stack myStack
	for _, token := range tokens {
		var num1, num2, res int
		switch token {
		case "+":
			num1 = stack.Pop()
			num2 = stack.Pop()
			res = num2 + num1
			stack.Push(res)
		case "-":
			num1 = stack.Pop()
			num2 = stack.Pop()
			res = num2 - num1
			stack.Push(res)
		case "*":
			num1 = stack.Pop()
			num2 = stack.Pop()
			res = num2 * num1
			stack.Push(res)
		case "/":
			num1 = stack.Pop()
			num2 = stack.Pop()
			res = num2 / num1
			stack.Push(res)
		default:
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}
	return stack.Pop()
}

// https://leetcode.com/problems/evaluate-reverse-polish-notation
func main() {
	fmt.Printf("%v", evalRPN2([]string{"2", "1", "+", "3", "*"}))
}
