/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return isMirrored(root.Left, root.Right)
}

func isMirrored(p, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil {
        return false
    }
    
    return p.Val == q.Val && isMirrored(p.Left, q.Right) && isMirrored(p.Right, q.Left)
}
