package main

func longestConsecutive(nums []int) int {
	mm := make(map[int]bool, len(nums))
	for _, num := range nums {
		mm[num] = true
	}
	maax := 0
	for k, _ := range mm {
		if !mm[k-1] {
			// so this is starter
			length := 0
			for {
				length++
				if !mm[k+length] {
					break
				}
			}
			maax = max(maax, length)
		}
	}
	return maax
}

// https://leetcode.com/problems/longest-consecutive-sequence
func main() {
	// note that we use map instead of set because we don't have set internally in golang
	longestConsecutive([]int{100, 4, 200, 1, 3, 3, 2})
}
