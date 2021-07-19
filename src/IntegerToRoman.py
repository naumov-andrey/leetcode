# https://leetcode.com/problems/integer-to-roman/

class Solution(object):
    def intToRomanHelper_(self, num, literals):
        if num == 0:
            return ''
        if num < 4:
            return ''.join(literals[0] * num)
        if num == 4:
            return ''.join(literals[0:2])
        if num < 9:
            return ''.join([literals[1], literals[0] * (num - 5)])
        return ''.join([literals[0], literals[2]])

    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        units = [0] * 4
        i = 0
        while num > 0:
            units[i] = num % 10
            num = num // 10
            i += 1
        roman = ['M' * units[3],
                 self.intToRomanHelper_(units[2], ['C', 'D', 'M']),
                 self.intToRomanHelper_(units[1], ['X', 'L', 'C']),
                 self.intToRomanHelper_(units[0], ['I', 'V', 'X']), ]
        return ''.join(roman)
