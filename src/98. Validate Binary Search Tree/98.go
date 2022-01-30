/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTWithRightBound(root.Left, root.Val) && isValidBSTWithLeftBound(root.Right, root.Val)
}

func isValidBSTWithLeftBound(root *TreeNode, l int) bool {
	return root == nil ||
		(root.Val > l &&
			isValidBSTWithBounds(root.Left, l, root.Val) &&
			isValidBSTWithLeftBound(root.Right, root.Val))
}

func isValidBSTWithRightBound(root *TreeNode, r int) bool {
	return root == nil ||
		(root.Val < r &&
			isValidBSTWithRightBound(root.Left, root.Val) &&
			isValidBSTWithBounds(root.Right, root.Val, r))
}

func isValidBSTWithBounds(root *TreeNode, l, r int) bool {
	return root == nil ||
		(root.Val > l &&
			root.Val < r &&
			isValidBSTWithBounds(root.Left, l, root.Val) &&
			isValidBSTWithBounds(root.Right, root.Val, r))
}
