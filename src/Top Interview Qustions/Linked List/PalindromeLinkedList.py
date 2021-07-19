# https://leetcode.com/problems/palindrome-linked-list/

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def isPalindrome(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        runner, half = head, head
        stack = []
        while runner is not None:
            runner = runner.next
            stack.append(half.val)
            if runner is not None:
                runner = runner.next
                half = half.next
        while half is not None:
            if half.val != stack.pop():
                return False
            half = half.next
        return True
