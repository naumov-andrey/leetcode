# https://leetcode.com/problems/first-unique-character-in-a-string/

import collections


class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: int
        """
        ctr = collections.Counter(s)
        for i, letter in enumerate(s):
            if ctr[letter] == 1:
                return i
        return -1
