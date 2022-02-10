func containsDuplicate(nums []int) bool {
    s := make(map[int]struct{})
    
    for _, v := range nums {
        if _, ok := s[v]; ok {
            return true
        }
        s[v] = struct{}{}
    }
    
    return false
}

// func containsDuplicate(nums []int) bool {
//     sort.Ints(nums)
    
//     for i := 1; i < len(nums); i++ {
//         if nums[i] == nums[i - 1] {
//             return true
//         }
//     }
    
//     return false
// }

// func containsDuplicate(nums []int) bool {
//     for i := 0; i < len(nums); i++ {
//         for j := i + 1; j < len(nums); j++ {
//             if nums[i] == nums[j] {
//                 return true
//             }
//         }
//     }
    
//     return false
// }

// func containsDuplicate(nums []int) bool {
//     s := make(map[int]struct{})
    
//     for _, v := range nums {
//         s[v] = struct{}{}
//     }
    
//     return len(s) < len(nums)
// }
