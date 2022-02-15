type Solution struct {
    original []int
    data     []int
    
    swap     func(int, int)
}


func Constructor(nums []int) Solution {
    original := make([]int, len(nums))
    copy(original, nums)
    
    data := nums
    
    swap := func(i, j int) {
        data[i], data[j] = data[j], data[i]
    }
    
    return Solution{original, data, swap}
}


func (this *Solution) Reset() []int {
    copy(this.data, this.original)
    return this.data
}


func (this *Solution) Shuffle() []int {
    rand.Shuffle(len(this.data), this.swap)
    return this.data
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
 