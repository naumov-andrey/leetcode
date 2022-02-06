func binarySearch(nums []int, lo, hi, val int) int {
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        
        if val == nums[mid] {
            return mid
        }
        if val < nums[mid] {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    
    return -1
}

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return make([][]int, 0)
    }
    sort.Ints(nums)
    
    triplets := make(map[[3]int][]int)
    
    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            if k := binarySearch(nums, j + 1, len(nums) - 1, - (nums[i] + nums[j])); k != -1 {
                triplets[[3]int{nums[i], nums[j], nums[k]}] = []int{nums[i], nums[j], nums[k]}
            } 
        }
    }
    
    res := make([][]int, 0, len(triplets))
    for _, v := range triplets {
        res = append(res, v)
    }
    
    return res
}
