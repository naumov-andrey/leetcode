# https://leetcode.com/problems/symmetric-tree/

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution(object):
    def compareTreeNodes(self, first, second):
        if first is None and second is None:
            return True
        if first is not None and second is not None:
            return first.val == second.val and (self.compareTreeNodes(first.left, second.right)
                                                and self.compareTreeNodes(first.right, second.left))
        return False

    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.compareTreeNodes(root.left, root.right)
