# https://leetcode.com/problems/power-of-three/

class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        power = 0
        while 3 ** power < n:
            power += 1
        return 3 ** power == n
