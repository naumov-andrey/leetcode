func intersection(nums1 []int, nums2 []int) []int {
    s := make(map[int]bool)
    
    for _, v := range nums1 {
        s[v] = true
    }
    
    res := make([]int, 0)
    for _, v := range nums2 {
        if ok, _ := s[v]; ok {
            res = append(res, v)
            s[v] = false
        }
    }
    
    return res
}
