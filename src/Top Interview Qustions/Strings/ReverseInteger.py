# https://leetcode.com/problems/reverse-integer/

class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        str_x = str(x)
        if x < 0:
            str_x = str_x[1:][::-1]
            if (len(str_x) == 10 and str_x > str(2 ** 31)) or len(str_x) > 10:
                return 0
            str_x = '-{}'.format(str_x)
        else:
            str_x = str_x[::-1]
            if (len(str_x) == 10 and str_x > str(2 ** 31 - 1)) or len(str_x) > 10:
                return 0
        return int(str_x)
