package main

import (
	"container/heap"
	"fmt"
)

type NumFreq struct {
	num       int
	frequency int
}

type NumFreqHeap []NumFreq

func (h *NumFreqHeap) Len() int {
	return len(*h)
}

func (h *NumFreqHeap) Less(i, j int) bool {
	return (*h)[i].frequency < (*h)[j].frequency
}

func (h *NumFreqHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NumFreqHeap) Push(x interface{}) {
	*h = append(*h, x.(NumFreq))
}

func (h *NumFreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	freqHeap := make(NumFreqHeap, 0)
	heap.Init(&freqHeap)

	for num, freq := range freqMap {
		heap.Push(&freqHeap, NumFreq{num, freq})
		if len(freqHeap) > k {
			heap.Pop(&freqHeap)
		}
	}

	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(&freqHeap).(NumFreq).num
	}

	return result
}

//https://leetcode.com/problems/top-k-frequent-elements
func main() {
	fmt.Printf("%v", topKFrequent([]int{1, 2}, 2))
}
