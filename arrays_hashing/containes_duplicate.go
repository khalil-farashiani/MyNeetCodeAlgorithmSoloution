package main

func containsDuplicate(nums []int) bool {
	duplicateCounter := make(map[int]bool, len(nums))
	for _, num := range nums {
		if duplicateCounter[num] {
			return true
		}
		duplicateCounter[num] = true
	}
	return false
}

// https://leetcode.com/problems/contains-duplicate/
func main() {
	// testcase
	println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
