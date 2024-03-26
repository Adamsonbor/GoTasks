package main

import (
	"fmt"
)

// MY SOLUTION

// func removeDuplicates(nums []int) int {
// 	var (
// 		indx			int16
// 		last_value		int
// 		num_values		int16
// 	)

// 	last_value = nums[0]
// 	num_values = 0
// 	indx = 0
// 	for _, value := range nums {
// 		if value != last_value {
// 			nums[indx] = value
// 			last_value = value
// 			num_values = 1
// 			indx++

// 		} else {
// 			num_values ++

// 			if num_values < 3 {
// 				nums[indx] = value
// 				indx ++
// 			}
// 		}
// 	}
// 	return int(indx)
// }

// =========================================

// SOLUTION FROM ALIENS
func removeDuplicates(nums []int) int {
	var (
		i int
	)

	for _, v := range nums {
		if i < 2 || v > nums[i - 2] {
			nums[i] = v
			i++
		}
	}
	return i
}

func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,3}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,1,2,3,3}))
}
