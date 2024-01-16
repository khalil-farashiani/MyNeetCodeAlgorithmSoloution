package main

import (
	"fmt"
	"strings"
)

func strip(s string) string {
	var result strings.Builder
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func isPalindrome(s string) bool {
	s = strip(s)
	point2 := len(s) - 1
	for _, ch := range s {
		if ch != int32(s[point2]) {
			return false
		}
		point2--
	}
	return true
}

// https://leetcode.com/problems/valid-palindrome
func main() {
	fmt.Printf("%v", isPalindrome("A man, a plan, a canal: Panama"))
}
