package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var (
		n1 int = m - 1
		n2 int = n - 1
		k int = n + m - 1
	)

	for n2 >= 0 {
		:
	}
}

// 1, 2, 3, 0, 0, 0
// 1, 5, 6

// 5, 6, 7, 0, 0, 0
// 1, 2, 3

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{5,6,6}, 3)
	merge([]int{1,2,7,0,0,0}, 3, []int{5,6,6}, 3)
	merge([]int{8,9,10,0,0,0}, 3, []int{5,6,6}, 3)
}
