# https://leetcode.com/problems/remove-nth-node-from-end-of-list/

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        cur = head
        cur_delay = head
        while n > 0:
            cur = cur.next
            n -= 1
        if cur is None:
            return head.next
        while cur.next is not None:
            cur = cur.next
            cur_delay = cur_delay.next
        cur_delay.next = cur_delay.next.next
        return head
