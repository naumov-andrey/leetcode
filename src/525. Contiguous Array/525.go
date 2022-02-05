func findMaxLength(nums []int) int {
    res := 0
    diffToIndex := map[int]int{0: -1}
    
    diff := 0
    for i, v := range nums {
        if v == 0 {
            diff -= 1
        } else {
            diff += 1
        }
        
        if idx, ok := diffToIndex[diff]; ok {
            if i - idx > res {
                res = i - idx
            }
        } else {
            diffToIndex[diff] = i
        }
    }
    
    return res
}
