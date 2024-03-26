package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var (
		min		int
		profit	int
		total	int
	)

	min = prices[0]
	for i := range prices {
		if prices[i] < min || (i > 0 && prices[i] < prices[i - 1]) {
			total += profit
			min = prices[i]
			profit = 0

		// } else if prices[i] - min >= profit {
		} else {
			profit = prices[i] - min

		}
	}
	return total + profit
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}), 7)
	fmt.Println(maxProfit([]int{1,2,3,4,5}), 4)
	fmt.Println(maxProfit([]int{7,6,4,3,1}), 0)
	fmt.Println(maxProfit([]int{7,1}), 0)
	fmt.Println(maxProfit([]int{7}), 0)
	fmt.Println(maxProfit([]int{2,1,2,0,1}), 2)
	fmt.Println(maxProfit([]int{2,4,1}), 2)
}

// 7,1,5,3,6,4
// if price < min : min = price 
// else if price - min > profit {
// 		profit = price - min
// } else {
// 		total_profit += profit
//		min = price
// }

// идем по массиву. если число меньше нашего то обновляем наше число
// если число больше нашего и разности между ними больше profit то обновляем profit
// если число меньше предыдущего то total_profit += profit profit = 0 min = price
