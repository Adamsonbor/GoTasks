package main

import (
	"fmt"
)

func reverse(arr []int, beg int, end int) {
	var (
		tmp int
	)

	for beg < end {
		tmp = arr[end]
		arr[end] = arr[beg]
		arr[beg] = tmp
		end --
		beg ++
	}
}

func rotate(nums []int, k int)  {
	var length int = len(nums)

	k = k % length
	reverse(nums, 0, length - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, length - 1)
}

func main() {
	rotate([]int{1,2,3,4,5,6,7,8,9}, 10)
	rotate([]int{-1}, 2)
	rotate([]int{-1,2,3,4,5}, 2)
}


// [1,2,3,4,5,6,7] 3
// [5,6,7,4,1,2,3]
// [5,6,7,4,1,2,3]
// [5,6,7,1,2,3,4]

// [1,2,3,4,5,6,7,8,9,10,11] 3
// [11,10,9,8,7,6,5,4,3,2,1]
// [9,10,11,8,7,6,5,4,3,2,1]
// [9,10,11,1,2,3,4,5,6,7,8]

