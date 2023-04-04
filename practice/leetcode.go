package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	// 说白了就是找一个数组中两个差值最大的元素 并返回差值
	// 不过有个条件，就是被减数一定要在减数后面
	// btw.现在看见一道题就想用双指针...4.4
	var temp, profit int
	for ind := 0; ind < len(prices)-1; ind++ {
		// 先判断相邻的下一个数是不是比自己小，如果比自己小，这轮循环可以忽略了
		if prices[ind+1] <= prices[ind] {
			fmt.Println(prices[ind+1], "<=", prices[ind])
			continue
		}

		fmt.Print(prices[ind+1], ">=", prices[ind], " ")
		temp = getMaxNum(prices[ind+1:]) - prices[ind]
		fmt.Println("ind|", ind, "-> temp|", temp)
		if temp > profit {
			profit = temp
		}
	}
	return profit
}

// 获取一个价格数组中的最大价格
func getMaxNum(prices []int) int {
	max := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		}
	}
	return max
}

// 获取最大价格
func getMaxProfit(prices []int) int {
	
	return 0
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(nums[0:1])
	//fmt.Println("finally res|", maxProfit(nums))
}
