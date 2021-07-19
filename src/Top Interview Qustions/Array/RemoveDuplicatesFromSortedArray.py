# https://leetcode.com/problems/remove-duplicates-from-sorted-array/

class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        repeats = 0
        for i in range(1, len(nums)):
            if nums[i] == nums[i - 1]:
                repeats += 1
                continue
            nums[i - repeats] = nums[i]
        return len(nums) - repeats
