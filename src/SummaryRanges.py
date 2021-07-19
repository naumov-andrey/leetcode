# https://leetcode.com/problems/summary-ranges/

class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        result = []

        i = 0
        while i < len(nums):
            r = l = nums[i]
            i += 1
            while i < len(nums) and nums[i] - nums[i - 1] == 1:
                r = nums[i]
                i += 1

            if r == l:
                result.append(f'{l}')
            else:
                result.append(f'{l}->{r}')

        return result
