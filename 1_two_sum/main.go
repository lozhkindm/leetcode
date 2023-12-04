package main

import (
	"runtime"
)

func main() {
	twoSum([]int{3, 1, 4, 3, 5, 6, 7}, 6)
}

func twoSum(nums []int, target int) []int {
	runtime.GC()
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		toFind := target - num
		if _, ok := m[toFind]; ok {
			return []int{m[toFind], i}
		}
		m[num] = i
	}
	return []int{}
}
