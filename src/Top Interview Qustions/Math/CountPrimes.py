# https://leetcode.com/problems/count-primes/

class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """
        isPrime = [True for _ in range(n)]
        i = 2
        while i ** 2 <= n:
            i += 1
            if not isPrime[i - 1]:
                continue
            j = (i - 1) ** 2
            while j < n:
                isPrime[j] = False
                j += i - 1

        count = 0
        for i in range(2, n):
            if isPrime[i]:
                count += 1
        return count
