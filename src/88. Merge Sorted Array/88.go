func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j := 0, 0
    for ; i < m + n && j < n; i++ {
        if nums1[i] >= nums2[j] {
            c := 0
            for j + c < n && nums1[i] >= nums2[j + c] {
                c++
            }
            
            for k := m + n - 1; k >= i + c; k-- {
                nums1[k] = nums1[k - c]
            }
            
            for k := 0; k < c; k++ {
                nums1[i + k] = nums2[j + k]
            }
            
            i += c
            j += c
        }
    }
    for ; j < n; j++ {
        nums1[m + j] = nums2[j]
    }
}
