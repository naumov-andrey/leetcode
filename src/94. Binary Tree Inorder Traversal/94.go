/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func inorderTraversal(root *TreeNode) []int {
//     if root == nil {
//         return nil
//     }

//     l := inorderTraversal(root.Left)
//     m := root.Val
//     r := inorderTraversal(root.Right)

//     return append(append(l, m), r...)
// }

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	s := make([]*TreeNode, 0)
	cur := root

	for cur != nil || len(s) != 0 {
		for cur != nil {
			// push
			s = append(s, cur)

			cur = cur.Left
		}
		// pop
		cur = s[len(s)-1]
		s = s[:len(s)-1]

		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}
