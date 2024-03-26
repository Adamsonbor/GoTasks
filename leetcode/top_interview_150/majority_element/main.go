package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	var (
		count int
		value int
	)

	for _, v := range nums {
		if count == 0 {
			value = v
			count ++
		} else if value == v {
			count ++
		} else {
			count --
		}
	}
	return value
}

func main() {
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}
