package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	var (
		last_index	int = len(nums) - 1
		steps		int = nums[0]
		indx		int
	)

	for indx < last_index && steps > 0 {
		steps --
		indx ++
		if nums[indx] > steps {
			steps = nums[indx]
		}
	}

	if indx == last_index {
		return true
	}
	return false
}

func test(nums []int, expected bool) {
	if expected == canJump(nums) {
		fmt.Println("\033[32mOK\033[0m")
	} else {
		fmt.Println("\033[31mKO\033[0m")
	}
}

func main() {
	test([]int{2,3,1,1,4}, true)
	test([]int{3,2,1,0,4}, false)
	test([]int{2,0}, true)
	test([]int{2,5,0,0}, true)
	test([]int{1,2,3}, true)
	test([]int{3}, true)
	test([]int{0,3}, false)
}
