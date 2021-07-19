# https://leetcode.com/problems/linked-list-cycle/

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        is_in_list = {}
        while True:
            if head is None:
                return False
            if head not in is_in_list:
                is_in_list[head] = True
            else:
                return True
            head = head.next
