package main

import "fmt"

func productExceptSelf(nums []int) []int {
	//We first calculate right products of index then left products * right products
	sliceSize := len(nums)
	result := make([]int, sliceSize, sliceSize)
	result[0], result[sliceSize-1] = 1, 1

	for i := 1; i < sliceSize; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := sliceSize - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		result[i] *= rightProduct
	}

	return result
}

//https://leetcode.com/problems/product-of-array-except-self
func main() {
	fmt.Printf("%v", productExceptSelf([]int{1, 2, 3, 5, 2}))
}
