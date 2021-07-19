# https://leetcode.com/problems/roman-to-integer/

class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        n = len(s)
        roman_numerals = {'M': 1000,
                          'CM': 900,
                          'D': 500,
                          'CD': 400,
                          'C': 100,
                          'XC': 90,
                          'L': 50,
                          'XL': 40,
                          'X': 10,
                          'IX': 9,
                          'V': 5,
                          'IV': 4,
                          'I': 1}
        i = 0
        while i < n:
            if i < n - 1 and s[i:i + 2] in roman_numerals:
                result += roman_numerals[s[i:i + 2]]
                i += 1
            else:
                result += roman_numerals[s[i]]
            i += 1
        return result
