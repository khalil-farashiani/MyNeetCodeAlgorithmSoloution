package main

import (
	"fmt"
	"sort"
)

type myStruct []struct {
	num int
	rep int
}

func (m myStruct) Len() int {
	return len(m)
}
func (m myStruct) Less(i, j int) bool {
	return m[i].rep > m[j].rep
}
func (m myStruct) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func findIsINmyStruct(s myStruct, i int) (int, bool) {
	for index, item := range s {
		if i == item.num {
			return index, true
		}
	}
	return 0, false
}

func topKFrequent(nums []int, k int) []int {
	var result []int
	freq := make(myStruct, 0, len(nums))
	for _, num := range nums {
		if index, ok := findIsINmyStruct(freq, num); ok {
			freq[index].rep++
		} else {
			freq = append(freq, struct {
				num int
				rep int
			}{num: num, rep: 1})
		}
	}
	sort.Sort(freq)
	if len(freq) < k {
		for _, fr := range freq {
			result = append(result, fr.num)
		}
	} else {
		for i, fr := range freq {
			result = append(result, fr.num)
			if i+1 == k {
				break
			}
		}
	}
	return result
}

//https://leetcode.com/problems/top-k-frequent-elements
func main() {
	fmt.Printf("%v", topKFrequent([]int{1, 2}, 2))
}
