# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if len(prices) <= 1:
            return 0
        cur_min, cur_max, max_profit = prices[0], prices[0], 0
        for i in prices:
            if i < cur_min:
                cur_min = i
                cur_max = i
            elif i > cur_max:
                cur_max = i
            max_profit = max(max_profit, cur_max - cur_min)
        return max_profit
