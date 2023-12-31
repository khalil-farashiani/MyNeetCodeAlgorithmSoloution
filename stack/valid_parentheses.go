package main

import "fmt"

type Stack []rune

func (s *Stack) push(v rune) {
	*s = append(*s, v)
}

func (s *Stack) pop() rune {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return ret
}

func isValid(s string) bool {
	myStack := Stack{}
	for _, ch := range s {
		if ch == '[' || ch == '(' || ch == '{' {
			myStack.push(ch)
		} else {
			if len(myStack) == 0 {
				return false
			}
			rr := myStack.pop()
			if rr == '[' && ch != ']' ||
				rr == '{' && ch != '}' ||
				rr == '(' && ch != ')' {
				return false
			}
		}
	}
	if len(myStack) == 0 {
		return true
	}
	return false
}

// https://leetcode.com/problems/valid-parentheses
func main() {
	fmt.Println(isValid("]"))
}
