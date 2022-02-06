func abs(v int) int {
    if v < 0 {
        return -v
    }
    return v
}

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    
    closest := nums[0] + nums[1] + nums[2]
    for i := 0; i < len(nums) - 2; i++ {
        j, k := i + 1, len(nums) - 1
        
        for k > j {
            cur := nums[i] + nums[j] + nums[k]
            
            if cur == target {
                return target
            }
            
            if  abs(target - cur) < abs(target - closest) {
                closest = cur
            }
            
            if cur < target {
                j++
            } else if cur > target {
                k--
            }
        }
    }
    
    return closest
}
