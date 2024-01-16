package main

import (
	"sort"
)

type myMap map[int][2]int

func (m myMap) Len() int          { return len(m) }
func (m myMap) Swap(i int, j int) { (m)[i], (m)[j] = (m)[j], (m)[i] }
func (m myMap) Less(i int, j int) bool {
	if (m)[i][0] != (m)[j][0] {
		return (m)[i][0] > (m)[j][0]
	}
	return (m)[i][1] > (m)[j][1]
}

func carFleet(target int, position []int, speed []int) int {
	var carMap = myMap{}
	for i, pos := range position {
		carMap[i] = [2]int{pos, speed[i]}
	}
	sort.Sort(carMap)

	i := 0
	flag := false
	for {
		if i == len(position) && flag {
			break
		}
		if i == len(position) {
			i = 0
		}
		car := carMap[i]
		car[0] += car[1]
		carMap[i] = car
		if carMap[0][0] >= target {
			flag = true
		}
		i++
	}

	res := make([][2]int, len(carMap))
	for k, v := range carMap {
		res[k] = v
	}
	count := 1
	lastPos := res[0][0]
	for i, v := range res {
		if i == 0 {
			continue
		}
		if v[0] >= target && lastPos >= target {
			continue
		}
		if v[0] < lastPos && v[0] < target {
			count++
			lastPos = v[0]
		}
	}
	return count
}

func main() {
	print(carFleet(12, []int{15, 4, 10, 12, 17, 19, 11, 14, 6, 0, 2}, []int{4, 8, 1, 2, 8, 8, 10, 7, 10, 8, 6}))
}
