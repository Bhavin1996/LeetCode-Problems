/* In this problem we can buy and sell only one time so little bit different from it's 2 nd type.
so approach for this is like this we iterate over the prices array and we do a substraction from lowest buy price.
So set buyprice for 1st element and maxprofit as 0. Then follow above approach and solve it by iterating over it.
*/

package leetcode

func bestTimeToBuyAndSell_1(prices []int) int {
	buyPrice := prices[0]
	maxProfit := 0
	for _, val := range prices {
		if val < buyPrice { // This makes sure that we always have buyPrice at min val
			buyPrice = val
		}
		res := val - buyPrice
		if maxProfit < res {
			maxProfit = res
		}
	}
	return maxProfit
}
