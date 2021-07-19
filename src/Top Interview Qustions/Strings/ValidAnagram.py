# https://leetcode.com/problems/valid-anagram/

from collections import Counter


class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        s_cntr = Counter(s)
        t_cntr = Counter(t)
        return s_cntr == t_cntr
