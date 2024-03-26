package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var (
		cur		int
		prof	int
	)
	
	for i := range prices {
		if prices[i] < prices[cur] {
			cur = i
		} else if prices[i] - prices[cur] > prof {
			prof = prices[i] - prices[cur]
		}
	}

	return prof
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{1,2}))
	fmt.Println(maxProfit([]int{2,4,1}))
}

// [7,1,5,3,6,4]
// [7,3,9,1,5,4]
// 
