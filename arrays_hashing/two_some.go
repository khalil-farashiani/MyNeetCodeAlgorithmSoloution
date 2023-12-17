package main

import "fmt"

func twoSum(nums []int, target int) []int {
	complement := make(map[int]int)
	for i, num := range nums {
		if j, ok := complement[num]; ok {
			return []int{i, j}
		}
		complement[target-num] = i
	}
	return nil
}

// 2 for soloution
// func twoSum(nums []int, target int) []int {
// 	for index, num := range nums {
// 		for i := index + 1; i < len(nums); i++ {
// 			if num+nums[i] == target {
// 				return []int{i, index}
// 			}
// 		}
// 	}
// 	return []int{0, 0}
// }

// https://leetcode.com/problems/two-sum/
func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 9))
}
