package leetcode

func bestTimeToBuyAndSell_1(prices []int) int {
	buyPrice := prices[0]
	maxProfit := 0
	for _, val := range prices {
		if val < buyPrice {
			buyPrice = val
		}
		res := val - buyPrice
		if maxProfit < res {
			maxProfit = res
		}
	}
	return maxProfit
}
