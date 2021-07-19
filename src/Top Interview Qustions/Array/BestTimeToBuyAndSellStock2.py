# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        max_profit = 0
        if len(prices) > 1:
            i = 0
            while i < len(prices) - 1:
                profit = prices[i + 1] - prices[i]
                i += 1
                if profit >= 0:
                    while i < len(prices) - 1 and prices[i] < prices[i + 1]:
                        profit += prices[i + 1] - prices[i]
                        i += 1
                    max_profit += profit
        return max_profit
