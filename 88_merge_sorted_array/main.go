package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[m+n+1] = nums1[m]
			nums1[m] = 0
			m--
		} else {
			nums1[m+n+1] = nums2[n]
			n--
		}
	}
	for ; n >= 0; n-- {
		nums1[n] = nums2[n]
	}
}
