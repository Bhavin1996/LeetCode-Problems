package leetcode

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	sellPrice := 0
	end := len(prices)
	if len(prices) == 1 {
		return 0
	} else if len(prices) == 2 {
		if prices[0] < prices[1] {
			return prices[1] - prices[0]
		} else {
			return 0
		}
	}
	for i := 1; i < end; i++ {
		if prices[i] < prices[i-1] && prices[i] < buyPrice {
			if i == end-1 {
				return 0
			}
			buyPrice = prices[i]
		} else if prices[i] > prices[i-1] && prices[i] > sellPrice {
			sellPrice = prices[i]
		}
	}
	return sellPrice - buyPrice
}
