# https://leetcode.com/problems/climbing-stairs/

class Solution(object):
    def __init__(self):
        self.n_values = {1: 1, 2: 2}

    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n in self.n_values:
            return self.n_values[n]
        self.n_values[n] = self.climbStairs(n - 1) + self.climbStairs(n - 2)
        return self.n_values[n]
