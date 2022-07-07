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

// 省去buy變數
func (p Problem) MaxProfitFollowUp(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
