/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.
The solution to this problem is quite simple actually as we need only to check for bigger price to pre_ele.
In this problem what we do it we are gonna iterate over the prices and if the curr_elem > pre_element then we calculate
the substraction of these element and add it into sum (i.e. MaxProfit). That's it!
*/

package leetcode

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 1 {
		return 0
	} else if len(prices) == 2 {
		if prices[1] > prices[0] {
			return (prices[1] - prices[0])
		} else {
			return 0
		}
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}
