# https://leetcode.com/problems/number-of-1-bits/

class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        binary_n = bin(n)[2:]
        count = 0
        for bit in binary_n:
            if bit == '1':
                count += 1
        return count
