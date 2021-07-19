# https://leetcode.com/problems/merge-sorted-array/

class Solution(object):
    def merge(self, nums1, m, nums2, n):
        """
        :type nums1: List[int]
        :type m: int
        :type nums2: List[int]
        :type n: int
        :rtype: None Do not return anything, modify nums1 in-place instead.
        """
        if n == 0:
            return
        if m == 0:
            for i in range(n):
                nums1[len(nums1) - n + i] = nums2[len(nums2) - n + i]
            return

        cur1 = len(nums1) - m - n
        cur2 = len(nums2) - n

        while nums1[cur1] <= nums2[cur2] and m != 0:
            m -= 1
            cur1 += 1

        for i in range(len(nums1) - n, cur1, -1):
            nums1[i] = nums1[i - 1]

        nums1[cur1] = nums2[cur2]
        self.merge(nums1, m, nums2, n - 1)
