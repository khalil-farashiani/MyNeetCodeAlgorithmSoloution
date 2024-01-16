package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	// Set low and high boundaries
	low, high := 0, len(arr)-1

	for low <= high { // Continue search as long as low <= high
		// Calculate middle index
		mid := low + (high-low)/2
		// Check if element is found
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			// If target is greater, focus on the right half
			low = mid + 1
		} else {
			// If target is less, focus on the left half
			high = mid - 1
		}
	}

	// If element not found, return -1
	return -1
}

func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		if j := binarySearch(numbers[i+1:], target-num); j != -1 {
			return []int{i + 1, i + j + 2}
		}
	}
	return nil
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
func main() {
	fmt.Printf("%v", twoSum([]int{2, 3, 4}, 6))
}
