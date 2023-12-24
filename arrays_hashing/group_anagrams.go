package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sorted := string(runes)
		groups[sorted] = append(groups[sorted], str)
	}
	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// this have O(n^2) time complexity
// func groupAnagrams(strs []string) [][]string {
// 	var result [][]string
// 	var CalculatedIndex = make(map[int]bool)
// 	for i, str := range strs {
// 		if CalculatedIndex[i] {
// 			continue
// 		}
// 		anagrams := []string{str}
// 		for j := i + 1; j < len(strs); j++ {
// 			if isAnagram(strs[j], str) {
// 				CalculatedIndex[j] = true
// 				anagrams = append(anagrams, strs[j])
// 			}
// 		}
// 		result = append(result, anagrams)
// 	}
// 	return result
// }

// https://leetcode.com/problems/group-anagrams
func main() {
	fmt.Printf("%v", groupAnagrams([]string{"rat", "tar", "rar"}))
}
