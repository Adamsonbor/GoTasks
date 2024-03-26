package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
    var (
		indx int8 = 0
	)

	for _, v := range nums {
		if v != val {
			nums[indx] = v
			indx++
		}
	}
	return int(indx)
}

func main() {
	fmt.Println(removeElement([]int{3,2,2,3}, 2))
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
	fmt.Println(removeElement([]int{3,2,2,3}, 2))
}
