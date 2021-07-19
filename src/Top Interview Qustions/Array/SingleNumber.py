# https://leetcode.com/problems/single-number/

from collections import Counter


class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cntr = Counter(nums)
        for i in nums:
            if cntr[i] == 1:
                return i
