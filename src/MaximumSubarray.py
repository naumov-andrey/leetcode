# https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/566/

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        # return self.maxSubArrayImpl(nums, 0, nums[0])
        index = 0
        current = 0
        result = nums[0]
        while index < len(nums):
            result = max(result, current + nums[index])
            if current + nums[index] < 0:
                current = 0
            else:
                current += nums[index]
            index += 1
        return result


    def maxSubArrayImpl(self, nums: List[int], current: int, result: int) -> int:
        if not nums:
            return result

        result = max(result, current + nums[0])

        if current + nums[0] < 0:
            return self.maxSubArrayImpl(nums[1:], 0, result)

        return self.maxSubArrayImpl(nums[1:], current + nums[0], result)
