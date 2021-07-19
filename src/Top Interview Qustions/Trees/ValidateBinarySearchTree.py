# https://leetcode.com/problems/validate-binary-search-tree/

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution(object):
    def isValidBSTImpl(self, root, interval):
        isValid = interval[0] < root.val < interval[1]
        if root.left is not None:
            isValid = isValid and root.val > root.left.val and self.isValidBSTImpl(root.left, (
            interval[0], root.val))
        if root.right is not None:
            isValid = isValid and root.right.val > root.val and self.isValidBSTImpl(root.right, (
            root.val, interval[1]))
        return isValid

    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.isValidBSTImpl(root, (-float('inf'), float('inf')))
