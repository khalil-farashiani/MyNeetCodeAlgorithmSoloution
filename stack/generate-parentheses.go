package main

import (
	"fmt"
	"strings"
)

// only add open parenthesis when open < n
// only add close parenthesis when close < open
// valid if openN == closeN == n
func generateParenthesis(n int) []string {
	var stack []string
	var res []string

	var backtrack func(int, int)
	backtrack = func(openN int, closedN int) {
		if openN == n && closedN == n && openN == closedN {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			pop(&stack)
		}
		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			pop(&stack)
		}
	}
	backtrack(0, 0)
	return res

}

func pop(list *[]string) {
	//since pop() function in go doesn't exist
	length := len(*list)
	*list = (*list)[:length-1]
}

// https://leetcode.com/problems/generate-parentheses/
func main() {
	fmt.Printf("%v", generateParenthesis(8))
}
