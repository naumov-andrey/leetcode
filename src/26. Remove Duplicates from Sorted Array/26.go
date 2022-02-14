// func removeDuplicates(nums []int) int {
//     shift := 0
    
//     for i := 0; i < len(nums); i++ {
//         for i > 0 && i < len(nums) && nums[i - 1] == nums[i] {
//             i++
//             shift++
//         }
        
//         if i < len(nums) {
//             nums[i - shift] = nums[i]
//         }
//     }
    
//     return len(nums) - shift
// }

func removeDuplicates(nums []int) int {
    shift := 0
    
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] == nums[i] {
            shift++
        } else {
            nums[i - shift] = nums[i]
        }
    }
    
    return len(nums) - shift
}
