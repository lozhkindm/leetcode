package main

func main() {
	removeDuplicates([]int{1, 2, 3})
}

func removeDuplicates(nums []int) int {
	var l, r int
	for r = 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}
