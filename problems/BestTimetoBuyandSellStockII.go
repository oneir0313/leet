package problems

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func (p Problem) MaxProfit(prices []int) int {
	buy, profit := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[buy] {
			profit += prices[i] - prices[buy]
		}
		buy = i
	}
	return profit
}
