package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var (
		n1 int32 = int32(m) - 1
		n2 int32 = int32(n) - 1
		k int32 = int32(n) + int32(m) - 1
	)

	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[k] = nums1[n1]
			n1--
		} else {
			nums1[k] = nums2[n2]
			n2--
		}
		k--
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{5,6,6}, 3)
	merge([]int{1,2,7,0,0,0}, 3, []int{5,6,6}, 3)
	merge([]int{8,9,10,0,0,0}, 3, []int{5,6,6}, 3)
}
