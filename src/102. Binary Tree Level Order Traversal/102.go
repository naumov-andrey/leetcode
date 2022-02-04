/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return make([][]int, 0)
    }
    res := make([][]int, 1)
    
    lvl, lvlValues := goDownOneLevel([]*TreeNode{root})
    res[0] = lvlValues
    for len(lvl) != 0 {
        lvl, lvlValues = goDownOneLevel(lvl)
        res = append(res, lvlValues)
    }
    return res
}

func goDownOneLevel(lvl []*TreeNode) ([]*TreeNode, []int) {
    nextLvl := make([]*TreeNode, 0)
    lvlValues := make([]int, len(lvl))
    
    for i, n := range lvl {
        if n.Left != nil {
            nextLvl = append(nextLvl, n.Left)
        }
        if n.Right != nil {
            nextLvl = append(nextLvl, n.Right)
        }
        
        lvlValues[i] = n.Val
    }
    
    return nextLvl, lvlValues
}
