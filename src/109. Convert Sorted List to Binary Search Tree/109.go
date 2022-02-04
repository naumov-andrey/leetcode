/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
   return  sortedListToBSTHelper(head, nil)
}

func sortedListToBSTHelper(l, r *ListNode) *TreeNode {
    if l == nil || l == r {
        return nil
    }
    
    half, runner := l, l
    for runner != r && runner.Next != r {
        half = half.Next
        runner = runner.Next.Next
    }
    
    root := &TreeNode{half.Val, nil, nil}
    root.Left = sortedListToBSTHelper(l, half)
    root.Right = sortedListToBSTHelper(half.Next, r)
    
    return root
}
