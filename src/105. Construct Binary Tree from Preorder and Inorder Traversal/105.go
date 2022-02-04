/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func index(nums []int, val int) int {
    for i, v := range nums {
        if val == v {
            return i
        }
    }
    return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    
    root := &TreeNode{preorder[0], nil, nil}
    rootInorderIdx := index(inorder, root.Val)
    
    root.Left = buildTree(preorder[1:1 + rootInorderIdx], inorder[:rootInorderIdx])
    root.Right = buildTree(preorder[1 + rootInorderIdx:], inorder[rootInorderIdx + 1:])
    
    return root
}
