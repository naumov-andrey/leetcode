/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	c := (len(nums) - 1) / 2
	root := &TreeNode{nums[c], nil, nil}

	root.Left = sortedArrayToBST(nums[:c])
	root.Right = sortedArrayToBST(nums[c+1:])
	return root
}
