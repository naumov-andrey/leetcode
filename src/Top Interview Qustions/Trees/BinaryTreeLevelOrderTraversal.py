# https://leetcode.com/problems/binary-tree-level-order-traversal/

from collections import deque


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        result = []

        to_process = deque([root])
        while to_process:

            level = []
            next_level = deque()
            while to_process:

                current = to_process.popleft()
                if not current:
                    continue
                level.append(current.val)
                next_level.extend((current.left, current.right))

            if level:
                result.append(level)
            to_process.extend(next_level)

        return result
