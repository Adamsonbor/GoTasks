package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var (
		indx int16 = 1
		val int
	)

	if len(nums) == 0 {
		return 0
	}

	val = nums[0]
	for _, v := range nums {
		if v != val {
			nums[indx] = v
			indx++
			val = v
		}
	}
	return int(indx)
}

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
	fmt.Println(removeDuplicates([]int{0}))
	fmt.Println(removeDuplicates([]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}))
}
